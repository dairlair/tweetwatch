//
// Package twitterclient provides wrapper around Twitter Streaming API
// The package accepts in config the Storage Interface which provides methods for retrieve active streams and store twits with their steam
//
package twitterclient

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
)

// Interface defines the main object interface which is created by this package.
type Interface interface {
	// Creates Twitter Streaming API client and validates credentials.
	Start() error
	// Add stream to watch it.
	AddStream(entity.StreamInterface)
	// Returns currently loaded and watched streams.
	GetStreams() map[int64]entity.StreamInterface
	// Runs watching for twits according to specified streams.
	Watch(chan entity.TweetStreamsInterface) error
	// Stops watching for all specified streams.
	Unwatch()
}
