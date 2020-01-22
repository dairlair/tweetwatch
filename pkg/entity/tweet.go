package entity

import "time"

// TweetInterface defines methods implemented by the platform's entity twit.
type TweetInterface interface {
	GetID() int64
	GetTwitterID() int64
	GetTwitterUserID() int64
	GetTwitterUsername() string
	GetFullText() string
	GetCreatedAt() time.Time
}

// Tweet is a basic structure implementing TweetInterface
type Tweet struct {
	ID              int64
	TwitterID       int64
	TwitterUserID   int64
	TwitterUsername string
	FullText        string
	CreatedAt       time.Time
}

// GetID returns the Tweet ID from the tweetwatch platform
func (t *Tweet) GetID() int64 {
	return t.ID
}

// GetTwitterID returns the tweet id from twitter
func (t *Tweet) GetTwitterID() int64 {
	return t.TwitterID
}

// GetTwitterUserID returns the user id from twitter
func (t *Tweet) GetTwitterUserID() int64 {
	return t.TwitterUserID
}

// GetTwitterUserID returns the user id from twitter
func (t *Tweet) GetTwitterUsername() string {
	return t.TwitterUsername
}

// GetFullText returns twit's full text
func (t *Tweet) GetFullText() string {
	return t.FullText
}

// GetCreatedAt returns the user id from twitter
func (t *Tweet) GetCreatedAt() time.Time {
	return t.CreatedAt
}
