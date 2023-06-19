#/bin/bash

# deploy skywalking backend, ui
kubectl apply -f skywalking.yaml


cat <<EOF | kubectl apply -f -
apiVersion: telemetry.istio.io/v1alpha1
kind: Telemetry
metadata:
  name: test
spec:
  tracing:
  - customTags:
      my_new_foo_tag:
        literal:
          value: foo
    providers:
    - name: skywalking
    randomSamplingPercentage: 100
EOF