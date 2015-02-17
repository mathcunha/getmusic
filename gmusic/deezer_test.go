package gmusic

import (
	"testing"
)

func TestDeezerPlaylist(t *testing.T) {
	s := Deezer{}

	music, err := GetPlaylist("http://www.deezer.com/playlist/1127687863", &s)

	if err != nil {
		t.Errorf("config error %v", err)
	}

	t.Logf("returned musics %v", music)
}
