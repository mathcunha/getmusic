package main

import (
	"flag"
	"fmt"
	"github.com/mathcunha/getmusic/gmusic"
)

var playlist string

func init() {
	flag.StringVar(&playlist, "playlist", "http://www.deezer.com/playlist/1127687863", "Deezer playlist url")
}

func main() {

	flag.Parse()

	s := gmusic.Deezer{}

	fmt.Println("Loading the playlist")
	m, _ := gmusic.GetPlaylist(playlist, &s)

	fmt.Println("Searching the magnet links")
	magnet := gmusic.GetMagnetLinks(m)

	for i, v := range m {
		fmt.Printf("%v - [%v] \n", v, magnet[i])
	}
}
