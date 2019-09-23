package entity

// TopicInterface defines methods implemented by the platform's entity topic.
type TopicInterface interface {
	GetID() int64
	GetUserID() int64
	GetName() string
	GetTracks() []string
	GetCreatedAt() string
}

// Topic is a basic structure implementing TopicInterface
type Topic struct {
	ID        int64
	UserID    int64
	Name      string
	Tracks    []string
	CreatedAt string
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
	return t.Tracks
}

// GetCreatedAt returns twit's full text
func (t *Topic) GetCreatedAt() string {
	return t.CreatedAt
}
