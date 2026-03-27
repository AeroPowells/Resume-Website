FROM registry.access.redhat.com/ubi9/ubi:latest

# UBI9 ships curl-minimal which conflicts with curl — --allowerasing replaces it
RUN dnf install -y --allowerasing \
    git \
    curl \
    wget \
    tar \
    gcc \
    make \
    && dnf clean all

# Install Go — detect arch so this works on both x86_64 and Apple Silicon (aarch64)
ARG GO_VERSION=1.23.4
RUN ARCH=$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/') \
    && curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz \
    | tar -C /usr/local -xz
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/home/developer/go"
ENV PATH="${GOPATH}/bin:${PATH}"

# Install Node.js — use .tar.gz (UBI9 doesn't have xz installed)
ARG NODE_VERSION=20.18.1
RUN ARCH=$(uname -m | sed 's/x86_64/x64/;s/aarch64/arm64/') \
    && curl -fsSL https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-linux-${ARCH}.tar.gz \
    | tar -C /usr/local -xz --strip-components=1

# Set up a non-root user
ARG USERNAME=developer
ARG USER_UID=1000
ARG USER_GID=${USER_UID}
RUN groupadd --gid ${USER_GID} ${USERNAME} \
    && useradd --uid ${USER_UID} --gid ${USER_GID} -m ${USERNAME} \
    && chown -R ${USERNAME}:${USERNAME} /workspace 2>/dev/null || true

WORKDIR /workspace
USER ${USERNAME}
