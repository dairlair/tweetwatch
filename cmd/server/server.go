package main

import (
	"strings"

	"github.com/dairlair/twitwatch/pkg/cmd/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	grpcServer "github.com/dairlair/twitwatch/pkg/protocol/grpc"
)

func main() {
	config := readConfig()
	log.Infof("Config: %v\n", config)
	// Here we will run server... Coming soon
	srv := server.NewInstance(&config)
	srv.Start()
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
		GRPC: grpcServer.Config{
			ListenAddress: viper.GetString("grpc.listen"),
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
