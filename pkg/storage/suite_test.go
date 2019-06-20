package storage

import (
	"github.com/jackc/pgx"
	"github.com/stretchr/testify/suite"
)

type StorageSuite struct {
	suite.Suite
	cfg      PostgresConfig
	connPool *pgx.ConnPool
}

func NewStorageSuite(cfg PostgresConfig) StorageSuite {
	return StorageSuite{cfg: cfg}
}

func (s *StorageSuite) SetupSuite() {
	s.connPool = createPostgresConnection(s.cfg)
}

func (s *StorageSuite) TearDownSuite() {
	s.connPool.Close()
}
