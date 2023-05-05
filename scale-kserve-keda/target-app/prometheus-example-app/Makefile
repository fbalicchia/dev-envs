VERSION:=$(shell cat VERSION)
LDFLAGS="-X main.appVersion=$(VERSION)"
.PHONY: all
all: build docker-build docker-push install

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


##@ Build

.PHONY: build
build: ##build exec.
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o prometheus-example-app --installsuffix cgo main.go

.PHONY: docker-build
docker-build: ##build docker image with the manager.
	docker build -t localhost:5001/prometheus-example-app:$(VERSION) .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push localhost:5001/prometheus-example-app:$(VERSION)

.PHONY: install
install:  ## Install applications
	kubectl apply -f manifests/deployment.yaml;kubectl apply -f manifests/service-monitor.yaml
