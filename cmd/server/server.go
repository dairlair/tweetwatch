package main

import (
	"strings"

	"github.com/dairlair/twitwatch/pkg/cmd/server"
	grpcServer "github.com/dairlair/twitwatch/pkg/protocol/grpc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	config := readConfig()
	log.Infof("Config: %v\n", config)

	srv := server.NewInstance(&config)
	err := srv.Start()
	if err != nil {
		log.Errorf("twitwatch start failed: %s", err)
	}
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