package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gocolly/colly/v2"
)

var src_url = "https://www.songlyrics.com/alkaline-trio/"
var songs = []string{"time-to-waste-lyrics", "the-poison-lyrics", "burn-lyrics", "mercy-me-lyrics", "sadie-lyrics", "fall-victim-lyrics", "i-was-a-prayer-lyrics", "prevent-this-tragedy-lyrics", "back-to-hell-lyrics", "your-neck-lyrics", "smoke-lyrics"}
var lyrics []string

func main() {
	fmt.Println("Start parsing lyrics")
	for _, song_id := range songs {
		lyrics = append(lyrics, fixLyrics(getLyrics(src_url+song_id))...)
	}

	fmt.Println("Lyrics parsed successfully")

	fmt.Println("Starting server on http://127.0.0.1:8000 ")

	desu()
}

func getLyrics(url string) string {
	var lyric string

	c := colly.NewCollector()

	c.OnHTML("p[id=songLyricsDiv]", func(h *colly.HTMLElement) {
		lyric = h.Text /* Get text from html */
	})

	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			fmt.Printf("Can't access %s ", url)
			return
		}
	})

	c.Visit(url)

	return lyric
}

func fixLyrics(a string) []string {
	noempty := strings.Replace(a, "\n\n", "\n", 1)
	inline := strings.Split(noempty, "\n")
	return inline
}

/* Simple web server */
func desu() {
	http.HandleFunc("/", desune)
	http.ListenAndServe(":8000", nil)
}

func desune(w http.ResponseWriter, _ *http.Request) {
	lyric_length := rand.Intn(40-10) + 10
	for i := 0; i < lyric_length; i++ {
		fmt.Fprintf(w, "%s\n", lyrics[rand.Intn(len(lyrics))])
	}
	fmt.Fprintf(w, "\n\nCopyright (c) exathedev")
}
