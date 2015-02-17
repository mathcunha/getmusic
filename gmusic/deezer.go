package gmusic

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Deezer struct {
	Data []DeezerSong
}

type DeezerSong struct {
	Title  string `json:"SNG_TITLE"`
	Artist string `json:"ART_NAME"`
	Album  string `json:"ALB_TITLE"`
}

func (s *Deezer) GetPlaylistService(body []byte) ([]Music, error) {

	pattern := regexp.MustCompile("data: {(?P<data>.+)")

	mp := map[string]string{}

	FindPattern(pattern, body, &mp)

	content := fmt.Sprint("{", mp["data"])

	reader := strings.NewReader(content)
	dec := json.NewDecoder(reader)

	if err := dec.Decode(s); err == nil {
		music := make([]Music, len(s.Data), len(s.Data))

		for i, v := range s.Data {
			music[i] = Music{v.Artist, v.Album, v.Title, ""}
		}
		return music, nil
	} else {
		log.Printf("error reading data - %v \n", err)
		return nil, err
	}
}
