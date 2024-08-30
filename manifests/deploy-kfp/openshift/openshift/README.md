This folder contains manifests to help with deploying Kubeflow Standalone on Openshift. 

You can deploy standard, or with multi-user auth enabled. 
Before deploying, you will first need to set the remote fork from which to pull the base manifests from. From this we 
apply some patches to get the deployment to work on Openshift. You can inspect these patch files/additions in: 

- manifests/deploy-kfp/openshift/openshift/base
- manifests/deploy-kfp/openshift/openshift/auth


### Configure your repo: 

We'll use the upstream, default branch. 
```bash
# Adjust this to your fork if needed, no trailing slash
REPO=https://github.com/kubeflow/pipelines
BRANCH=master
```

### Deploy base no auth/multi-user

```bash
git clone https://github.com/opendatahub-io/dsp-dev-tools.git
cd manifests/deploy-kfp/openshift/openshift/base
./add_resources.sh $REPO $BRANCH
```

### Deploy with auth/multi-user

```bash
git clone https://github.com/opendatahub-io/dsp-dev-tools.git
cd manifests/deploy-kfp/openshift/openshift/auth
./add_resources.sh $REPO $BRANCH
```

### Add your own images: 

To add your own images: 

```bash
# Argo wf controller
SRC=gcr.io/ml-pipeline/workflow-controller
TARGET=quay.io/argoproj/workflow-controller:v3.4.16
kustomize edit set image ${SRC}=${TARGET}

# api-server
SRC=gcr.io/ml-pipeline/api-server
TARGET=opendatahub/ds-pipelines-api-server:latest
kustomize edit set image ${SRC}=${TARGET}

# Persistent Agent
SRC=gcr.io/ml-pipeline/persistenceagent
TARGET=quay.io/hukhan/persistenceagent:latest
kustomize edit set image ${SRC}=${TARGET}

# frontend
SRC=gcr.io/ml-pipeline/frontend
TARGET=quay.io/opendatahub/ds-pipelines-frontend:latest
kustomize edit set image ${SRC}=${TARGET}

# scheduledworkflow
SRC=gcr.io/ml-pipeline/scheduledworkflow
TARGET=quay.io/opendatahub/ds-pipelines-scheduledworkflow:latest
kustomize edit set image ${SRC}=${TARGET}
```

> Note the resource format, read more about it [here](https://github.com/kubernetes-sigs/kustomize/blob/master/examples/remoteBuild.md).

### Add custom driver/launcher images 

These are set as ENV vars to api server, so they need to be changed manually: 

```bash
LAUNCHER_IMAGE=quay.io/opendatahub/ds-pipelines-launcher:latest
DRIVER_IMAGE=quay.io/opendatahub/ds-pipelines-driver:latest


# from repo root
cd manifests/deploy-kfp/openshift/openshift/base
var=$LAUNCHER_IMAGE yq -i '.spec.template.spec.containers[0].env[0].value = strenv(var)' api-server-patch.yaml
var=$DRIVER_IMAGE yq -i '.spec.template.spec.containers[0].env[1].value = strenv(var)' api-server-patch.yaml

```