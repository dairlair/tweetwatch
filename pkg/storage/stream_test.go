package storage

import (
	pb "github.com/dairlair/twitwatch/pkg/api/v1"
)

func (suite StorageSuite) TestAddStream_Successful() {
	id, err := suite.storage.AddStream(&pb.Stream{Track: "something"})
	suite.True(id > 0)
	suite.Equal(err, nil)
}
