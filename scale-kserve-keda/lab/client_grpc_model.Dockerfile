FROM python:3.9-slim-bullseye


COPY ./kservelab/grpc_client.py /home/kserve/
COPY ./replicas-input.json /home/kserve/input.json
ENV INGRESS_HOST=value1
ENV INGRESS_PORT=8081
ENV MODEL_NAME=linear-regression-model-grpc

RUN  pip install --no-cache-dir kserve
RUN useradd kserve -m -u 1000 -d /home/kserve
