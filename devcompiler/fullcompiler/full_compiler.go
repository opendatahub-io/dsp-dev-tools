// Copyright 2024 Giulio Frasca
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fullcompiler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/kubeflow/pipelines/api/v2alpha1/go/pipelinespec"
	"github.com/kubeflow/pipelines/backend/src/v2/compiler/argocompiler"
	"google.golang.org/protobuf/encoding/protojson"
	"sigs.k8s.io/yaml"
)

type FullCompiler struct{}

type MockPipelineIR struct {
	PipelineSpec  map[string]interface{} `json:"pipelineSpec"`
	RuntimeConfig map[string]interface{} `json:"runtimeConfig"`
}

func (fc FullCompiler) LoadIRFile(path string, platformSpecPath string) (*pipelinespec.PipelineJob, *pipelinespec.SinglePlatformSpec) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	job := &pipelinespec.PipelineJob{}
	if err := protojson.Unmarshal(content, job); err != nil {
		log.Fatalf("Failed to parse pipeline job, error: %s, job: %v", err, string(content))
	}

	platformSpec := &pipelinespec.PlatformSpec{}
	if platformSpecPath != "" {
		content, err = os.ReadFile(platformSpecPath)
		if err != nil {
			log.Fatal(err)
		}
		if err := protojson.Unmarshal(content, platformSpec); err != nil {
			log.Fatalf("Failed to parse platform spec, error: %s, spec: %v", err, string(content))
		}
		return job, platformSpec.Platforms["kubernetes"]
	}
	return job, nil
}

func (fc FullCompiler) Compile(inputPipelinePackage, pipelineComponent, outputIRPath, outputWorkflowPath, platformSpecPatch string) {
	err := fc.CompilePipelineToIR(inputPipelinePackage, pipelineComponent, outputIRPath)
	if err != nil {
		log.Fatal(err)
	}

	err = fc.EnrichIRJSON(outputIRPath)
	if err != nil {
		log.Fatal(err)
	}

	err = fc.CompileIRToWorkflow(outputIRPath, outputWorkflowPath, platformSpecPatch)
	if err != nil {
		log.Fatal(err)
	}
}

func (fc FullCompiler) CompilePipelineToIR(inputPipelinePackage, pipelineComponent, outputIRPath string) error {
	log.Printf("Executing Python Pipeline (%s) to Generate PipelineIR (%s)", inputPipelinePackage, outputIRPath)
	pyscript := fmt.Sprintf("from %s import %s; from kfp import compiler; compiler.Compiler().compile(%s, '%s');", inputPipelinePackage, pipelineComponent, pipelineComponent, outputIRPath)
	cmd := exec.Command("python", "-c", pyscript)
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (fc FullCompiler) EnrichIRJSON(irPath string) error {
	log.Printf("Enriching Intermediate Rep JSON (%s) to allow compilation into Workflow", irPath)
	// TODO(gfrasca): Actually generate the runtime config
	mockRuntimeConfigJSON := []byte(`{"parameters":{"text":{"stringValue":"foo"}}}`)
	var mockRuntimeConfig map[string]interface{}
	json.Unmarshal([]byte(mockRuntimeConfigJSON), &mockRuntimeConfig)

	data, err := os.Open(irPath)
	if err != nil {
		return err
	}
	defer data.Close()

	byteValue, _ := io.ReadAll(data)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	output := MockPipelineIR{
		PipelineSpec:  result,
		RuntimeConfig: mockRuntimeConfig,
	}

	outputJSON, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(irPath, outputJSON, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fc FullCompiler) CompileIRToWorkflow(inputIRPath, outputWorkflowPath, platformSpecPatch string) error {
	log.Printf("Compiling PipelineIR JSON (%s) to Argo Workflow YAML (%s)", inputIRPath, outputWorkflowPath)
	job, platformSpec := fc.LoadIRFile(inputIRPath, platformSpecPatch)

	wf, err := argocompiler.Compile(job, platformSpec, nil)
	if err != nil {
		log.Fatal(err)
	}

	got, err := yaml.Marshal(wf)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(outputWorkflowPath, got, 0777)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
