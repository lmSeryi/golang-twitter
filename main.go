package main

import (
	"log"

	"github.com/lmSeryi/golang-twitter/db"
	"github.com/lmSeryi/golang-twitter/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Without connection")
		return
	}
	handlers.Handlers()
}
