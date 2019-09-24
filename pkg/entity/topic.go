package entity

import "time"

// TopicInterface defines methods implemented by the platform's entity topic.
type TopicInterface interface {
	GetID() int64
	GetUserID() int64
	GetName() string
	GetStreams() []StreamInterface
	GetTracks() []string
	GetCreatedAt() time.Time
	GetIsActive() bool
}

// Topic is a basic structure implementing TopicInterface
type Topic struct {
	ID        int64
	UserID    int64
	Name      string
	Streams   []StreamInterface
	CreatedAt time.Time
	IsActive bool
}

// GetID returns the Topic ID from the twitwatch platform
func (t *Topic) GetID() int64 {
	return t.ID
}

// GetID returns the topic's owner User ID from the twitwatch platform
func (t *Topic) GetUserID() int64 {
	return t.UserID
}

// GetName returns the topic's name
func (t *Topic) GetName() string {
	return t.Name
}

// GetTrack returns the topic's track
func (t *Topic) GetTracks() []string {
	var tracks []string
	for _, stream := range t.GetStreams() {
		tracks = append(tracks, stream.GetTrack())
	}
	return tracks
}

// GetCreatedAt returns twit's full text
func (t *Topic) GetCreatedAt() time.Time {
	return t.CreatedAt
}

// GetIsActive returns twit's full text
func (t *Topic) GetIsActive() bool {
	return t.IsActive
}

// GetIsActive returns twit's full text
func (t *Topic) GetStreams() []StreamInterface {
	return t.Streams
}