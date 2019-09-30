package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
	log "github.com/sirupsen/logrus"
)

// AddTopic inserts topic into database
func (storage *Storage) AddTopic(topic entity.TopicInterface) (result entity.TopicInterface, err error) {
	const addTopicSQL = `
		INSERT INTO topic (
			user_id
			, name
			, tracks
		) VALUES (
			$1, $2, $3
		) RETURNING 
			topic_id
			, user_id
			, name
			, tracks
			, created_at
			, is_active
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

func (storage *Storage) GetTopicByID(topicID int64) (result entity.TopicInterface, err error) {
	const sql = `
		SELECT
			topic_id
			, user_id
			, name
			, tracks
			, created_at
			, is_active
		FROM topic 
		WHERE 
			topic_id = $1 
			AND is_deleted = FALSE
	`
	row := storage.connPool.QueryRow(sql, topicID)
	topic := entity.Topic{}
	err = row.Scan(
		&topic.ID,
		&topic.UserID,
		&topic.Name,
		&topic.Tracks,
		&topic.CreatedAt,
		&topic.IsActive,
	)
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

// AddTopic inserts topic into database
func (storage *Storage) UpdateTopic(topic entity.TopicInterface) (result entity.TopicInterface, err error) {
	return storage.GetTopicByID(topic.GetID())
}

func (storage *Storage) GetUserTopics(userId int64) (result []entity.TopicInterface, err error) {
	const sql = `
		SELECT
			topic_id
			, user_id
			, name
			, tracks
			, created_at
			, is_active
		FROM topic 
		WHERE 
			user_id = $1 
			AND is_deleted = FALSE
	`
	var topics []entity.TopicInterface

	rows, err := storage.connPool.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		topic := entity.Topic{}
		err := rows.Scan(
			&topic.ID,
			&topic.UserID,
			&topic.Name,
			&topic.Tracks,
			&topic.CreatedAt,
			&topic.IsActive,
		)
		if err != nil {
			return nil, err
		}
		topics = append(topics, &topic)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return topics, nil
}