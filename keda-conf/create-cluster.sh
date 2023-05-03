kind delete clusters --all
cat <<EOF | kind create cluster -n keda --config -
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  image: kindest/node:v1.24.1
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
EOF
kubectl create ns argocd
kubectl create ns paas-application-staging
kubectl apply -f https://raw.githubusercontent.com/argoproj/argo-cd/master/manifests/install.yaml -n argocd
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
helm upgrade --install keda kedacore/keda --namespace keda --create-namespace --wait
helm install http-add-on kedacore/keda-add-ons-http --namespace keda
helm install xkcd ./xkcd  --namespace keda
cd websocket-app/;docker build -t websockets-test:1.0 .
kind load docker-image websockets-test:1.0 -n keda
kubectl apply -f websocket-deploy.yaml

