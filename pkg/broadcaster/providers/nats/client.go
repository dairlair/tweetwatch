package nats

import (
	"fmt"
	"github.com/dairlair/tweetwatch/pkg/entity"
	stan "github.com/nats-io/go-nats-streaming"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	conn stan.Conn
}

func NewClient(config Config) *Client {
	conn, err := stan.Connect(config.ClusterID, config.ClientID, func(options *stan.Options) error {
		options.NatsURL = config.URL
		return nil
	})
	if err != nil {
		log.Errorf("Can't connect: %v. Make sure a NATS Streaming Server is running at: %s", err, config.URL)
		return nil
	}
	log.Infof("Connected to NATS Streaming cluster [%s] as client [%s]", config.ClusterID, config.ClientID)

	return &Client{
		conn: conn,
	}
}

func (c Client) Broadcast(channel string, tweetStreams entity.TweetStreamsInterface) {
	eventMsg := []byte(fmt.Sprintf("Tweet #%d", tweetStreams.GetTweet().GetTwitterID()))
	err := c.conn.Publish(channel, eventMsg)
	if err != nil {
		log.Errorf("Error publishing message %s: %v\n", err.Error())
	}
}