package service

import (
	storageMock "github.com/dairlair/tweetwatch/pkg/storage/mocks"
	twitterclientMock "github.com/dairlair/tweetwatch/pkg/twitterclient/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvConfigurationReading(t *testing.T) {
	tc := createTwitterclientMock()
	s := createStorageMock()
	srv := NewService(s, tc, nil)
	assert.NotNil(t, srv)
}

func createTwitterclientMock() *twitterclientMock.Interface {
	return &twitterclientMock.Interface{}
}

func createStorageMock() *storageMock.Interface {
	return &storageMock.Interface{}
}