package gmusic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func RestHandlerV1(w http.ResponseWriter, r *http.Request) {
	a_path := strings.Split(r.URL.Path, "/")
	if "GET" != r.Method {
		http.Error(w, "method not allowed, try curl -X GET https://getmusic.herokuapp.com/streamer/playlist/id", http.StatusMethodNotAllowed)
		return
	}
	if "deezer" != a_path[2] {
		http.Error(w, "no handler to path "+r.URL.Path, http.StatusNotFound)
		return
	} else {
		callDeezerResource(w, r)
	}
}

func callDeezerResource(w http.ResponseWriter, r *http.Request) {
	a_path := strings.Split(r.URL.Path, "/")

	if len(a_path) < 5 {
		http.Error(w, "method not allowed, try curl -X GET https://getmusic.herokuapp.com/deezer/playlist/id", http.StatusNotFound)
		return
	}
	if "playlist" != a_path[3] {
		http.Error(w, "method not allowed, try curl -X GET https://getmusic.herokuapp.com/deezer/playlist/id", http.StatusNotFound)
	}
	if "dummy" == a_path[4] {
		music := getDummyPlaylist()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(music)
	} else if len(a_path[4]) > 0 {
		s := Deezer{}
		url := fmt.Sprint("http://www.deezer.com/playlist/", a_path[4])
		music, err := GetPlaylist(url, &s)
		if err != nil {
			http.Error(w, fmt.Sprint("error downloading ", url), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(music)
	} else {
		http.Error(w, "id is empty, try curl -X GET https://getmusic.herokuapp.com/deezer/playlist/id", http.StatusNotFound)
	}
}

func getDummyPlaylist() []Music {
	return []Music{Music{"Nome_1", "Artista", "Album_1", "2000"}, Music{"Nome_2", "Artistão", "Album_2", "2000"}}
}
