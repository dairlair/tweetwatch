package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/dairlair/tweetwatch/pkg/api/v1"
	"github.com/dairlair/tweetwatch/pkg/entity"
	storageMocks "github.com/dairlair/tweetwatch/pkg/storage/mocks"
	twitterclientMocks "github.com/dairlair/tweetwatch/pkg/twitterclient/mocks"
)

func TestCreateStream_RequestCreation(t *testing.T) {
	track := "something"
	req := pb.CreateStreamRequest{Stream: &pb.Stream{Track: track}}
	assert.Equal(t, track, req.GetStream().GetTrack(), "they should be equal")

}

func TestCreateStream_Successful(t *testing.T) {
	pbStream := pb.Stream{Track: "something"}
	entityStream := entity.Stream{Track: pbStream.GetTrack()}

	var id int64 = 1
	storageMock := storageMocks.Interface{}
	storageMock.On("AddStream", &entityStream).Return(id, nil)
	storageMock.On("GetStreams").Return(make([]entity.StreamInterface, 0))
	twitterclientMock := twitterclientMocks.Interface{}
	s := NewTweetwatchServiceServer(&storageMock, &twitterclientMock)

	req := pb.CreateStreamRequest{Stream: &pbStream}
	resp, err := s.CreateStream(context.Background(), &req)
	assert.Nil(t, err, "Error must be equal nil")
	assert.Equal(t, id, resp.Id, "Returned id mus be equal ID returned from storage")
}

func TestCreateStream_FailedOnStorage(t *testing.T) {
	pbStream := pb.Stream{Track: "something"}
	entityStream := entity.Stream{Track: pbStream.GetTrack()}

	storageMock := storageMocks.Interface{}
	storageMock.On("AddStream", &entityStream).Return(int64(0), errors.New("Integrity violation"))
	twitterclientMock := twitterclientMocks.Interface{}
	s := NewTweetwatchServiceServer(&storageMock, &twitterclientMock)

	req := pb.CreateStreamRequest{Stream: &pbStream}
	resp, err := s.CreateStream(context.Background(), &req)

	assert.Nil(t, resp, "Response must be nil where storage returns error")
	assert.NotNil(t, err, "Service must returns error")
}

func TestCreateStream_WrongApiVersion(t *testing.T) {
	stream := pb.Stream{Track: "something"}
	storageMock := storageMocks.Interface{}
	storageMock.On("AddStream", &stream).Return(int64(1), nil)
	twitterclientMock := twitterclientMocks.Interface{}
	s := NewTweetwatchServiceServer(&storageMock, &twitterclientMock)

	req := pb.CreateStreamRequest{Stream: &stream, Api: "v0"}
	resp, err := s.CreateStream(context.Background(), &req)

	assert.Nil(t, resp, "Response must be nil where storage returns error")
	assert.NotNil(t, err, "Service must returns error")
}

func TestGetStreams_Successful(t *testing.T) {
	storageMock := storageMocks.Interface{}
	storageMock.On("GetStreams").Return([]entity.StreamInterface{}, nil)
	twitterclientMock := twitterclientMocks.Interface{}
	s := NewTweetwatchServiceServer(&storageMock, &twitterclientMock)

	req := pb.GetStreamsRequest{}
	resp, err := s.GetStreams(context.Background(), &req)

	assert.Nil(t, err, "Error must be equal nil")
	assert.NotNil(t, resp, "Response must be not null")
	assert.IsType(t, []*pb.Stream{}, resp.GetStreams())
}
