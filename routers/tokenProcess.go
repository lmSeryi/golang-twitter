package routers

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/lmSeryi/golang-twitter/db"
	"github.com/lmSeryi/golang-twitter/models"
)

var Email string

var IdUser string

/* TokenProcess extact token values */
func TokenProcess(tk string) (*models.Claim, bool, string, error) {
	code := []byte("lmSeryiJuasJuas")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(toke *jwt.Token) (interface{}, error) {
		return code, nil
	})

	if err == nil {
		_, found, _ := db.UserExists(claims.Email)
		if found {
			Email = claims.Email
			IdUser = claims.ID.Hex()
		}
		return claims, found, IdUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}
