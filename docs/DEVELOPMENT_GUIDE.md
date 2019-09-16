# Developers guide

### gRPC proto updates
To regenerate gRPC service from updated proto files (located in /api/proto) run this command:
```shell
./third_party/protoc-gen.sh
```

For more information see https://github.com/golang/protobuf#installation

### Mockery mocks

To regenerate mock used in tests (i.g.: /pkg/storage/mocks) run this command:

```shell
./tools/regenerate-mocks.sh
```

#### Dial server though grpcurl

List of services
```shell
grpcurl -proto api/proto/v1/twitwatch-service.proto localhost:1308 list
```

List of service methods
```shell
grpcurl -proto api/proto/v1/twitwatch-service.proto localhost:1308 list v1.TwitwatchService
```

Create stream
```shell
grpcurl -plaintext -proto api/proto/v1/twitwatch-service.proto -d '{"api": "v1", "stream": {"track": "Tesla"}}' localhost:1308 v1.TwitwatchService.CreateStream
```

Get streams
```shell
grpcurl -plaintext -proto api/proto/v1/twitwatch-service.proto -d '{"api": "v1"}' localhost:1308 v1.TwitwatchService.GetStreams
```

Sign up
```shell
grpcurl -plaintext -proto api/proto/v1/twitwatch-service.proto -d '{"api": "v1", "email": "john.doe@example.com", "password": "secret"}' localhost:1308 v1.TwitwatchService.SignUp
```

Sign in
```shell
grpcurl -plaintext -proto api/proto/v1/twitwatch-service.proto -d '{"api": "v1", "email": "john.doe@example.com", "password": "secret"}' localhost:1308 v1.TwitwatchService.SignIn
```