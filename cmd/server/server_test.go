package main

import (
	"github.com/bxcodec/faker/v3"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvConfigurationReading(t *testing.T) {
	postgresDsn := faker.URL()
	err := os.Setenv("POSTGRES_DSN", postgresDsn)
	assert.Nil(t, err, "Error when tried to set environment variable POSTGRES_DSN")
	config, _, err := readConfig()
	assert.Nil(t, err)
	assert.Equal(t, postgresDsn, config.Postgres.DSN)
}