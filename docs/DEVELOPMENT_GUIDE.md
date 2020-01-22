# Developers guide

## Run server locally

```shell
go run ./cmd/server/server.go
# Check the CORS headers
curl "http://localhost:1308/login" -X OPTIONS -H "Sec-Fetch-Mode: cors" -H "Access-Control-Request-Method: POST" -H "Origin: http://localhost:3000" -H "Access-Control-Request-Headers: content-type" --compressed -i
# Signup after start
http post :1308/signup email=john@example.com password=secret
# Or login
http post :1308/login email=john@example.com password=secret
# Grab the token and get a topics list with JWT authorization
http :1308/topics "Authorization:${jwt}"
# Validate your creds via
http post :1308/login "Authorization:${jwt}"
# And create your first topic
http POST :1308/topics "Authorization:${jwt}" name="Tesla Inc." isActive:=true
# Get topics list after that
http :1308/topics "Authorization:${jwt}"
# Update created topic
http PUT :1308/topics/1 "Authorization:${jwt}" name="Tesla Inc." isActive:=false
# Delete topic
http DELETE :1308/topics/1 "Authorization:${jwt}"
# Add stream to some topic
http POST :1308/topics/1/streams "Authorization:${jwt}" track="qwerty"
# Get streams from some topic
http :1308/topics/1/streams "Authorization:${jwt}"
# Update created stream
http PUT :1308/topics/1/streams/1 "Authorization:${jwt}" track="zxcvbn"
# Get tweets of topic
http :1308/topics/1/tweets "Authorization:${jwt}"
```

### Swagger stubs regenerate

To regenerate swagger stubs run this command:

```shell script
./tools/commander swagger
```

### Mockery mocks

To regenerate mock used in tests (i.g.: /pkg/storage/mocks) run this command:./c

```shell script
./tools/commander mocks
```

### API Client generator

Install this package (more info on [github](https://github.com/OpenAPITools/openapi-generator)):
```shell script
npm install @openapitools/openapi-generator-cli -g
```

then run:
```shell script
./tools/commander client
```