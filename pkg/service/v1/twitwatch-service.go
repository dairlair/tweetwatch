package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apiV1 "github.com/dairlair/twitwatch/pkg/api/v1"
	"github.com/dairlair/twitwatch/pkg/storage"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// twitwatchServiceServer is implementation of v1.ToDoServiceServer proto interface
type twitwatchServiceServer struct {
	storage storage.Interface
}

// NewTwitwatchServiceServer creates TwitWatch service
func NewTwitwatchServiceServer(s storage.Interface) apiV1.TwitwatchServiceServer {
	return &twitwatchServiceServer{storage: s}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *twitwatchServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// Create new stream
func (s *twitwatchServiceServer) CreateStream(ctx context.Context, req *apiV1.CreateStreamRequest) (*apiV1.CreateStreamResponse, error) {
	// Check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	// Insert stream entity data
	id, err := s.storage.AddStream(req.GetStream())
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to save stream-> "+err.Error())
	}

	return &apiV1.CreateStreamResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}
