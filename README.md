# TwitWatch
A Twitter API based daemon for twits analyses purposes

## Developers guide

To regenerate gRPC service from updated proto files (located in /api/proto) run this command:
```shell
./third_party/protoc-gen.sh
```

Run daemon locally
```shell
go run cmd/server/server.go
```

For more information see https://github.com/golang/protobuf#installation

## Operation guide

TBD