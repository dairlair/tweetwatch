package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
	log "github.com/sirupsen/logrus"
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
		) RETURNING topic_id, user_id, name, tracks, created_at
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

	createdTopic := entity.Topic{}
	if err := tx.QueryRow(
			addTopicSQL,
			topic.GetUserID(),
			topic.GetName(),
			topic.GetTracks(),
		).Scan(
			&createdTopic.ID,
			&createdTopic.UserID,
			&createdTopic.Name,
			&createdTopic.Tracks,
			&createdTopic.CreatedAt,
		); err != nil {
		return nil, pgError(err)
	}

	for _, track := range topic.GetTracks() {
		stream, err := storage.addStream(tx, &entity.Stream{Track:track, TopicID: createdTopic.ID})
		if err != nil {
			log.Errorf("error: %s", err)
			return nil, err
		}
		log.Debugf("stream created: %v", stream)
	}

	if err := tx.Commit(); err != nil {
		return nil, pgError(err)
	}

	result = &createdTopic

	return result, nil
}