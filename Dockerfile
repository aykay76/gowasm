FROM golang:1.22 AS build-stage
WORKDIR /
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go

FROM scratch
WORKDIR /
COPY --from=build-stage /main.wasm /main.wasm
ENTRYPOINT ["/main.wasm"]
