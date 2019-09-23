#!/bin/bash
echo "Regenerate server from swagger specs..."
rm -rf pkg/api/*
swagger generate server -t pkg/api -f ./api/swagger-spec/tweetwatch-server.yml --exclude-main -A tweetwatch -P models.UserResponse