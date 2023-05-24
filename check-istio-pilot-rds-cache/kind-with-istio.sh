#!/usr/bin/env bash
set -o errexit


main(){
  kubernetes
  istio
  certmanager
  metrics
}


metrics(){
  kubectl apply -f $PWD/addons/prometheus.yaml
  kubectl apply -f $PWD/addons/grafana.yaml
}



kubernetes() {
  echo "${em}① Kubernetes${me}"
  kind delete clusters --all
  cat << EOF |  kind create cluster --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  image: kindest/node:v1.24.1
  extraPortMappings:
  - containerPort: 31080
    hostPort: 80
  - containerPort: 31443
    hostPort: 443
EOF
}

istio() {
  echo "${em}② istio${me}"
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
  components:
    pilot:
      enabled: true
      k8s:
        env:
        - name: PILOT_ENABLE_RDS_CACHE
          value: "false"


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


cat << EOF > ./ingress-gateway.yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: test-gateway
spec:
  selector:
    istio: ingressgateway
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - "*"
EOF

kubectl apply -f ./ingress-gateway.yaml -n istio-system
}


certmanager(){
  # Install Cert Manager
  kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.3.0/cert-manager.yaml
  kubectl wait --for=condition=available --timeout=600s deployment/cert-manager-webhook -n cert-manager
  echo "😀 Successfully installed Cert Manager"
}


main @


