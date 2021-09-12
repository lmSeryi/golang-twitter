package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lmSeryi/golang-twitter/db"
	"github.com/lmSeryi/golang-twitter/models"
)

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	payload := models.CreateTweet{
		UserId:  IdUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.CreateTweet(payload)
	if err != nil {
		http.Error(w, "An error has happened while creating tweet "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "The tweet has not been created "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
