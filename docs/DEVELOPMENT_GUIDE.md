# Developers guide

## Run server locally

```shell
go run ./cmd/server/server.go
# Run some command after start
http post :1308/signup email=john@example.com password=secret
# Grab the token and get a topics list with JWT authorization
http :1308/topics "Authorization:${jwt}"
# Validate your creds via
http post :1308/login "Authorization:${jwt}"
# And create your first topic
http POST :1308/topics "Authorization:${jwt}" name="Tesla Inc." tracks:='["Tesla","Elon Musk"]'
http POST :1308/topics "Authorization:${jwt}" name="Tesla Inc." tracks:='["Trump"]'
http POST :1308/topics "Authorization:${jwt}" name="Disney" tracks:='["Mickey Mouse","Donald Duck"]'
# Get topics list after that
http :1308/topics "Authorization:Basic ${basic}"
```

# Swagger stubs regenerate

To regenerate swagger stubs run this command:

```shell
./tools/swagger.sh
```

### Mockery mocks

To regenerate mock used in tests (i.g.: /pkg/storage/mocks) run this command:./c

```shell
./tools/regenerate-mocks.sh
```