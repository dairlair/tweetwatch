package storage

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/dchest/uniuri"
	"golang.org/x/crypto/bcrypt"
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

	if err != nil {
		return "", pgError(err)
	}

	hash := hashAndSalt([]byte(password))
	token = uniuri.New()

	if err := tx.QueryRow(signUpSQL,
		email,
		hash,
		token,
	).Scan(&token); err != nil {
		tx.Rollback()
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

	log.Infof("Retrieved hash: %v, token: %v", storedHash, storedToken)

	if !comparePasswords(storedHash, []byte(password)) {
		log.Infof("Invalid credentials")
		return "", errors.New("Invalid credentials")
	}

	return storedToken, nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Warn(err)
		return false
	}

	return true
}

/** @TODO Move this methods to security / auth package */
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Warn(err)
	}
	return string(hash)
}
