# Data Science Pipelines Developer Compilers

Development tool for compiling KFP/DSP Pipelines

## Full e2e Compiler

This Compiler Generates PipelineIR JSON and ArgoWFs from a given input Python package and Pipeline Definition Name


### Installation
This largely can be used as-is.

However, because this is intended for developers, there is a line in go.mod:
`replace github.com/kubeflow/pipelines v0.0.0-20240613070908-b57f9e858880 => /home/gfrasca/src/data-science-pipelines`, which points to the local source for DSP such that you can update and debug the internal IR Compiler dynamically.   Adjust the directory pointer accordingly (#TODO: show how to do this using appropriate `go` commands instead)

### Execution
 
To use this, follow these steps:
1. Place the pipeline definition you would like to compile in the `pipelines` directory.

2. Run the following command from repo root: 
    ```
    go run . \
        -pypackage=pipelines.<NAME OF YOUR PYTHON FILE> \
        -componentName=<COMPONENT NAME OF YOUR PIPELINE>
    ```


This will generate the IntermediateRepresentation JSON file in `pipelines/ir.json` and the compiled Argo Workflow YAML file in `pipelines/workflow.yaml`.  

Note that you can override these values using `-ir` and `-output` flags, respectively.


### Example Usage
There is a simple sample pipeline (`pipelines/pipeline.py`) with a Pipeline Definition named `pipeline`, as designated by the @dsl.pipeline decorator.  To generate an IR and ArgoWF, the full command could look like:

```
go run . \
   -pypackage=pipelines.pipeline \
   -componentName=pipeline \
   -ir=pipelines/ir.json \
   -output=pipelines/workflow.yaml
```

Adjust the parameter values to your needs


### Disclaimer
No input validation is provided, and as we are executing pipelines files directly this is inherently not safe for any form of live/production use.  For Local Development purposes only.
