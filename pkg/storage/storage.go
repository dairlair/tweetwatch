package storage

import (
	"sync"

	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
)

// NewStorage creates new Storage instance
func NewStorage(connPool *pgx.ConnPool) *Storage {
	return &Storage{
		connPool: connPool,
	}
}

// Storage provides ability to store and to retrieve twits and other entities
type Storage struct {
	mutex   sync.RWMutex
	connPool *pgx.ConnPool
}

func pgError(err error) error {
	log.Error(err)
	return err
}
