package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/jackc/pgx"
)

func (storage *Storage)  GetTopicTweets(topicID int64) (tweets []entity.TweetInterface, err error) {
	const sql = `
		SELECT 
			tweet_id
			, twitter_id
			, twitter_user_id
			, twitter_username
			, full_text
			, created_at
		FROM
			tweet 
		WHERE tweet_id IN (
			SELECT tweet_id FROM tweet_stream WHERE topic_id = $1
		)
		ORDER BY created_at DESC
	`

	rows, err := storage.connPool.Query(sql, topicID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tweet entity.Tweet
		if err := rows.Scan(
			&tweet.ID,
			&tweet.TwitterID,
			&tweet.TwitterUserID,
			&tweet.TwitterUsername,
			&tweet.FullText,
			&tweet.CreatedAt,
		); err != nil {
			return nil, err
		}
		tweets = append(tweets, &tweet)
	}

	return tweets, nil
}

// AddTwit just insert tweet into database
func addTweet(tx *pgx.Tx, tweet entity.TweetInterface) (id int64, err error) {
	const addTweetSQL = `
		INSERT INTO tweet (
			twitter_id
			, twitter_user_id
			, twitter_username
			, full_text
			, created_at
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING tweet_id
	`

	if err := tx.QueryRow(addTweetSQL,
		tweet.GetTwitterID(),
		tweet.GetTwitterUserID(),
		tweet.GetTwitterUsername(),
		tweet.GetFullText(),
		tweet.GetCreatedAt(),
	).Scan(&id); err != nil {
		return 0, pgError(err)
	}

	return id, nil
}

func addTweetStream(tx *pgx.Tx, tweetId int64, streamId int64, topicId int64) (id int64, err error) {
	const addTweetStreamSQL = `
		INSERT INTO tweet_stream (
			tweet_id
			, stream_id
			, topic_id
		) VALUES (
			$1, $2, $3
		) RETURNING tweet_id
	`

	if err := tx.QueryRow(addTweetStreamSQL,
		tweetId,
		streamId,
		topicId,
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
	if err != nil {
		return id, err
	}

	for _, stream := range tweetStreams.GetStreams() {
		_, err = addTweetStream(tx, tweetId, stream.GetID(), stream.GetTopicID())
		if err != nil {
			return 0, pgError(err)
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, pgError(err)
	}

	return id, nil
}