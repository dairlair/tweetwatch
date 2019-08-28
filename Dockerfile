# Build server stage
FROM golang:latest AS builder
ADD . /tweetwatch
WORKDIR /tweetwatch
# We run go build with flag "-mod vendor" to use vendor version of packages stored in the repository.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -a -o /tweetwatch-server cmd/server/server.go

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /tweetwatch-server ./
RUN chmod +x ./tweetwatch-server
ENTRYPOINT ["./tweetwatch-server"]
EXPOSE 1308