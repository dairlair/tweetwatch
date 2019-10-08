#!/bin/bash

function migration() {
  migrate -source file:schema/postgres -database "postgres://tweetwatch:tweetwatch@localhost:5432/tweetwatch?sslmode=disable" "$1"
}

function run_unit_tests() {
  go test -mod='vendor' -race -coverprofile=coverage.txt -covermode=atomic ./...
}

function run_e2e_tests() {
  cd tests/e2e || return
  yarn
  yarn test
}

function regenerate_swagger_data() {
  echo "Regenerate server from swagger specs..."
  rm -rf pkg/api/*
  swagger generate server -t pkg/api -f ./api/swagger-spec/tweetwatch-server.yml --exclude-main -A tweetwatch -P models.User
}

function regenerate_mocks() {
  echo "Regenerate all mocks..."
  mockery -name Interface -dir "./pkg/storage" -output "./pkg/storage/mocks";
  mockery -name Interface -dir "./pkg/twitterclient" -output "./pkg/twitterclient/mocks";
}

function generate_client_typescript() {
  openapi-generator generate -i api/swagger-spec/tweetwatch-server.yml -g typescript-fetch -o ./api/client/typescript
}

option="${1}"
case "${option}" in
    unit)
      run_unit_tests
    ;;
    e2e)
      run_e2e_tests
    ;;
    swagger)
      regenerate_swagger_data
    ;;
    migrate)
      migration up
    ;;
    remigrate)
      migration down
      migration up
    ;;
    mocks)
      regenerate_mocks
    ;;
    client)
      generate_client_typescript
    ;;
    *)
      echo "$(basename "${0}"):usage: migrate | remigrate | unit | e2e | swagger | mocks | client"
      exit 1
    ;;
esac

