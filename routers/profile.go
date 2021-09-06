package routers

import (
	"encoding/json"
	"net/http"
	"github.com/lmSeryi/golang-twitter/db"
)

func Profile(w http.ReponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")

	if len(Id) < 1{
		http.Error(w, "Invalid Id", 400)
		return
	}

	profile, err := db.SearchProfile(Id)

	if err != nil{
		http.Error(w, "Error while searching profile " + err.Error(), 400)
		return
	}
	w.Header().Set("Context-Type", "application/json")
	w.Writeheader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
