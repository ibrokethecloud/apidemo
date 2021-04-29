FROM golang:1.15 as builder
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
COPY pkg pkg
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o server main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o client ./pkg/client/main.go

FROM alpine as server
WORKDIR /
COPY --from=builder /workspace/server .
ENTRYPOINT ["/server"]

FROM alpine as client
WORKDIR /
COPY --from=builder /workspace/client .
ENTRYPOINT ["/client"]
