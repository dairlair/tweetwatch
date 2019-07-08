package storage

import (
	"os"
	"testing"

	"github.com/jackc/pgx"
	"github.com/stretchr/testify/suite"
)

type StorageSuite struct {
	suite.Suite
	cfg      PostgresConfig
	connPool *pgx.ConnPool
	storage  Interface
}

func NewStorageSuite(cfg PostgresConfig) StorageSuite {
	return StorageSuite{cfg: cfg}
}

func (suite *StorageSuite) SetupSuite() {
	suite.connPool = CreatePostgresConnection(suite.cfg)
	suite.storage = NewStorage(suite.connPool)
}

func (suite *StorageSuite) TearDownSuite() {
	suite.connPool.Close()
}

func TestStorageSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for storage")
	}
	cfg := PostgresConfig{
		DSN: os.Getenv("TWITWATCH_TEST_POSTGRES_DSN"),
	}
	storageSuite := NewStorageSuite(cfg)
	suite.Run(t, &storageSuite)
}
