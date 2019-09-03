// Package twitterclient provides wrapper around Twitter Streaming API
// The package accepts in config the Storage Interface which provides methods for retrieve active streams and store twits with treir steam
package twitterclient

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
)

// StorageInterface declares dependency for twitterclient.
type StorageInterface interface {
	AddTwit(entity.TwitInterface) (id int64, err error)
	// Twitterclient need to retrieve from
	GetActiveStreams() (streams []entity.StreamInterface, err error)
}

// InstanceInterface defines the main object interface which is created by this package.
type InstanceInterface interface {
	// Creates Twitter Streaming API client and validates credentials.
	Start() error
	// Add stream to watch it.
	AddStream(entity.StreamInterface)
	// Returns currently loaded and watched streams.
	GetStreams() map[int64]entity.StreamInterface
	// Runs watching for twits according to specified streams.
	Watch() error
	// Stops watching for all specified streams.
	Unwatch()
}
