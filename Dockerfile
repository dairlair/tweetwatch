########################################################################################################################
# Build server binary stage
########################################################################################################################
FROM golang:latest AS builder
ADD . /tweetwatch
WORKDIR /tweetwatch
# We run go build with flag "-mod vendor" to use vendor version of packages stored in the repository.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -a -o /tweetwatch-server cmd/server/server.go

########################################################################################################################
# Final stage                                                                                                          #
########################################################################################################################
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /tweetwatch-server /tweetwatch/server
RUN chmod +x /tweetwatch/server
# These files are required to migrate database
COPY --from=builder /tweetwatch/schema /tweetwatch/schema
COPY --from=builedr /tweetwatch/tools/migrate.linux-amd64 /tweetwatch/migrate
RUN chmod +x /tweetwatch/migrate
ENTRYPOINT ["/tweetwatch/server"]
EXPOSE 1308