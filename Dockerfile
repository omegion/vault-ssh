ARG GO_VERSION=1.16-alpine3.12
ARG FROM_IMAGE=alpine:3.14.3

FROM --platform=${BUILDPLATFORM} golang:${GO_VERSION} AS builder

ARG TARGETOS
ARG TARGETARCH

LABEL org.opencontainers.image.source="https://github.com/omegion/vault-ssh"

WORKDIR /app

COPY ./ /app

RUN apk update && \
  apk add ca-certificates gettext git make && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

RUN make build TARGETOS=$TARGETOS TARGETARCH=$TARGETARCH

FROM ${FROM_IMAGE}

COPY --from=builder /app/dist/vault-ssh /bin/vault-ssh

ENTRYPOINT ["vault-ssh"]
