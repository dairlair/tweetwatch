# TwitWatch
A Twitter API based daemon for twits analyses purposes

## Developers guide

Apply migrations
```shell
migrate -source file:schema/postgres -database "postgres://test:test@localhost:5432/twitwatch_test?sslmode=disable" up
```

To regenerate gRPC service from updated proto files (located in /api/proto) run this command:
```shell
./third_party/protoc-gen.sh
```

For more information see https://github.com/golang/protobuf#installation

Run daemon locally
```shell
go run cmd/server/server.go
```

Dial server though grpcurl
```shell
grpcurl -proto api/proto/v1/twitwatch-service.proto localhost:1234 list
```

```shell
grpcurl -proto api/proto/v1/twitwatch-service.proto localhost:1234 list v1.TwitwatchService
```