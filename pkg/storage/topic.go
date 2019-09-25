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
		) RETURNING topic_id, user_id, name, created_at, is_active
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
			&createdTopic.CreatedAt,
			&createdTopic.IsActive,
		); err != nil {
		return nil, pgError(err)
	}

	for _, stream := range topic.GetStreams() {
		st := entity.Stream{
			TopicID: createdTopic.ID,
			Track:   stream.GetTrack(),
		}
		createdStream, err := storage.addStream(tx, &st)
		if err != nil {
			log.Errorf("error: %s", err)
			return nil, err
		}
		log.Infof("stream created: %v", createdStream)
		createdTopic.Streams = append(createdTopic.Streams, createdStream)
	}

	if err := tx.Commit(); err != nil {
		return nil, pgError(err)
	}

	result = &createdTopic

	return result, nil
}

func (storage *Storage) GetUserTopics(userId int64) (result []entity.TopicInterface, err error) {
	const sql = `
	`
	var topics []entity.TopicInterface

	return topics, nil
}