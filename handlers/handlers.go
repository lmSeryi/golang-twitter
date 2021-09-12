package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/lmSeryi/golang-twitter/middlew"
	"github.com/lmSeryi/golang-twitter/routers"
)

/* Handlers setting port and raise server */
func Handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/sign-up", middlew.CheckDb(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDb(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.CheckDb(middlew.JwtValidation(routers.Profile))).Methods("GET")
	router.HandleFunc("/update-profile", middlew.CheckDb(middlew.JwtValidation(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/create-tweet", middlew.CheckDb(middlew.JwtValidation(routers.CreateTweet))).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
