package storage

import (
	"errors"
	"fmt"
	"sync"

	pb "github.com/dairlair/twitwatch/pkg/api/v1"
	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
)

// Interface must be implemented by postgres based storage or something else.
type Interface interface {
	AddStream(stream *pb.Stream) (id int64, err error)
	GetStreams() (streams []*pb.Stream, err error)
	AddTwit(twit *pb.Twit) (id int64, err error)
}

// NewStorage creates new Storage instance
func NewStorage(connPool *pgx.ConnPool) *Storage {
	return &Storage{
		connPool: connPool,
	}
}

// @TODO Add method to graceful shutdown with "defer s.connPool.Close()" and all other things

// PostgresConfig contains options for the Postgres connections pool
type PostgresConfig struct {
	DSN string
}

// Storage provides ability to store and to retrieve twits and other entities
type Storage struct {
	mutex    sync.RWMutex
	connPool *pgx.ConnPool
}

func pgError(err error) error {
	log.Error(err)
	return err
}

func CreatePostgresConnection(config PostgresConfig) *pgx.ConnPool {
	pgConf, err := pgx.ParseURI(config.DSN)
	if err != nil {
		msg := fmt.Sprintf("Can not parse Postgres DSN: %s", err)
		panic(errors.New(msg))
	}
	log.Infof("PostgreSQL: host=%s, port=%d, username=%s, database=%s",
		pgConf.Host,
		pgConf.Port,
		pgConf.User,
		pgConf.Database,
	)
	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     pgConf.Host,
			Port:     pgConf.Port,
			User:     pgConf.User,
			Password: pgConf.Password,
			Database: pgConf.Database,
			RuntimeParams: map[string]string{
				"application_name": "twitwatch",
			},
		},
		MaxConnections: 1,
	})

	if err != nil {
		msg := fmt.Sprintf("Can not connect to Postgres: %s", err)
		panic(errors.New(msg))
	}

	return connPool
}
