from kfp import compiler, dsl

@dsl.component()
def print_op(s: str):
    print(s)

@dsl.pipeline(name='my-pipeline')
def pipeline():
    print_op(s='foo')
