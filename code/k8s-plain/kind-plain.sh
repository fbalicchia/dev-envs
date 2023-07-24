#!/usr/bin/env bash
set -o errexit

main() {
  echo "${em}Allocating...${me}"
  kubernetes
  certmanager
  metrics
  prometheus
}


kubernetes() {
  echo "${em}① Kubernetes${me}"
  kind delete clusters --all
docker container stop kind-registry || true
docker container rm kind-registry || true
# create registry container unless it already exists
reg_name='kind-registry'
reg_port='5000'
if [ "$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)" != 'true' ]; then
  docker run \
    -d --restart=always -p "127.0.0.1:${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi


# connect the registry to the cluster network if not already connected
if [ "$(docker inspect -f='{{json .NetworkSettings.Networks.kind}}' "${reg_name}")" = 'null' ]; then
  docker network connect "kind" "${reg_name}"
fi

cat << EOF |  kind create cluster --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
    endpoint = ["http://${reg_name}:5000"]
nodes:
- role: control-plane
  image: kindest/node:v1.27.3
  extraPortMappings:
  - containerPort: 31080
    hostPort: 80
  - containerPort: 31443
    hostPort: 443
EOF


# Document the local registry
# https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-registry-hosting
  namespace: kube-public
data:
  localRegistryHosting.v1: |
    host: "localhost:${reg_port}"
    help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
EOF
}

certmanager() {
  echo "${em}③  certmanager${me}"
  kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.3.0/cert-manager.yaml
  kubectl wait --for=condition=available --timeout=600s deployment/cert-manager-webhook -n cert-manager
  echo "😀 Successfully installed Cert Manager"
}


metrics() {
  echo "${em}④  metrics${me}"
  kubectl apply -f ./metrics-server-ignore-ssl.yaml
  echo "😀 Successfully installed metrics server"
}


prometheus() {
  echo "${em}⑤ prometheus${me}"
  kubectl apply -f ./prometheus.yaml
  echo "😀 Successfully installed Prometheus"
}



main "$@"



