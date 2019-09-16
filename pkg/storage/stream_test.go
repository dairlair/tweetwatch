package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
)

func (suite StorageSuite) TestAddStream_Successful() {
	stream, err := suite.storage.AddStream(&entity.Stream{Track: "something"})
	suite.NotNil(stream)
	suite.True(stream.GetID() > 0)
	suite.Equal(err, nil)
}
