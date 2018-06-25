package data

import (
	"encoding/json"
	"io"
)

// Story represents the whole story in gopher.json
type Story map[string]Chapter

// Chapter is a struct representing a detail of each chapter in gopher.json.
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option is a struct representing which chapter we can go next.
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// ReadStory read data in JSON format from given io.Reader
func ReadStory(r io.Reader) (Story, error) {
	var story Story

	d := json.NewDecoder(r)
	err := d.Decode(&story)
	return story, err

}
