package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/lmSeryi/golang-twitter/models"
	"time"
)

/* GenerateJWT genereate JWT encrypt */
func GenerateJWT(t models.User) (string, error) {
	code := []byte("lmSeryiJuasJuas")
	payload := jwt.MapClaims{
		"_id":       t.ID.Hex(),
		"email":     t.Email,
		"name":      t.Name,
		"lastName":  t.LastName,
		"birthday":  t.Birthday,
		"biography": t.Biography,
		"location":  t.Location,
		"webSite":   t.WebSite,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(code)

	if err != nil {
		return token, err
	}
	return tokenStr, nil
}
