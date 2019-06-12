package storage

import (
	apiV1 "github.com/dairlair/twitwatch/pkg/api/v1"
)

const (
	addStreamSQL = `
		INSERT INTO stream (
			track
		) VALUES (
			$1
		) RETURNING stream_id
	`
)

// AddStream inserts stream into database
func (storage *Storage) AddStream(stream *apiV1.Stream) (id int64, err error) {
	tx, err := storage.connPool.Begin()
	if err != nil {
		return 0, pgError(err)
	}

	if err := tx.QueryRow(addStreamSQL, stream.Track).Scan(&id); err != nil {
		tx.Rollback()
		return 0, pgError(err)
	}

	if err := tx.Commit(); err != nil {
		return 0, pgError(err)
	}

	return id, nil
}
