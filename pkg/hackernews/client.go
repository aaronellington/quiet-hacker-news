package hackernews

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const defaultBaseURL string = "https://hacker-news.firebaseio.com/v0"

// Client is a Hacker News Client
type Client struct {
	BaseURL string
}

// Item is a news story
type Item struct {
	Title string
	URL   string
	Host  string
}

// TopStories gets all the top stories IDs
func (c Client) TopStories() ([]int, error) {
	var ids []int

	err := c.request("/topstories.json", &ids)
	if err != nil {
		return []int{}, err
	}

	return ids, nil
}

// Item gets a news item by ID
func (c Client) Item(id int) (Item, error) {
	item := Item{}
	path := fmt.Sprintf("/item/%d.json", id)

	err := c.request(path, &item)
	if err != nil {
		return Item{}, err
	}

	parsedURL, err := url.Parse(item.URL)
	if err != nil {
		return Item{}, err
	}

	item.Host = strings.TrimPrefix(parsedURL.Hostname(), "www.")

	return item, nil
}

func (c Client) request(path string, target interface{}) error {
	if c.BaseURL == "" {
		c.BaseURL = defaultBaseURL
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, path)
	r, err := http.Get(fullURL)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(r.Body)
	err = dec.Decode(target)
	if err != nil {
		return err
	}

	return nil
}
