package security

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), plainPwd) == nil
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Warn(err)
	}
	return string(hash)
}
