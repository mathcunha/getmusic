package gmusic

import (
	"fmt"
	"regexp"
)

func FindPattern(re *regexp.Regexp, body []byte, mp *map[string]string) {
	n1 := re.SubexpNames()
	r3 := re.FindAllSubmatch(body, -1)
	mapa := *mp
	if len(r3) > 0 {
		for j, r2 := range r3 {
			if len(r2) > 0 {
				for i, n := range r2 {
					mapa[KeyName(n1[i], j)] = string(n)
				}
			}
		}
	}
}

func KeyName(key string, i int) string {
	if len(key) != 0 && i != 0 {
		return fmt.Sprintf("%v_%v", key, i)
	}
	return key
}
