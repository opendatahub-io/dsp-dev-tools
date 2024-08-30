#!/usr/bin/env bash

REPO=$1
BRANCH=$2

kustomize edit add resource ${REPO}//manifests/kustomize/base/installs/generic?ref=${BRANCH}
kustomize edit add resource ${REPO}//manifests/kustomize/base/pipeline/cluster-scoped?ref=${BRANCH}
kustomize edit add resource ${REPO}//manifests/kustomize/base/metadata/base?ref=${BRANCH}
kustomize edit add resource ${REPO}//manifests/kustomize/third-party/argo/installs/cluster?ref=${BRANCH}
kustomize edit add resource ${REPO}//manifests/kustomize/third-party/minio/base?ref=${BRANCH}
kustomize edit add resource ${REPO}//manifests/kustomize/third-party/mysql/base?ref=${BRANCH}