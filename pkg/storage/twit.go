package storage

import (
	"github.com/dairlair/twitwatch/pkg/entity"
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
func (storage *Storage) AddTwit(twit entity.TwitInterface) (id int64, err error) {
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

	if err := tx.Commit(); err != nil {
		return 0, pgError(err)
	}

	return id, nil
}
