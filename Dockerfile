FROM golang:1.19.4

WORKDIR /usr/src/rancher-redeploy-workload

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .

RUN mkdir bin

RUN CGO_ENABLED=0 go build -a -o bin/rancher-redeploy-workload .

FROM alpine:3.17

RUN apk update && apk add bash

COPY --from=0 /usr/src/rancher-redeploy-workload/bin/rancher-redeploy-workload /usr/local/bin/rancher-redeploy-workload

COPY docker-entrypoint.sh /usr/local/bin/

CMD ["/usr/local/bin/docker-entrypoint.sh"]
