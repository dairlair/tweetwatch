# Operations guide

## Apply migrations command:

Command migrate docs is available [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

```shell
migrate -source file:schema/postgres -database "postgres://tweetwatch:tweetwatch@localhost:5432/tweetwatch?sslmode=disable" up
```

## Run daemon locally
When configured through config file (tweetwatch.yml)
```shell
go run cmd/server/server.go
```

...or with config through environment variables:

```shell
SERVER_LOG_LEVEL=warning
POSTGRES_DSN="postgres://tweetwatch:tweetwatch@localhost:5432/tweetwatch?sslmode=disable" \
GRPC_LISTEN=":1308" \
TWITTER_CONSUMER_KEY="SOME_TWITTER_CONSUMER_KEY" \
TWITTER_CONSUMER_SECRET="SOME_TWITTER_CONSUMER_KEY" \
TWITTER_ACCESS_TOKEN="SOME_TWITTER_ACCESS_TOKEN" \
TWITTER_ACCESS_SECRET="SOME_TWITTER_ACCESS_SECRET" \
go run cmd/server/server.go
```