package storage

import (
	"fmt"
	. "github.com/dairlair/tweetwatch/pkg/entity"
	"time"
)

func (suite StorageSuite) TestAddStream_Successful() {
	email := fmt.Sprintf("tester%d@example.com", time.Now().Second())
	userId, err := suite.storage.SignUp(email, "secret")
	suite.Nil(err, "User must be created successfully")

	topic, err := suite.storage.AddTopic(&Topic{
		UserID:    *userId,
		Name:      fmt.Sprintf("Test topic %d", time.Now().Second()),
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
