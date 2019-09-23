package storage

import (
	"errors"

	"github.com/dairlair/tweetwatch/pkg/security"
)

// SignUpHandler registers new user in system
func (storage *Storage) SignUp(email string, password string) (id *int64, err error) {
	const signUpSQL = `
		INSERT INTO "user" (
			email
			, hash
		) VALUES (
			$1, $2
		) RETURNING user_id
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

	hash := security.HashAndSalt([]byte(password))

	var userId int64
	if err := tx.QueryRow(signUpSQL,
		email,
		hash,
	).Scan(&userId); err != nil {
		return nil, pgError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, pgError(err)
	}

	return &userId, nil
}

// Login authenticate user, set temp token for him and returns token
func (storage *Storage) Login(email string, password string) (id *int64, err error) {
	const signInSQL = `
		SELECT 
			hash
			, user_id
		FROM
			"user"
		WHERE LOWER(email) = $1
	`

	var storedHash string

	if err := storage.connPool.QueryRow(signInSQL, email).Scan(&storedHash, id); err != nil {
		return nil, pgError(err)
	}

	if !security.ComparePasswords(storedHash, []byte(password)) {
		return nil, errors.New("invalid credentials")
	}

	return id, nil
}
