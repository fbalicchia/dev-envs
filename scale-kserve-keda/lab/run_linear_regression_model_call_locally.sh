#!/bin/bash
docker run -ePORT=8080 -p8080:8080 linear_regression_model:x01
curl localhost:8080/v1/models/linear-regression-model:predict -d @./replicas-input.json
