# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.234.0/containers/python-3/.devcontainer/base.Dockerfile

# [Choice] Debian OS version (use bullseye on local amd64/Apple Silicon): buster, bullseye
ARG VARIANT="bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/rust:${VARIANT} AS build

RUN rustup default stable
RUN rustup target add wasm32-wasi

# TinyGo installation, see https://tinygo.org/getting-started/install/linux/ for instructions
ARG TINYGO_URL="https://github.com/tinygo-org/tinygo/releases/download/v0.30.0/tinygo0.30.0.linux-amd64.tar.gz"
RUN curl -sL "$TINYGO_URL" -o tinygo_amd64.tar.gz && \
    tar -xzf tinygo_amd64.tar.gz && \
    mv tinygo /usr/local/tinygo

RUN cargo install wit-bindgen-cli --git https://github.com/bytecodealliance/wit-bindgen --rev 1af7e87066854894ab140d2a630a0bc23c8b126f

## wasm-tools
RUN curl -LO https://github.com/bytecodealliance/wasm-tools/releases/download/wasm-tools-1.0.48/wasm-tools-1.0.48-x86_64-linux.tar.gz && \
    tar -xzf wasm-tools-1.0.48-x86_64-linux.tar.gz && \
    mv wasm-tools-1.0.48-x86_64-linux/wasm-tools /usr/local/bin

## spin
RUN curl -LO https://github.com/fermyon/spin/releases/download/v2.0.1/spin-v2.0.1-linux-amd64.tar.gz && \
    tar -xzf spin-v2.0.1-linux-amd64.tar.gz && \
    mv spin /usr/local/bin

## go
RUN curl -LO https://go.dev/dl/go1.21.4.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/usr/local/go"
RUN go install -v golang.org/x/tools/gopls@latest

FROM mcr.microsoft.com/vscode/devcontainers/rust:${VARIANT}
COPY --from=build /usr/local/bin /usr/local/bin
COPY --from=build /usr/local/go /usr/local/go
COPY --from=build /usr/local/cargo/bin /usr/local/cargo/bin
COPY --from=build /usr/local/tinygo /usr/local/tinygo
ENV PATH="/usr/local/go/bin:${PATH}"
ENV PATH="/usr/local/tinygo/bin:${PATH}"
ENV GOROOT="/usr/local/go"
ENV TINYGOROOT="/usr/local/tinygo"
