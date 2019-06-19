package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/dairlair/twitwatch/pkg/api/v1"
	storageMocks "github.com/dairlair/twitwatch/pkg/storage/mocks"
)

func TestCreateStream_RequestCreation(t *testing.T) {
	track := "something"
	req := pb.CreateStreamRequest{Stream: &pb.Stream{Track: track}}
	assert.Equal(t, track, req.GetStream().GetTrack(), "they should be equal")

}

func TestCreateStream_Successful(t *testing.T) {
	stream := pb.Stream{Track: "something"}

	var id int64 = 1
	storageMock := storageMocks.Interface{}
	storageMock.On("AddStream", &stream).Return(id, nil)
	s := NewTwitwatchServiceServer(&storageMock)

	req := pb.CreateStreamRequest{Stream: &stream}
	resp, err := s.CreateStream(context.Background(), &req)
	assert.Nil(t, err, "Error must be equal nil")
	assert.Equal(t, id, resp.Id, "Returned id mus be equal ID returned from storage")
}

func TestCreateStream_FailedOnStorage(t *testing.T) {
	stream := pb.Stream{Track: "something"}

	storageMock := storageMocks.Interface{}
	storageMock.On("AddStream", &stream).Return(int64(0), errors.New("Integrity violation"))
	s := NewTwitwatchServiceServer(&storageMock)

	req := pb.CreateStreamRequest{Stream: &stream}
	resp, err := s.CreateStream(context.Background(), &req)

	assert.Nil(t, resp, "Response must be nil where storage returns error")
	assert.NotNil(t, err, "Service must returns error")
}

func TestCreateStream_WrongApiVersion(t *testing.T) {
	stream := pb.Stream{Track: "something"}
	storageMock := storageMocks.Interface{}
	storageMock.On("AddStream", &stream).Return(int64(1), nil)
	s := NewTwitwatchServiceServer(&storageMock)

	req := pb.CreateStreamRequest{Stream: &stream, Api: "v0"}
	resp, err := s.CreateStream(context.Background(), &req)

	assert.Nil(t, resp, "Response must be nil where storage returns error")
	assert.NotNil(t, err, "Service must returns error")
}

func TestGetStreams_Successfull(t *testing.T) {
	storageMock := storageMocks.Interface{}
	s := NewTwitwatchServiceServer(&storageMock)

	req := pb.GetStreamsRequest{}
	resp, err := s.GetStreams(context.Background(), &req)
	assert.Nil(t, err, "Error must be equal nil")
	assert.NotNil(t, resp, "Response must be not null")
}
