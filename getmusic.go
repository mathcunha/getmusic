package main

import (
	"github.com/mathcunha/getmusic/gmusic"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", staticHandler)
	http.HandleFunc("/api/", gmusic.RestHandlerV1)
	http.ListenAndServe(":8080", nil)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
