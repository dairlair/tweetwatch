package storage

import (
	"errors"

	"github.com/dairlair/tweetwatch/pkg/security"
	"github.com/dchest/uniuri"
)

// SignUp registers new user in system
func (storage *Storage) SignUp(email string, password string) (token string, err error) {
	const signUpSQL = `
		INSERT INTO "user" (
			email
			, hash
			, token
		) VALUES (
			$1, $2, $3
		) RETURNING token
	`

	tx, err := storage.connPool.Begin()
	if err != nil {
		return "", pgError(err)
	}
	defer func() {
		if err != nil {
			pgRollback(tx)
		}
	}()

	hash := security.HashAndSalt([]byte(password))
	token = uniuri.New()

	if err := tx.QueryRow(signUpSQL,
		email,
		hash,
		token,
	).Scan(&token); err != nil {
		return "", pgError(err)
	}

	if err := tx.Commit(); err != nil {
		return "", pgError(err)
	}

	return token, nil
}

// SignIn authenticate user, set temp token for him and returns token
func (storage *Storage) SignIn(email string, password string) (token string, err error) {
	const signInSQL = `
		SELECT 
			hash
			, token
		FROM
			"user"
		WHERE LOWER(email) = $1
	`

	var storedHash string
	var storedToken string

	if err := storage.connPool.QueryRow(signInSQL, email).Scan(&storedHash, &storedToken); err != nil {
		return "", pgError(err)
	}

	if !security.ComparePasswords(storedHash, []byte(password)) {
		return "", errors.New("invalid credentials")
	}

	return storedToken, nil
}
