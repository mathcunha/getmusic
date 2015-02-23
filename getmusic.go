package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/view/", staticHandler)
	http.ListenAndServe(":8080", nil)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
