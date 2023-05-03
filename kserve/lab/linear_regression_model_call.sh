#!/bin/bash

curl localhost:8080/v1/models/linear-regression-model:predict -d @./replicas-input.json
#curl localhost:8080/v1/models/linear-regression-model:predict -d @./replicas-input.json