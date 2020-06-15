package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"

	rankmanager "github.com/Ekram-B2/rankmanager/rank"
)

func main() {
	// 1. Set up router object to define paths which wrap execution logic
	r := chi.NewRouter()

	// 2. Define the endpoints required of for the task
	r.Get("/determineRank", rankmanager.HandleRequestToDetermineRank)

	// 3. Determine the binding port
	var bindingPort string
	if os.Getenv("DEPLOYMENT_TYPE") == "1" {
		// Hardcoded the port number in development mode
		bindingPort = ":8081"
	} else {
		bindingPort = ":" + os.Getenv("PORT")
	}
	// 4. start up an http server object at the port
	http.ListenAndServe(bindingPort, r)
}
