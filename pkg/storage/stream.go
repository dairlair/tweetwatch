package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
	"github.com/jackc/pgx"
)

func (storage *Storage) addStream(tx *pgx.Tx, stream entity.StreamInterface) (result entity.StreamInterface, err error) {
	const addStreamSQL = `
		INSERT INTO stream (
			topic_id
			, track
		) VALUES (
			$1, $2
		) RETURNING stream_id, topic_id, track
	`

	createdStream := entity.Stream{}
	if err := tx.QueryRow(
			addStreamSQL,
			stream.GetTopicID(),
			stream.GetTrack(),
		).Scan(
			&createdStream.ID,
			&createdStream.TopicID,
			&createdStream.Track,
		); err != nil {
		return nil, pgError(err)
	}

	result = &createdStream

	return result, nil
}

// AddStream inserts stream into database
// @DEPRECATED
//func (storage *Storage) AddStream(stream entity.StreamInterface) (result entity.StreamInterface, err error) {
//	const addStreamSQL = `
//		INSERT INTO stream (
//			track
//		) VALUES (
//			$1
//		) RETURNING stream_id
//	`
//
//	tx, err := storage.connPool.Begin()
//	if err != nil {
//		return nil, pgError(err)
//	}
//	defer func() {
//		if err != nil {
//			pgRollback(tx)
//		}
//	}()
//
//	var id int64
//	if err := tx.QueryRow(addStreamSQL, stream.GetTrack()).Scan(&id); err != nil {
//		return nil, pgError(err)
//	}
//
//	if err := tx.Commit(); err != nil {
//		return nil, pgError(err)
//	}
//
//	result = entity.NewStream(id, stream.GetTrack())
//
//	return result, nil
//}

// GetStreams returns all existing streams
func (storage *Storage) GetStreams() (streams []entity.StreamInterface, err error) {
	const getStreamsSQL = `
		SELECT 
			stream_id
			, track
		FROM
			stream
	`

	rows, err := storage.connPool.Query(getStreamsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var stream entity.Stream
		if err := rows.Scan(
			&stream.ID,
			&stream.Track,
		); err != nil {
			return nil, err
		}
		streams = append(streams, &stream)
	}

	return streams, nil
}

// GetStreams returns all active streams (streams with flag "is_active" = TRUE)
func (storage *Storage) GetActiveStreams() (streams []entity.StreamInterface, err error) {
	// @TODO Refactor when tweet table got is_active flag.
	return storage.GetStreams()
}