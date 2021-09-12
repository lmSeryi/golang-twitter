package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lmSeryi/golang-twitter/db"
	"github.com/lmSeryi/golang-twitter/models"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Invalid credentials "+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.UpdateProfile(t, IdUser)
	if err != nil {
		http.Error(w, "Error while updating profile "+err.Error(), 500)
		return
	}

	if !status {
		http.Error(w, "The profile has not been updated "+err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
