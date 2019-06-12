package storage

import (
	"sync"

	pb "github.com/dairlair/twitwatch/pkg/api/v1"
	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
)

// Interface must be implemented by postgres based storage or something else.
type Interface interface {
	AddStream(stream *pb.Stream) (id int64, err error)
	AddTwit(twit *pb.Twit) (id int64, err error)
}

// NewStorage creates new Storage instance
func NewStorage(connPool *pgx.ConnPool) *Storage {
	return &Storage{
		connPool: connPool,
	}
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
