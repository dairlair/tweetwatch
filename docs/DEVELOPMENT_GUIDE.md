# Developers guide

## Run server locally

```shell
go run ./cmd/server/server.go
# Signup after start
http post :1308/signup email=john@example.com password=secret
# Or login
http post :1308/login email=john@example.com password=secret
# Grab the token and get a topics list with JWT authorization
http :1308/topics "Authorization:${jwt}"
# Validate your creds via
http post :1308/login "Authorization:${jwt}"
# And create your first topic
http POST :1308/topics "Authorization:${jwt}" name="Tesla Inc." tracks:='["Tesla","Elon Musk"]' isActive:=true
http POST :1308/topics "Authorization:${jwt}" name="Tesla Inc." tracks:='["Trump"]' isActive:=true
http POST :1308/topics "Authorization:${jwt}" name="Disney" tracks:='["Mickey Mouse","Donald Duck"]' isActive:=true
# Get topics list after that
http :1308/topics "Authorization:${jwt}"
# Update created topic
http PUT :1308/topics/1 "Authorization:${jwt}" name="Tesla Inc." tracks:='["BFR","Elon Musk"]' isActive:=true
# Check the CORS headers
 curl "http://localhost:1308/login" -X OPTIONS -H "Sec-Fetch-Mode: cors" -H "Access-Control-Request-Method: POST" -H "Origin: http://localhost:3000" -H "Access-Control-Request-Headers: content-type" --compressed -i
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