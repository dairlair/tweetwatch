package entity

import "time"

// TopicInterface defines methods implemented by the platform's entity topic.
type TopicInterface interface {
	GetID() int64
	GetUserID() int64
	GetName() string
	GetCreatedAt() time.Time
	GetIsActive() bool
}

// Topic is a basic structure implementing TopicInterface
type Topic struct {
	ID        int64
	UserID    int64
	Name      string
	CreatedAt time.Time
	IsActive bool
}

// GetID returns the Topic ID from the tweetwatch platform
func (t *Topic) GetID() int64 {
	return t.ID
}

// GetID returns the topic's owner User ID from the tweetwatch platform
func (t *Topic) GetUserID() int64 {
	return t.UserID
}

// GetName returns the topic's name
func (t *Topic) GetName() string {
	return t.Name
}

// GetCreatedAt returns twit's full text
func (t *Topic) GetCreatedAt() time.Time {
	return t.CreatedAt
}

// GetIsActive returns twit's full text
func (t *Topic) GetIsActive() bool {
	return t.IsActive
}