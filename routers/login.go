package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lmSeryi/golang-twitter/jwt"
	"github.com/lmSeryi/golang-twitter/models"
)

/* Login do login */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().add("Content-Type", "application/json")

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Credentials not valid", 400)
		return
	}
	if len(t.Email) == 0{
		http.Error(w, "Email is required", 400)
		return
	}
	doc, exist := db.Login(t.Email, t.Password)
	if !exist{
		http.Error(w, "Credentials not valid", 400)
		return
	}
	
	jwtKey, err := jwt.GenerateJWT(doc)
	if err != nil {
		http.Error(w, "Error while getting token", 500)
		return
	}

	resp := models.AnsLogin{
		Token : jwtKey
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncode(w).Encode(resp)

	expirationTime := time.Now().add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "Token",
		Value: jwtKey,
		Expires: expirationTime
	})
}
