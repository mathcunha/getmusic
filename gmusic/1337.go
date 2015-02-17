package gmusic

import (
	"fmt"
	"net/url"
	"regexp"
)

func GetMagnetLinks(m []Music) []string {
	retorno := make([]string, len(m), len(m))
	for i, v := range m {
		retorno[i] = GetMagnetLink(v)
	}

	return retorno
}

func GetMagnetLink(m Music) string {
	values := url.Values{}
	values.Set("", fmt.Sprint(m.Artist, " ", m.Name, " ", m.Year))
	query := values.Encode()
	query = query[1:len(query)]

	//Searching the music
	url := fmt.Sprint("http://1337x.to/search/", query, "/1/")
	body := GetUrl(url)
	href := LeetContent(body)

	//Searching the magnet link
	url = fmt.Sprint("http://1337x.to", href)
	body = GetUrl(url)
	href = LeetMagnet(body)

	return href
}

func LeetMagnet(b []byte) string {
	pattern := regexp.MustCompile("<ul class=\"download-links\">(?P<data>(.|\n)+)<ul class=\"category-name\">")
	mp := map[string]string{}
	FindPattern(pattern, b, &mp)

	//fmt.Printf("content [%v]", mp["data"])

	pattern = regexp.MustCompile("<a href=\"(?P<href>.*)\" oncli")
	FindPattern(pattern, []byte(mp["data"]), &mp)

	return mp["href"]
}

func LeetContent(b []byte) string {
	pattern := regexp.MustCompile("<span class=\"coll-2\">se</span>(?P<data>(.|\n)+)<div class=\"pagging-box\">")
	mp := map[string]string{}
	FindPattern(pattern, b, &mp)

	//fmt.Printf("content [%v]", mp["data"])

	pattern = regexp.MustCompile("<strong><a href=\"(?P<href>.*)\"")
	FindPattern(pattern, []byte(mp["data"]), &mp)

	return mp["href"]
}
