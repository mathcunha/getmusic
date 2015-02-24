package main

import (
	"github.com/mathcunha/getmusic/gmusic"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/view/", staticHandler)
	http.HandleFunc("/api/", gmusic.RestHandlerV1)
	http.ListenAndServe(getPort(), nil)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
