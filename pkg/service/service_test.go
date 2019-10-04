package service

import (
	. "github.com/dairlair/tweetwatch/pkg/entity"
	storageMock "github.com/dairlair/tweetwatch/pkg/storage/mocks"
	twitterclientMock "github.com/dairlair/tweetwatch/pkg/twitterclient/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestEnvConfigurationReading(t *testing.T) {
	tc := createTwitterclientMock()
	tc.On("AddStreams", []StreamInterface{}).Return()
	tc.On("Start").Return(nil)
	tc.On("Watch", mock.AnythingOfType("chan entity.TweetStreamsInterface")).Return(nil)
	s := createStorageMock()
	s.On("GetStreams").Return([]StreamInterface{}, nil)
	srv := NewService(s, tc)
	assert.NotNil(t, srv)
}

func createTwitterclientMock() *twitterclientMock.Interface {
	return &twitterclientMock.Interface{}
}

func createStorageMock() *storageMock.Interface {
	return &storageMock.Interface{}
}