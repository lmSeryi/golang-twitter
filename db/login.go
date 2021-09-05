package db

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/lmSeryi/golang-twitter/models"
)

func Login(email string, password string) (models.User, bool) {
	usu, found, _ := UserExists(email)

	if !found {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
