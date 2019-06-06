package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	faker "github.com/bxcodec/faker/v3"
)

func TestEnvConfigurationReading(t *testing.T) {
	postgresDsn := faker.URL()

	os.Setenv("POSTGRES_DSN", postgresDsn)

	config := readConfig()

	assert.Equal(t, postgresDsn, config.Postgres.DSN)
}
