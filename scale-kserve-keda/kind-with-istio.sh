#!/usr/bin/env bash
set -o errexit
kind delete clusters --all
docker container stop kind-registry
docker container rm kind-registry
# create registry container unless it already exists
reg_name='kind-registry'
reg_port='5001'
if [ "$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)" != 'true' ]; then
  docker run \
    -d --restart=always -p "127.0.0.1:${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi


# connect the registry to the cluster network if not already connected
if [ "$(docker inspect -f='{{json .NetworkSettings.Networks.kind}}' "${reg_name}")" = 'null' ]; then
  docker network connect "kind" "${reg_name}"
fi

cat << EOF |  kind create cluster --name scale --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
    endpoint = ["http://${reg_name}:5000"]
nodes:
- role: control-plane
  image: kindest/node:v1.24.1
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


cat << EOF > ./istio-minimal-operator.yaml
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
spec:
  values:
    global:
      proxy:
        autoInject: enabled
      useMCP: false
    gateways:
      istio-ingressgateway: 
        name: cluster-local-gateway
        runAsRoot: true
  addonComponents:
    pilot:
      enabled: true
    tracing:
      enabled: false
    kiali:
      enabled: false
    prometheus:
      enabled: false
  components:
    ingressGateways:
      - name: istio-ingressgateway
        enabled: true
      - name: cluster-local-gateway
        enabled: true
        label:
          istio: cluster-local-gateway
          app: cluster-local-gateway
        k8s:
          service:
            type: ClusterIP
            ports:
            - port: 15020
              name: status-port
            - port: 80
              name: http2
            - port: 443
              name: https
EOF
istioctl manifest apply -f istio-minimal-operator.yaml -y

cat << EOF > ./patch-ingressgateway-nodeport.yaml
spec:
  type: NodePort
  ports:
  - name: http2
    nodePort: 31080
    port: 80
    protocol: TCP
    targetPort: 80
EOF
kubectl patch service istio-ingressgateway -n istio-system --patch "$(cat ./patch-ingressgateway-nodeport.yaml)"

# Install Cert Manager
kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.3.0/cert-manager.yaml
kubectl wait --for=condition=available --timeout=600s deployment/cert-manager-webhook -n cert-manager
cd ..
echo "😀 Successfully installed Cert Manager"




kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
kubectl create serviceaccount -n kubernetes-dashboard admin-user
kubectl apply --filename https://github.com/knative/serving/releases/download/knative-v1.9.0/serving-crds.yaml
kubectl apply --filename https://github.com/knative/serving/releases/download/knative-v1.9.0/serving-core.yaml
kubectl apply --filename https://github.com/knative/net-istio/releases/download/knative-v1.9.0/release.yaml
kubectl patch configmap/config-domain \
  --namespace knative-serving \
  --type merge \
  --patch '{"data":{"127.0.0.1.nip.io":""}}'


kubectl apply -f https://github.com/kserve/kserve/releases/download/v0.10.1/kserve.yaml
kubectl wait --for=condition=ready pod -l control-plane=kserve-controller-manager -n kserve --timeout=300s
kubectl apply -f https://github.com/kserve/kserve/releases/download/v0.10.1/kserve-runtimes.yaml
echo "😀 Successfully installed KServe"
kubectl create ns kserve-test
kubectl patch -n knative-serving cm config-deployment \
    --type='json' \
    -p='[{"op":"add","path":"/data/registries-skipping-tag-resolving","value":"kind.local,ko.local,dev.local,localhost:5001"}]'

helm upgrade --install keda kedacore/keda --namespace keda --create-namespace --wait

kubectl create ns metrics