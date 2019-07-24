package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/dairlair/tweetwatch/pkg/api/v1"
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/dairlair/tweetwatch/pkg/storage"
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
func NewTwitwatchServiceServer(s storage.Interface) pb.TwitwatchServiceServer {
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
func (s *twitwatchServiceServer) CreateStream(ctx context.Context, req *pb.CreateStreamRequest) (*pb.CreateStreamResponse, error) {
	// Check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	// Insert stream entity data
	id, err := s.storage.AddStream(entity.NewStream(req.GetStream().GetId(), req.GetStream().GetTrack()))
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to save stream-> "+err.Error())
	}

	return &pb.CreateStreamResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// GetStreams Returns list of streams
func (s *twitwatchServiceServer) GetStreams(ctx context.Context, req *pb.GetStreamsRequest) (*pb.GetStreamsResponse, error) {
	// Check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	streams, err := s.storage.GetStreams()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve streams-> "+err.Error())
	}

	pbStreams := []*pb.Stream{}
	for _, stream := range streams {
		pbStream := pb.Stream{Id: stream.GetID(), Track: stream.GetTrack()}
		pbStreams = append(pbStreams, &pbStream)
	}

	return &pb.GetStreamsResponse{
		Api:     apiVersion,
		Streams: pbStreams,
	}, nil
}

func (s *twitwatchServiceServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	// Check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	token, err := s.storage.SignUp(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to sign up-> "+err.Error())
	}

	return &pb.SignUpResponse{
		Api:   apiVersion,
		Token: token,
	}, nil
}

func (s *twitwatchServiceServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	// Check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	token, err := s.storage.SignIn(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to sign in-> "+err.Error())
	}

	return &pb.SignInResponse{
		Api:   apiVersion,
		Token: token,
	}, nil
}
