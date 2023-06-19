PORT ?= 8000

build:
	docker build . -t envoy-lua-test

run:
	docker run -p 8080:8080 -e PORT=8080 --rm -t solsson/http-echo
	docker run --rm -it -v $(PWD)/envoy.yaml:/etc/envoy/envoy.yaml -p 8000:8000 -t envoy-lua-test -c /etc/envoy/envoy.yaml
