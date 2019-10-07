package entity

// StreamInterface describes entity Stream
type StreamInterface interface {
	GetID() int64
	GetTopicID() int64
	GetTrack() string
}

// Stream contains info required by twitter client to retrieve data from Twitter Streaming API and to store stream into the database
type Stream struct {
	ID    int64
	TopicID int64
	Track string
}

func NewStreams(tracks []string) []StreamInterface {
	var streams []StreamInterface
	for _, track := range tracks {
		stream := Stream{
			Track:   track,
		}
		streams = append(streams, &stream)
	}
	return streams
}

// GetID returns the Stream ID
func (s *Stream) GetID() int64 {
	return s.ID
}

// GetID returns the Stream ID
func (s *Stream) GetTopicID() int64 {
	return s.TopicID
}

// GetTrack returns the stream's track
func (s *Stream) GetTrack() string {
	return s.Track
}

func GetStreamIDs(streams []StreamInterface) (IDs []int64) {
	for _, stream := range streams {
		IDs = append(IDs, stream.GetID())
	}
	return IDs
}
