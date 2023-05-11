FROM golang:1.19.3 as builder

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o external-scaler server.go


FROM alpine:latest

RUN addgroup -S fbalicchia && adduser -S fbalicchia -G fbalicchia &&  mkdir -p /fbalicchia 
RUN chown -R "fbalicchia:fbalicchia" /fbalicchia 
WORKDIR /fbalicchia

COPY --from=builder /src/external-scaler /fbalicchia/external-scaler
USER fbalicchia
ENTRYPOINT ["/fbalicchia/external-scaler"]