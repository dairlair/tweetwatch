package main

import (
	"errors"
	"fmt"
	"github.com/dairlair/tweetwatch/pkg/broadcaster/providers/nats"
	"strings"

	"github.com/dairlair/tweetwatch/pkg/cmd/server"
	"github.com/dairlair/tweetwatch/pkg/storage"
	"github.com/dairlair/tweetwatch/pkg/twitterclient"
	"github.com/dairlair/tweetwatch/pkg/twitterclient/providers/gotwitter"
	"github.com/dairlair/tweetwatch/pkg/twitterclient/providers/void"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	log.SetLevel(log.WarnLevel)
}

func main() {
	config, providers, err := readConfig()
	if err != nil {
		log.Errorf("tweetwatch configurations failed: %s", err)
		return
	}
	setLogLevel(config.LogLevel)
	log.Infof("config: %v\n", config)
	srv := server.NewInstance(&config, providers)
	srv.Start()
}

func setLogLevel(logLevel string) {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Warningf("Unknown log level [%s]", logLevel)
	} else {
		log.SetLevel(level)
	}
}

func readConfig() (server.Config, server.Providers, error) {
	configureViper()

	err := viper.ReadInConfig()
	if err != nil {
		log.Warnf("config file not read: %s", err)
	}

	twitterProviderName := viper.GetString("twitter.provider")
	twitterProvider, err := getTwitterProvider(twitterProviderName)
	if err != nil {
		return server.Config{}, server.Providers{}, err
	}

	providers := server.Providers{
		TwitterClientProvider: twitterProvider,
	}

	return server.Config{
		LogLevel: viper.GetString("server.logLevel"),
		Postgres: storage.PostgresConfig{
			DSN: viper.GetString("postgres.dsn"),
		},
		RESTListen: viper.GetInt("rest.port"),
		Twitterclient: twitterclient.Config{
			TwitterConsumerKey:    viper.GetString("twitter.consumerKey"),
			TwitterConsumerSecret: viper.GetString("twitter.consumerSecret"),
			TwitterAccessToken:    viper.GetString("twitter.accessToken"),
			TwitterAccessSecret:   viper.GetString("twitter.accessSecret"),
		},
		NATS: nats.Config{
			URL:       viper.GetString("nats.url"),
			ClusterID: viper.GetString("nats.clusterId"),
			ClientID:  viper.GetString("nats.clientId"),
		},
	}, providers, nil
}

func configureViper() {
	viper.AutomaticEnv()
	viper.SetConfigName("tweetwatch")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/tweetwatch")
	viper.AddConfigPath("$HOME/.tweetwatch")
	viper.AddConfigPath("./")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetDefault("rest.post", "1308")
	viper.SetDefault("twitter.provider", "go-twitter")
	viper.SetDefault("server.logLevel", "warning")
}

func getTwitterProvider(provider string) (server.TwitterClientProvider, error) {
	log.Infof("twitter provider: %s", provider)
	switch provider {
	case "void":
		return void.NewInstance, nil
	case "go-twitter":
		return gotwitter.NewInstance, nil
	}

	msg := fmt.Sprintf("unknown twitter provider [%s]. available values: \"go-twitter\", \"void\"\n", provider)
	return nil, errors.New(msg)
}