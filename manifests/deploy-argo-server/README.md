This folder contains manifests to deploy Argo Workflows Standalone UI on Openshift,

This is useful for debugging Data Science Pipelines in a more technical view than offered by
the more 'specialized-for-data-scientists' views provided by the DSP or KFP UIs 

### Deploy Argo Workflows Standalone UI

```bash
git clone https://github.com/opendatahub-io/dsp-dev-tools.git
cd manifests/deploy-argo-server
oc apply -n <some-namespace> -k standalone/openshift
```
