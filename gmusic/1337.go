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
	for i := 0; i < 2; i++ {

		values := url.Values{}
		if i == 0 {
			values.Set("", fmt.Sprint(m.Artist, " ", m.Name, " ", m.Year))
		} else {
			values.Set("", fmt.Sprint(m.Artist, " ", m.Album, " ", m.Year))
		}
		//fmt.Println(values)
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

		if i == 0 {
			if len(href) > 0 {
				//return by music name
				return href
			}
		} else {
			//return by album name
			return href
		}
	}
	return ""
}

func LeetMagnet(b []byte) string {
	pattern := regexp.MustCompile("<ul class=\"download-links\">(?P<data>(.|\n)+)Magnet Download")
	mp := map[string]string{}
	FindPattern(pattern, b, &mp)

	//fmt.Printf("torrent page content [%v]", mp["data"])

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

	//fmt.Printf("link [%v]", mp["href"])

	return mp["href"]
}
