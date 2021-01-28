package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tw-core/models"
)

/*GenerateToken function that generate a token  */
func GenerateToken(T models.User) (string, error) {
	key := []byte("supermaster")
	payload := jwt.MapClaims{
		"email":         T.Email,
		"name":          T.Name,
		"lastName":      T.LastName,
		"biobliography": T.Bibliography,
		"location":      T.Location,
		"website":       T.WebSite,
		"_id":           T.ID.Hex(),
		"exp":           time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString(key)
}
