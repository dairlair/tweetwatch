package storage

import (
	pb "github.com/dairlair/tweetwatch/pkg/api/v1"
)

func (suite StorageSuite) TestAddTwit_Successful() {
	twit := pb.Twit{
		CreatedAt: "2019-01-01 00:00:00",
	}
	streamIds := []int64{1, 2}
	id, err := suite.storage.AddTwit(&twit, streamIds)
	suite.True(id > 0)
	suite.Equal(err, nil)
}
