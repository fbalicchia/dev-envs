#!/usr/bin/env bash
IMG_NAME=linear_regression_model:x01
docker build --no-cache  -f linear_regression_model.Dockerfile -t localhost:5001/${IMG_NAME} .
docker push localhost:5001/${IMG_NAME}

