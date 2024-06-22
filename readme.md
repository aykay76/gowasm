# gowasm

## Pre-requisites:

### golang 1.21

### [wasmtime](https://github.com/bytecodealliance/wasmtime)
`curl https://wasmtime.dev/install.sh -sSf | bash`
https://docs.docker.com/engine/alternative-runtimes/#wasmtime

Enable containerd image pulling in Docker Settings -> General. (Restart Docker)
Enable WASM in beta features (Restart Docker)

## Main:

https://go.dev/blog/wasi

`go mod init github.com/aykay76/gowasm`

.. create main.go ..

`GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go`

`wasmtime main.wasm`

## Containerise

https://docs.docker.com/desktop/wasm/

.. Create Dockerfile .. 

`docker buildx build -t gowasm:latest .`

## Run

`docker run --runtime=io.containerd.wasmedge.v1 --platform=wasi/wasm gowasm:latest`
