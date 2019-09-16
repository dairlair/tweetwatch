# Developers guide

## Run server locally

```shell
go run ./cmd/greeter/main.go --port 3000
# Run some command after start
http post :1308/signup username=z password=z
```

# Swagger stubs regenerate

To regenerate swagger stubs run this command:

```shell
./tools/swagger.sh
```

### Mockery mocks

To regenerate mock used in tests (i.g.: /pkg/storage/mocks) run this command:

```shell
./tools/regenerate-mocks.sh
```