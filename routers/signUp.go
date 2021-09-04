package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lmSeryi/golang-twitter/db"
	"github.com/lmSeryi/golang-twitter/models"
)

/* SignUp creates the regist in the DB */
func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error in data. "+err.Error(), 400)
		return
	}

	_, found, _ := db.UserExists(t.Email)
	if found {
		http.Error(w, "User already exists", 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "The email is required", 400)
		return
	}

	if len(t.Password) < 0 {
		http.Error(w, "Password should be have six characters", 400)
		return
	}

	if len(t.Password) < 0 {
		http.Error(w, "Password should be have six characters", 400)
		return
	}
	_, status, err := db.InsertRegist(t)
	if err != nil {
		http.Error(w, "Error has occured."+err.Error(), 500)
		return
	}

	if !status {
		http.Error(w, "Error has occured."+err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
