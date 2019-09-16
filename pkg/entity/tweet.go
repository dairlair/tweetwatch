package entity

// TweetInterface defines methods implemented by the platform's entity twit.
type TweetInterface interface {
	GetID() int64
	GetTwitterID() int64
	GetTwitterUserID() int64
	GetFullText() string
	GetCreatedAt() string
}

// Tweet is a basic structure implementing TweetInterface
type Tweet struct {
	ID            int64
	TwitterID     int64
	TwitterUserID int64
	FullText      string
	CreatedAt     string
}

// GetID returns the Tweet ID from the twitwatch platform
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

// GetFullText returns twit's full text
func (t *Tweet) GetFullText() string {
	return t.FullText
}

// GetCreatedAt returns the user id from twitter
func (t *Tweet) GetCreatedAt() string {
	return t.CreatedAt
}
