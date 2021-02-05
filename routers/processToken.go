package routers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tw-core/db"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tw-core/models"
)

//Email of user
var Email string

//IDUser id of user
var IDUser string

/*ProcessToken func process token  */
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	key := []byte("supermaster")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	fmt.Println(splitToken, token)
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("Invalid format token")
	}

	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {
		_, finded, _ := db.CheckUserExistByEmail(claims.Email)
		if finded {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, finded, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("Invalid token")
	}
	return claims, false, "", err
}
