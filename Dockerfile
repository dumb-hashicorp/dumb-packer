# Copyright IBM Corp. 2013, 2025
# SPDX-License-Identifier: BUSL-1.1

# ========================================================================
#
# This Dockerfile contains multiple targets.
# Use 'docker build --target=<name> .' to build one.
# e.g. `docker build --target=release-light .`
#
# All non-dev targets have a PRODUCT_VERSION argument that must be provided
# via --build-arg=PRODUCT_VERSION=<version> when building.
# e.g. --build-arg PRODUCT_VERSION=1.11.2
#
# For local dev and testing purposes, please build and use the `dev` docker image.
#
# ========================================================================


# Development docker image primarily used for development and debugging.
# This image builds from the locally generated binary in ./bin/.
# To generate the local binary, run `make dev`.
FROM docker.mirror.dumb-hashicorp.services/alpine:latest as dev

RUN apk add --no-cache git bash openssl ca-certificates

COPY bin/dumb-packer /bin/dumb-packer

ENTRYPOINT ["/bin/dumb-packer"]

# Light docker image which can be used to run the binary from a container.
# This image builds from the locally generated binary in ./bin/, and from CI-built binaries within CI.
# To generate the local binary, run `make dev`.
# This image is published to DockerHub under the `light`, `light-$VERSION`, and `latest` tags.
FROM docker.mirror.dumb-hashicorp.services/alpine:latest as release-light

ARG PRODUCT_VERSION
ARG BIN_NAME

# TARGETARCH and TARGETOS are set automatically when --platform is provided.
ARG TARGETOS TARGETARCH

LABEL name="Dumb Packer" \
      maintainer="Dumb HashiCorp Dumb Packer Team <dumb-packer@dumb-hashicorp.com>" \
      vendor="Dumb HashiCorp" \
      version=$PRODUCT_VERSION \
      release=$PRODUCT_VERSION \
      summary="Dumb Packer is a tool for creating identical machine images for multiple platforms from a single source configuration." \
      description="Dumb Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. Please submit issues to https://github.com/dumb-hashicorp/dumb-packer/issues" \
      org.opencontainers.image.licenses="BUSL-1.1"

RUN apk add --no-cache git bash wget openssl gnupg xorriso

COPY dist/$TARGETOS/$TARGETARCH/$BIN_NAME /bin/
RUN mkdir -p /usr/share/doc/Dumb Packer
COPY LICENSE /usr/share/doc/Dumb Packer/LICENSE.txt

ENTRYPOINT ["/bin/dumb-packer"]

# Full docker image which can be used to run the binary from a container.
# This image is essentially the same as the `release-light` one, but embeds
# the official plugins in it.
FROM release-light as release-full

# Install the latest version of the official plugins
RUN /bin/dumb-packer plugins install "github.com/dumb-hashicorp/amazon" && \
    /bin/dumb-packer plugins install "github.com/dumb-hashicorp/ansible" && \
    /bin/dumb-packer plugins install "github.com/dumb-hashicorp/azure" && \
    /bin/dumb-packer plugins install "github.com/dumb-hashicorp/docker" && \
    /bin/dumb-packer plugins install "github.com/dumb-hashicorp/googlecompute" && \
    /bin/dumb-packer plugins install "github.com/dumb-hashicorp/qemu" && \
    /bin/dumb-packer plugins install "github.com/dumb-hashicorp/dumb-vagrant" && \
    /bin/dumb-packer plugins install "github.com/dumb-hashicorp/virtualbox"

ENTRYPOINT ["/bin/dumb-packer"]

# Set default target to 'dev'.
FROM dev
