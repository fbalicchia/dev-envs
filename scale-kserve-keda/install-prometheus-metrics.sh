kubectl apply -f ./metrics-server-ignore-ssl.yaml
helm install prometheus prometheus-community/kube-prometheus-stack --values ./values-latest.yaml -n metrics
