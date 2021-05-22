ARG GO_VERSION=1.16-alpine3.12
ARG FROM_IMAGE=alpine:3.12

FROM golang:${GO_VERSION} AS builder

ARG VERSION

LABEL org.opencontainers.image.source="https://github.com/omegion/ssh-manager"

WORKDIR /app

COPY ./ /app

RUN apk update && \
  apk add ca-certificates gettext git make curl unzip && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

RUN make build-for-container VERSION=$VERSION

FROM ${FROM_IMAGE}

COPY --from=builder /app/dist/ssh-manager-linux /bin/ssh-manager

ENTRYPOINT ["ssh-manager"]
