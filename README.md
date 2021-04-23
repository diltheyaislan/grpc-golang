# grpc-golang

gRPC client and server communication using Golang 

## Starting server

```console
go run cmd/server/server.go
```

## Starting client

```console
go run cmd/client/client.go
```

## Compiling Protocol Buffers

```console
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```