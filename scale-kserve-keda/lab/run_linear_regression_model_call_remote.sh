#!/bin/bash
curl http://linear-regression-model-predictor-default.kserve-test.127.0.0.1.nip.io/v1/models/linear-regression-model:predict -d @./replicas-input.json
