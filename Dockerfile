FROM golang:1.16

WORKDIR /usr/src/rancher-redeploy-workload

COPY go.mod go.sum ./
RUN go mod download

COPY config.go .
COPY main.go .
COPY validator.go .

RUN mkdir bin

RUN CGO_ENABLED=0 go build -a -o bin/rancher-redeploy-workload .

FROM alpine:latest

RUN apk update && apk add bash

COPY --from=0 /usr/src/rancher-redeploy-workload/bin/rancher-redeploy-workload /usr/local/bin/rancher-redeploy-workload

COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

CMD ["/usr/local/bin/docker-entrypoint.sh"]
