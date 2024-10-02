package main

import (
	"flag"

	"github.com/gmfrasca/dsp-devcompiler/fullcompiler"
)

func main() {
	var pipelinePackageVar, pipelineCompVar, irVar, outputVar, pspVar string
	flag.StringVar(&pipelinePackageVar, "pypackage", "pipelines.pipeline", "Pipeline Python Definition Package containing PipelineDef")
	flag.StringVar(&pipelineCompVar, "componentName", "pipeline", "Name of Component in Python Pkg to compile")
	flag.StringVar(&irVar, "ir", "pipelines/ir.json", "Pipeline Intermediate Representation (IR) JSON File to read/write")
	flag.StringVar(&outputVar, "output", "pipelines/workflow.yaml", "Compiled Argo Workflow YAML Filepath to write to")
	flag.StringVar(&pspVar, "psp", "", "Pod Spec Patch to apply to Argo Compilation")
	flag.Parse()

	fc := fullcompiler.FullCompiler{}
	fc.Compile(pipelinePackageVar, pipelineCompVar, irVar, outputVar, pspVar)
}
