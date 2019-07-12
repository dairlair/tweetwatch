package entity

// TwitInterface defines methods implemented by the platform's entity twit.
type TwitInterface interface {
	GetID() int64
	GetTwitterID() int64
	GetTwitterUserID() int64
	GetFullText() string
	GetCreatedAt() string
}

// Twit is a basic structure implementing TwitInterface
type Twit struct {
	ID            int64
	TwitterID     int64
	TwitterUserID int64
	FullText      string
	CreatedAt     string
}

// GetID returns the Twit ID from the twitwatch platform
func (t *Twit) GetID() int64 {
	return t.ID
}

// GetTwitterID returns the tweet id from twitter
func (t *Twit) GetTwitterID() int64 {
	return t.TwitterID
}

// GetTwitterUserID returns the user id from twitter
func (t *Twit) GetTwitterUserID() int64 {
	return t.TwitterUserID
}

// GetFullText returns twit's full text
func (t *Twit) GetFullText() string {
	return t.FullText
}

// GetCreatedAt returns the user id from twitter
func (t *Twit) GetCreatedAt() string {
	return t.CreatedAt
}
