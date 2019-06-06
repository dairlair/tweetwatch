package main

import (
	"fmt"
	"github.com/dairlair/twitwatch/pkg/cmd/server"
	"github.com/spf13/viper"
	"strings"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := readConfig()
	fmt.Printf("Postgres DSN: %s\n", config.Postgres.DSN)
}

func readConfig() server.Config {
	configureViper()

	err := viper.ReadInConfig()
	if err != nil {
		log.Warnf("config file not read: %s", err)
	}

	return server.Config{
		Postgres: server.PostgresConfig{
			DSN: viper.GetString("postgres.dsn"),
		},
	}
}

func configureViper() {
	viper.AutomaticEnv()
	viper.SetConfigName("twitwatch")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/twitwatch")
	viper.AddConfigPath("$HOME/.twitwatch")
	viper.AddConfigPath("./")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}