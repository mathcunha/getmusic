package gmusic

import (
	"testing"
)

func TestMagnet(t *testing.T) {
	m := []Music{Music{"Maroon 5", "V", "Sugar", ""}}
	t.Logf("%v", GetMagnetLinks(m))
}
