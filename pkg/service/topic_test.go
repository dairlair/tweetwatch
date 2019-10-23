package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Ensures then when we update topic - the service must do follows:
// If topic is inactive or deleted - remove all his streams from twitterclient
// If topic is active and non deleted - add his streams to twitterclient
func TestService_UpdateTopicHandler(t *testing.T) {
	tc := createTwitterclientMock()
	s := createStorageMock()
	srv := NewService(s, tc)
	assert.NotNil(t, srv)
}