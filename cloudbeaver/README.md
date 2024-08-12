# cloudbeaver-k8s

Deploy cloudbeaver

# Steps: 

```bash
CLOUD_BEAVER_NAMESPACE=cloudbeaver
oc new-project $CLOUD_BEAVER_NAMESPACE
cd overlays/dsp
kustomize build . | oc -n $CLOUD_BEAVER_NAMESPACE apply -f -
```

# Deploy network policy if using default mariadb 

```bash
DSPA_NAME=sample
DSPA_NAMESPACE=dspa
cd dspa-network-policy

# Set the network policy to select the mariadb pod, otherwise the pod only allows access from 
var=$DSPA_NAME yq -i '.spec.podSelector.matchLabels.app = "mariadb-"+env(var) ' np.yaml

oc apply -n $DSPA_NAMESPACE -f np.yaml
```

Open the cloudbeaver route, make an account, add mariadb DB, find the secrets and svc dns from ocp console. 
Test connection. Once green, you are good to go.