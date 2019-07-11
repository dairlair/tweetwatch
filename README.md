# TwitWatch
A Twitter API based daemon for twits analyses purposes

## Developers guide

Apply migrations
```shell
migrate -source file:schema/postgres -database "postgres://twitwatch:twitwatch@localhost:5432/twitwatch_test?sslmode=disable" up
```

To regenerate gRPC service from updated proto files (located in /api/proto) run this command:
```shell
./third_party/protoc-gen.sh
```

To regenerate mock used in tests (i.g.: /pkg/storage/mocks) run this command:
```shell
cd pkg/storage
mockery -name Interface # Mock type Interface and save generated file into the "mocks" subdirectory
```

For more information see https://github.com/golang/protobuf#installation

Run daemon locally
```shell
go run cmd/server/server.go
```

#### Dial server though grpcurl
List of services
```shell
grpcurl -proto api/proto/v1/twitwatch-service.proto localhost:1234 list
```

List of service methods
```shell
grpcurl -proto api/proto/v1/twitwatch-service.proto localhost:1234 list v1.TwitwatchService
```

Create stream
```shell
grpcurl -plaintext -proto api/proto/v1/twitwatch-service.proto -d '{"api": "v1", "stream": {"track": "Tesla"}}' localhost:1234 v1.TwitwatchService.CreateStream
```

Get streams
```shell
grpcurl -plaintext -proto api/proto/v1/twitwatch-service.proto -d '{"api": "v1"}' localhost:1234 v1.TwitwatchService.GetStreams
```