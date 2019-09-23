# Developers guide

## Run server locally

```shell
go run ./cmd/server/server.go
# Run some command after start
http post :1308/signup email=john@example.com password=secret
# And login then...
basic=`echo "john@example.com:secret"|tr -d '\n'|base64 -i-`
http POST :1308/login "Authorization:Basic ${basic}"
# Or create your first topic
http POST :1308/topic "Authorization:Basic ${basic}" name="Tesla Inc." tracks:='["Tesla","Elon Musk"]'
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