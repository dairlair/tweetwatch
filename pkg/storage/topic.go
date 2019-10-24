package storage

import (
	"errors"
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/jackc/pgx"
)

// AddTopic inserts topic into database
func (storage *Storage) AddTopic(topic entity.TopicInterface) (result entity.TopicInterface, err error) {
	const addTopicSQL = `
		INSERT INTO topic (
			user_id
			, name
			, is_active
		) VALUES (
			$1, $2, $3
		) RETURNING 
			topic_id
			, user_id
			, name
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
			topic.GetIsActive(),
		).Scan(
			&createdTopic.ID,
			&createdTopic.UserID,
			&createdTopic.Name,
			&createdTopic.CreatedAt,
			&createdTopic.IsActive,
		); err != nil {
		return nil, pgError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, pgError(err)
	}

	result = &createdTopic

	return result, nil
}

func getTopicByID(tx *pgx.Tx, topicID int64) (result entity.TopicInterface, err error) {
	if topicID < 1 {
		return nil, errors.New("the topic ID is required")
	}
	const sql = `
		SELECT
			topic_id
			, user_id
			, name
			, created_at
			, is_active
		FROM topic 
		WHERE 
			topic_id = $1 
			AND is_deleted = FALSE
	`
	row := tx.QueryRow(sql, topicID)
	topic := entity.Topic{}
	err = row.Scan(
		&topic.ID,
		&topic.UserID,
		&topic.Name,
		&topic.CreatedAt,
		&topic.IsActive,
	)
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

func (storage *Storage) GetTopic(topicID int64) (topic entity.TopicInterface, err error) {
	if topicID < 1 {
		return nil, errors.New("the topic ID is required")
	}
	const sql = `
		SELECT
			topic_id
			, user_id
			, name
			, created_at
			, is_active
		FROM topic 
		WHERE 
			topic_id = $1 
			AND is_deleted = FALSE
	`
	row := storage.connPool.QueryRow(sql, topicID)
	topicEntity := entity.Topic{}
	err = row.Scan(
		&topicEntity.ID,
		&topicEntity.UserID,
		&topicEntity.Name,
		&topicEntity.CreatedAt,
		&topicEntity.IsActive,
	)
	if err != nil {
		return nil, err
	}
	topic = &topicEntity
	return topic, nil
}

// AddTopic inserts topic into database
func (storage *Storage) UpdateTopic(topic entity.TopicInterface) (result entity.TopicInterface, err error) {
	tx, err := storage.connPool.Begin()
	if err != nil {
		return nil, pgError(err)
	}
	defer func() {
		if err != nil {
			pgRollback(tx)
		}
	}()

	_, err = getTopicByID(tx, topic.GetID())
	if err != nil {
		return nil, err
	}

	// Update the main topic record
	const sql = `
		UPDATE topic SET
			name = $2
			, is_active = $3
		WHERE topic_id = $1
	`
	_, err = tx.Exec(
		sql,
		topic.GetID(),
		topic.GetName(),
		topic.GetIsActive(),
	)
	if err != nil {
		return nil,  err
	}


	// Read saved topic
	savedTopic, err := getTopicByID(tx, topic.GetID())
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, pgError(err)
	}

	return savedTopic, err
}

func (storage *Storage) GetUserTopics(userId int64) (result []entity.TopicInterface, err error) {
	const sql = `
		SELECT
			topic_id
			, user_id
			, name
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