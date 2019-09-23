package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
	"time"
)

// AddStream inserts stream into database
func (storage *Storage) AddTopic(topic entity.TopicInterface) (result entity.TopicInterface, err error) {
	const addTopicSQL = `
		INSERT INTO topic (
			user_id
			, name
			, tracks
		) VALUES (
			$1, $2, $3
		) RETURNING topic_id, created_at
	`
	tx, err := storage.connPool.Begin()
	if err != nil {
		return nil, pgError(err)
	}
	defer func() {
		if err != nil {
			pgRollback(tx)
		}
	}()

	var id int64
	var createdAt time.Time
	if err := tx.QueryRow(
			addTopicSQL,
			topic.GetUserID(),
			topic.GetName(),
			topic.GetTracks(),
		).Scan(&id, &createdAt); err != nil {
		return nil, pgError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, pgError(err)
	}

	result = &entity.Topic{
		ID:        id,
		UserID:    topic.GetUserID(),
		Name:      topic.GetName(),
		Tracks:    topic.GetTracks(),
		CreatedAt: createdAt.Format("2006-01-02T15:04:05-0700"),
	}

	return result, nil
}