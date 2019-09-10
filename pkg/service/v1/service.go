package v1

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/dairlair/tweetwatch/pkg/api/v1"
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/dairlair/tweetwatch/pkg/storage"
	"github.com/dairlair/tweetwatch/pkg/twitterclient"
)

// tweetwatchServiceServer is implementation of pb.TwitwatchServiceServer proto interface
// See github.com/dairlair/tweetwatch/pkg/api/v1/TwitwatchServiceServer
type tweetwatchServiceServer struct {
	storage storage.Interface
	tweetStreamsChannel chan entity.TweetStreamsInterface
	twitterClient twitterclient.Interface
}

// NewTweetwatchServiceServer creates TwitWatch service
func NewTweetwatchServiceServer(s storage.Interface, t twitterclient.Interface) pb.TwitwatchServiceServer {
	server := &tweetwatchServiceServer{
		storage: s,
		tweetStreamsChannel: make(chan entity.TweetStreamsInterface),
		twitterClient: t,
	}

	server.up()


	return server
}

// Create new stream
func (s *tweetwatchServiceServer) CreateStream(ctx context.Context, req *pb.CreateStreamRequest) (*pb.CreateStreamResponse, error) {
	// Check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	stream := entity.NewStream(req.GetStream().GetId(), req.GetStream().GetTrack())

	// Insert stream entity data
	id, err := s.storage.AddStream(stream)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to save stream-> "+err.Error())
	}

	// Ask twitter client reload streams. @TODO implement it and remove this test code.
	s.twitterClient.AddStream(stream)
	s.twitterClient.Unwatch()
	_ = s.twitterClient.Watch(s.tweetStreamsChannel)


	return &pb.CreateStreamResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// GetStreams Returns list of streams
func (s *tweetwatchServiceServer) GetStreams(ctx context.Context, req *pb.GetStreamsRequest) (*pb.GetStreamsResponse, error) {
	// Check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	streams, err := s.storage.GetStreams()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve streams-> "+err.Error())
	}

	var pbStreams []*pb.Stream
	for _, stream := range streams {
		pbStream := pb.Stream{Id: stream.GetID(), Track: stream.GetTrack()}
		pbStreams = append(pbStreams, &pbStream)
	}

	return &pb.GetStreamsResponse{
		Api:     apiVersion,
		Streams: pbStreams,
	}, nil
}

func (s *tweetwatchServiceServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
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

func (s *tweetwatchServiceServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
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

func (s *tweetwatchServiceServer) up() {
	log.Infof("Tweetwatch service up...")
	go func(input chan entity.TweetStreamsInterface) {
		for tweetStreams := range input {
			log.Infof("Store tweet to the database. %v", tweetStreams.GetTweet().GetID())
		}
	} (s.tweetStreamsChannel)
	log.Infof("Tweetwatch service is ready to accept tweets")
	err := s.twitterClient.Start()
	if err != nil {
		log.Fatalf("twitterclient error: %s\n", err)
	}
	_ = s.twitterClient.Watch(s.tweetStreamsChannel)
}