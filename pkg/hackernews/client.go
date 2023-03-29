package hackernews

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const defaultBaseURL string = "https://hacker-news.firebaseio.com/v0"

// Client is a Hacker News Client.
type Client struct {
	BaseURL string
}

// Item is a news story.
type Item struct {
	Title string
	URL   string
	Host  string
}

// TopStories gets all the top stories IDs.
func (c Client) TopStories(ctx context.Context) ([]int, error) {
	var ids []int

	err := c.request(ctx, "/topstories.json", &ids)
	if err != nil {
		return []int{}, err
	}

	return ids, nil
}

// Item gets a news item by ID.
func (c Client) Item(ctx context.Context, id int) (Item, error) {
	item := Item{}
	path := fmt.Sprintf("/item/%d.json", id)

	err := c.request(ctx, path, &item)
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

func (c Client) request(ctx context.Context, path string, target interface{}) error {
	if c.BaseURL == "" {
		c.BaseURL = defaultBaseURL
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, path)

	client := http.Client{}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, http.NoBody)
	if err != nil {
		return err
	}

	r, err := client.Do(request)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)

	err = dec.Decode(target)
	if err != nil {
		return err
	}

	return nil
}
