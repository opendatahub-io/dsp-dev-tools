import kfp
from pipeline_no_cache import fraud_training_pipeline

if __name__ == '__main__':
    kfp.compiler.Compiler().compile(
        pipeline_func=fraud_training_pipeline,
        package_path='fraud_detection_no_cache.yaml'
    )