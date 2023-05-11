HUB ?= localhost:5001/keda-external-scaler
VERSION ?= x01

build:
	docker build . -t $(HUB):$(VERSION)
push-image:
	docker push $(HUB):$(VERSION)

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...
