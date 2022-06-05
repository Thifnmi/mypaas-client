FROM golang:1.16-alpine AS builder

COPY . /go/src/mypaas-client

WORKDIR /go/src/mypaas-client

RUN apk add --update gcc git make musl-dev && \
    make build

FROM alpine:3.8

COPY --from=builder /go/src/mypaas-client/bin/tsuru /bin/tsuru

RUN apk update && \
    apk add --no-cache ca-certificates && \
    rm /var/cache/apk/*

CMD ["tsuru"]
