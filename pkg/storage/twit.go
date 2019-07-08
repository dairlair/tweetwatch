package storage

import (
	"github.com/dairlair/tweetwatch/pkg/entity"
)

const (
	addTwitSQL = `
		INSERT INTO twit (
			id
			, user_id
			, full_text
			, created_at
		) VALUES (
			$1, $2, $3, $4
		) RETURNING twit_id
	`
)

// AddTwit just insert tweet into database
func (storage *Storage) AddTwit(twit entity.TwitInterface, streamIds []int64) (id int64, err error) {
	tx, err := storage.connPool.Begin()

	if err != nil {
		return 0, pgError(err)
	}

	if err := tx.QueryRow(addTwitSQL,
		twit.GetID(),
		twit.GetTwitterUserID(),
		twit.GetFullText(),
		twit.GetCreatedAt(),
	).Scan(&id); err != nil {
		return 0, pgError(err)
	}

	err = addTwitStreams(storage.connPool, id, streamIds)
	if err != nil {
		return 0, pgError(err)
	}

	if err := tx.Commit(); err != nil {
		return 0, pgError(err)
	}

	return id, nil
}

func addTwitStreams(conn *pgx.ConnPool, twitId int64, streamIds []int64) error {
	for _, streamId := range streamIds {
		err := addTwitStream(conn, twitId, streamId)
		if err != nil {
			return err
		}
	}
	return nil
}

func addTwitStream(conn *pgx.ConnPool, twitId int64, streamId int64) error {
	const sql = `INSERT INTO twit_stream (twit_id, stream_id) VALUES ($1, $2)`
	_, err := conn.Exec(sql, twitId, streamId)
	return err
}