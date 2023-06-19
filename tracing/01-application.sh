export ISTIO_VERSION=1.16.2

# install istio
kubectl label namespace default istio-injection=enabled

# deploy the bookinfo applications
kubectl apply -f https://raw.githubusercontent.com/istio/istio/$ISTIO_VERSION/samples/bookinfo/platform/kube/bookinfo.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/$ISTIO_VERSION/samples/bookinfo/networking/bookinfo-gateway.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/$ISTIO_VERSION/samples/bookinfo/networking/destination-rule-all.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/$ISTIO_VERSION/samples/bookinfo/networking/virtual-service-all-v1.yaml

# generate traffic
kubectl apply -f https://raw.githubusercontent.com/mrproliu/skywalking-network-profiling-demo/main/resources/traffic-generator.yaml