package twitterstream

// Interface defines all methods that are required by twitterclient
// to retrieve twits from Twitter Streaming API.
type Interface interface {
	GetID() int64
	GetTrack() string
}

// Stream contains info required by twitter client to retrieve data from Twitter Streaming API
type Stream struct {
	ID    int64
	Track string
}

// GetID returns the Stream ID
func (s *Stream) GetID() int64 {
	return s.ID
}

// GetTrack returns the stream's track
func (s *Stream) GetTrack() string {
	return s.Track
}
