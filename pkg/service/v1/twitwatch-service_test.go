package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/dairlair/twitwatch/pkg/api/v1"
	storageMocks "github.com/dairlair/twitwatch/pkg/storage/mocks"
)

func TestCreateStreamRequestCreation(t *testing.T) {
	track := "something"
	req := pb.CreateStreamRequest{Stream: &pb.Stream{Track: track}}
	assert.Equal(t, track, req.GetStream().GetTrack(), "they should be equal")

}

func TestCreateStreamRequest(t *testing.T) {

	track := "something"
	stream := pb.Stream{Track: track}

	var id int64 = 1
	storageMock := storageMocks.Interface{}
	storageMock.On("AddStream", &stream).Return(id, nil)
	s := NewTwitwatchServiceServer(&storageMock)

	req := pb.CreateStreamRequest{Stream: &stream}
	resp, err := s.CreateStream(context.Background(), &req)
	assert.Nil(t, err, "Error must be equal nil")
	assert.Equal(t, id, resp.Id, "Returned id mus be equal ID returned from storage")
}
