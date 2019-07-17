package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
	log "github.com/sirupsen/logrus"
)

// AddStream inserts stream into database
func (storage *Storage) AddStream(stream entity.StreamInterface) (id int64, err error) {

	const addStreamSQL = `
		INSERT INTO stream (
			track
		) VALUES (
			$1
		) RETURNING stream_id
	`

	tx, err := storage.connPool.Begin()
	if err != nil {
		return 0, pgError(err)
	}

	if err := tx.QueryRow(addStreamSQL, stream.GetTrack()).Scan(&id); err != nil {
		txError := tx.Rollback()
		if txError != nil {
			log.Errorf("transaction closing failed. %s", txError)
		}
		return 0, pgError(err)
	}

	if err := tx.Commit(); err != nil {
		return 0, pgError(err)
	}

	return id, nil
}

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