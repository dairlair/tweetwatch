package storage

import (
	. "github.com/dairlair/tweetwatch/pkg/entity"
)

func (suite StorageSuite) TestAddStream_Successful() {
	id, err := suite.storage.AddTweetStreams(NewTweetStreams(&Tweet{}, []StreamInterface{&Stream{
		ID:      1,
		TopicID: 2,
		Track:   "Test",
	}}))
	suite.NotNil(id)
	suite.Equal(err, nil)
}
