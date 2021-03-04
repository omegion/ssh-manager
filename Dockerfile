ARG GO_VERSION=1.15-alpine3.12
ARG FROM_IMAGE=alpine:3.11

FROM golang:${GO_VERSION} AS builder

LABEL org.opencontainers.image.source="https://github.com/omegion/bw-ssh"

RUN apk update && \
  apk add ca-certificates gettext git make && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

COPY ./ /app

WORKDIR /app

RUN make build-for-container

FROM ${FROM_IMAGE}

RUN apk update && \
  apk add ca-certificates gettext jq curl openssl git postgresql && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

COPY --from=builder /app/dist/bw-ssh-linux /bin/bw-ssh

ENTRYPOINT ["bw-ssh"]
