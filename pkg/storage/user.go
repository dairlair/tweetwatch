package storage

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/dairlair/tweetwatch/pkg/security"
	"github.com/dchest/uniuri"
)

const (
	signUpSQL = `
		INSERT INTO "user" (
			email
			, hash
			, token
		) VALUES (
			$1, $2, $3
		) RETURNING token
	`
	signInSQL = `
		SELECT 
			hash
			, token
		FROM
			"user"
		WHERE LOWER(email) = $1
	`
)

// SignUp registers new user in system
func (storage *Storage) SignUp(email string, password string) (token string, err error) {
	tx, err := storage.connPool.Begin()
	defer func() {
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				log.Fatalf("rollback failed-> %s", txErr)
			}
		}
	}()

	if err != nil {
		return "", pgError(err)
	}

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
