package storage

import (
	. "github.com/dairlair/tweetwatch/pkg/entity"
	"time"
)

func (suite StorageSuite) TestAddStream_Successful() {

	userId, err := suite.storage.SignUp("tester@example.com", "secret")
	suite.Nil(err, "User must be created successfully")

	topic, err := suite.storage.AddTopic(&Topic{
		UserID:    *userId,
		Name:      "Test topic",
		CreatedAt: time.Time{},
		IsActive:  true,
	})

	suite.Nil(err, "Topic must be created successfully")

	id, err := suite.storage.AddTweetStreams(NewTweetStreams(&Tweet{
		ID:            1,
		TwitterID:     2,
		TwitterUserID: 3,
		FullText:      "Something...",
		CreatedAt:     time.Now(),
	}, []StreamInterface{&Stream{
		ID:      1,
		TopicID: topic.GetID(),
		Track:   "Test",
	}}))
	suite.NotNil(id)
	suite.Equal(nil, err)
}
