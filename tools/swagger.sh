#!/bin/bash
echo "Regenerate server from swagger specs..."
swagger generate server -t pkg/api -f ./api/swagger-spec/tweetwatch-server.yml --exclude-main -A greeter