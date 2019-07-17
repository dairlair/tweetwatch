package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
)

func (suite StorageSuite) TestAddStream_Successful() {
	id, err := suite.storage.AddStream(&entity.Stream{Track: "something"})
	suite.True(id > 0)
	suite.Equal(err, nil)
}
