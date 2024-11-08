FROM registry.fedoraproject.org/fedora-toolbox:40

ARG YQ_VERSION="v4.40.3"
ARG OKD_RELEASE="4.14.0-0.okd-2023-11-14-101924"
ARG KUSTOMIZE_VERSION="v5.5.0"
ARG MC_VERSION="mc.RELEASE.2024-11-05T11-29-45Z"

# Put useful tools here
RUN dnf install -y s3cmd wget pip podman skopeo &&\
    # Install yq
    curl -o /usr/local/bin/yq -L https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_amd64 &&\
    chmod +x /usr/local/bin/yq &&\
    # Install oc
    curl -L https://github.com/openshift/okd/releases/download/${OKD_RELEASE}/openshift-client-linux-${OKD_RELEASE}.tar.gz  | tar -xzf - -C /usr/local/bin  &&\
    chmod +x /usr/local/bin/oc && chmod +x /usr/local/bin/kubectl &&\
    # Install kustomize
    curl -L https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/${KUSTOMIZE_VERSION}/kustomize_${KUSTOMIZE_VERSION}_linux_amd64.tar.gz |  tar -xzf - -C /usr/local/bin  &&\
    # Install minio cli
    echo https://dl.min.io/client/mc/release/linux-amd64/${MC_VERSION} &&\
    curl -L https://dl.min.io/client/mc/release/linux-amd64/${MC_VERSION} -o /usr/local/bin/mc &&\
    chmod +x /usr/local/bin/mc &&\
    # Install huggingface cli
    pip install -U "huggingface_hub[cli]"


CMD /bin/bash
