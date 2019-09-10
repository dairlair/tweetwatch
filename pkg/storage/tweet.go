package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/jackc/pgx"
)

// AddTwit just insert tweet into database
func addTweet(tx *pgx.Tx, tweet entity.TweetInterface) (id int64, err error) {
	const addTweetSQL = `
		INSERT INTO tweet (
			id
			, user_id
			, full_text
			, created_at
		) VALUES (
			$1, $2, $3, $4
		) RETURNING tweet_id
	`

	if err := tx.QueryRow(addTweetSQL,
		tweet.GetID(),
		tweet.GetTwitterUserID(),
		tweet.GetFullText(),
		tweet.GetCreatedAt(),
	).Scan(&id); err != nil {
		return 0, pgError(err)
	}

	return id, nil
}

func addTweetStream(tx *pgx.Tx, tweetId int64, streamId int64) (id int64, err error) {
	const addTweetStreamSQL = `
		INSERT INTO tweet_stream (
			tweet_id
			, stream_id
		) VALUES (
			$1, $2
		) RETURNING tweet_id
	`

	if err := tx.QueryRow(addTweetStreamSQL,
		tweetId,
		streamId,
	).Scan(&id); err != nil {
		return 0, pgError(err)
	}

	return id, nil
}

func (storage *Storage) AddTweetStreams(tweetStreams entity.TweetStreamsInterface) (id int64, err error) {
	tx, err := storage.connPool.Begin()
	if err != nil {
		return 0, pgError(err)
	}
	defer func() {
		if err != nil {
			pgRollback(tx)
		}
	}()

	// Add tweet and get his ID.
	tweetId, err := addTweet(tx, tweetStreams.GetTweet())

	for _, stream := range tweetStreams.GetStreams() {
		_, err = addTweetStream(tx, tweetId, stream.GetID())
		if err != nil {
			return 0, pgError(err)
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, pgError(err)
	}

	return id, nil
}