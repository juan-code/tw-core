package db

import (
	"github.com/tw-core/models"
	"golang.org/x/crypto/bcrypt"
)

/*TryLogin is a function to try login user and compare password with hashpassword */
func TryLogin(email, password string) (models.User, bool) {
	user, finded, _ := CheckUserExistByEmail(email)
	if !finded {
		return user, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
