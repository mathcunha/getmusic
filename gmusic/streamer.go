package gmusic

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Streamer interface {
	GetPlaylistService(body []byte) ([]Music, error)
}

func GetPlaylist(url string, s Streamer) ([]Music, error) {
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("error calling service at %v  [%v]", url, err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("error reading body at %v  [%v]", url, err)
		return nil, err
	}

	return s.GetPlaylistService(body)

}
