# gRPC-golang-sample
This is a very simple example of a golang server and client using gRPC to communicate.

## Requirements
- protoc compiler `sudo apt install -y protobuf-compiler`
- golang packages
```
go get github.com/golang/protobuf/protoc-gen-go
go get google.golang.org/grpc
```


## To compile to proto file
```bash
protoc --go_out=plugins=grpc:. mathlib.proto
```
