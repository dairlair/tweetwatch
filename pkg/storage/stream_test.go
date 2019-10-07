package storage

import (
	. "github.com/dairlair/tweetwatch/pkg/entity"
)

func (suite StorageSuite) TestAddStream_Successful() {
	id, err := suite.storage.AddTweetStreams(NewTweetStreams(&Tweet{
		ID:            1,
		TwitterID:     2,
		TwitterUserID: 3,
		FullText:      "Something...",
		CreatedAt:     "2019-02-03 12:12:12",
	}, []StreamInterface{&Stream{
		ID:      1,
		TopicID: 2,
		Track:   "Test",
	}}))
	suite.NotNil(id)
	suite.Equal(nil, err)
}
