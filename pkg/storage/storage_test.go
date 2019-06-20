package storage

import (
	"os"
	"testing"

	pb "github.com/dairlair/twitwatch/pkg/api/v1"
	"github.com/stretchr/testify/suite"
)

type storageHandlerSuite struct {
	StorageSuite
}

func TestStorageSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test for storage")
	}
	cfg := PostgresConfig{
		DSN: os.Getenv("POSTGRES_DSN"),
	}
	storageHandlerSuiteTest := &storageHandlerSuite{
		NewStorageSuite(cfg),
	}
	suite.Run(t, storageHandlerSuiteTest)
}

func (suite storageHandlerSuite) TestAddStream_Successful() {
	storage := NewStorage(suite.StorageSuite.cfg)

	id, err := storage.AddStream(&pb.Stream{Track: "something"})
	suite.True(id > 0)
	suite.Equal(err, nil)
}
