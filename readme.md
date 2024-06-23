# gowasm

WASM isn't just for the web, it can be used for standalone applications too, as long as there is a runtime for it - of which there are now plenty:

- io.containerd.slight.v1
- io.containerd.spin.v2
- io.containerd.wasmedge.v1
- io.containerd.wasmtime.v1
- io.containerd.lunatic.v1
- io.containerd.wws.v1
- io.containerd.wasmer.v1

(https://docs.docker.com/desktop/wasm/)

And WebAssembly is closer to the hardware so should be more efficient to run regular code than virtualisation and containerisation (by this I mean container runtime, not packaging things into images).

So, this is an experiment to see how much more efficient a program can be in WASM on Kubernetes vs. an image running on a regular container runtime.

My aim is to keep as much as similar as possible between a WASM run and a containerised run. The obvious difference would be how the binary is created inside the container image, and the runtime itself. Other than that the two runs should be as similar as possible.

## Pre-requisites:

First of all, to build and run Golang programs that are compatible with WASM I need at least Golang 1.21.

Also for the WASM aspect I'll need a WASM runtime. I've opted for wasmtime from bytecode alliance. Others are available as discussed and maybe i'll extend this trial to do a comparison of each.

### [wasmtime](https://github.com/bytecodealliance/wasmtime)
`curl https://wasmtime.dev/install.sh -sSf | bash`

There are also some configuration changes I need to make in Docker Desktop for local testing:

- Enable containerd image pulling in Docker Settings -> General. (Restart Docker)
- Enable WASM in beta features (Restart Docker)

I will also need to make some changes to Kubernetes (i'm using AKS):

https://learn.microsoft.com/en-us/azure/aks/use-wasi-node-pools

## Main

First off, i'll need to make a Golang module that will support WASM:

https://go.dev/blog/wasi

`go mod init github.com/aykay76/gowasm`

.. create main.go ..

`GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go`

`wasmtime main.wasm`

## Containerise

I'll need to build an OCI image for the WASM application:

https://docs.docker.com/desktop/wasm/

.. Create Dockerfiles .. 

`docker build -t gowasm:wasm -f ./Dockerfilewasm .`
`docker build -t gowasm:go -f ./Dockerfilego .`

## Run

Now I can test running the image locally, on a WASM runtime:

`docker run --runtime=io.containerd.wasmtime.v1 gowasm:wasm`

.. and a normal container runtime

`docker run gowasm:go`