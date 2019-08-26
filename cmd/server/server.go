package main

import (
	"strings"

	"github.com/dairlair/tweetwatch/pkg/cmd/server"
	grpcServer "github.com/dairlair/tweetwatch/pkg/protocol/grpc"
	"github.com/dairlair/tweetwatch/pkg/storage"
	"github.com/dairlair/tweetwatch/pkg/twitterclient"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	config := readConfig()
	log.Infof("config: %v\n", config)

	srv := server.NewInstance(&config)
	err := srv.Start()
	if err != nil {
		log.Errorf("tweetwatch start failed: %s", err)
	}
}

func readConfig() server.Config {
	configureViper()

	err := viper.ReadInConfig()
	if err != nil {
		log.Warnf("config file not read: %s", err)
	}

	twitterProvider := viper.GetString("twitter.provider")
	log.Infof("twitter provider: %s", twitterProvider)

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
	viper.SetDefault("grpc.listen", ":1308")
	viper.SetDefault("twitter.provider", "go-twitter")
}

func injectTwitterClient(provider String) {

}