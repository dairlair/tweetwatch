package main

import (
	"strings"

	"github.com/dairlair/twitwatch/pkg/cmd/server"
	"github.com/dairlair/twitwatch/pkg/twitterclient"
	grpcServer "github.com/dairlair/twitwatch/pkg/protocol/grpc"
	"github.com/dairlair/twitwatch/pkg/storage"
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
		Postgres: storage.PostgresConfig{
			DSN: viper.GetString("postgres.dsn"),
		},
		GRPC: grpcServer.Config{
			ListenAddress: viper.GetString("grpc.listen"),
		},
		Twitterclient: twitterclient.Config{
			TwitterConsumerKey:    viper.GetString("twitter.consumerKey"),
			TwitterConsumerSecret: viper.GetString("twitter.consumerSecret"),
			TwitterAccessToken:    viper.GetString("twitter.accessToken"),
			TwitterAccessSecret:   viper.GetString("twitter.accessSecret"),
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
