import os

import kfp
from pipeline import fraud_training_pipeline
import subprocess

# Pipeline input parameters
metadata = {
    "datastore": {"url": "https://raw.githubusercontent.com/rh-aiservices-bu/fraud-detection/main/data/card_transdata.csv"},
    "hyperparameters": {"epochs": 2}
}

# DSPA Environment info
namespace = os.environ.get('NAMESPACE', "namespace")
dspa = os.environ.get('DSPA_NAME', "dspa")
run_name = os.environ.get('RUN_NAME', "run_name")


def terminal(cmd):
    proc = subprocess.run(cmd.split(), stdout=subprocess.PIPE)
    out = str(proc.stdout, encoding='utf-8')
    return out.strip()


if __name__ == '__main__':

    dspa_host = terminal(f"oc get routes -n {namespace} ds-pipeline-{dspa} --template={{{{.spec.host}}}}")
    route = f"https://{dspa_host}"
    token = terminal("oc whoami --show-token")

    client = kfp.Client(host=route, verify_ssl=False, existing_token=token)

    client.create_run_from_pipeline_func(
        fraud_training_pipeline,
        arguments=metadata,
        experiment_name="fraud-training",
        namespace="mlops-dev-zone",
        enable_caching=True
    )