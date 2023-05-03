kubectl create ns metrics
helm install prometheus prometheus-community/kube-prometheus-stack --values ./values-latest.yaml -n metrics
kubectl apply -f ./fix-app.yaml -n default
kubectl apply -f ./service-monitor-definition.yaml -n default
kubectl apply -f ./metrics-server-ignore-ssl.yaml