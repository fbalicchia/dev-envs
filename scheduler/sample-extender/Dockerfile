FROM golang:1.19-alpine as builder
ARG VERSION=0.0.1

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# build
WORKDIR /go/src/sample-scheduler-extender
COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on go mod download
COPY . .
RUN go build -ldflags "-s -w -X main.version=0.0.1" 

# runtime image
FROM gcr.io/google_containers/ubuntu-slim:0.14
COPY --from=builder /go/src/sample-scheduler-extender/sample-scheduler-extender sample-scheduler-extender
ENTRYPOINT ["sample-scheduler-extender"]
