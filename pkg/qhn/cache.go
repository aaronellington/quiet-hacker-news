package qhn

import (
	"fmt"
	"log"
	"time"

	"github.com/aaronellington/quiet-hacker-news/pkg/hackernews"
)

// UpdateCache updates the cache
func (app *App) UpdateCache() error {
	newItems, err := app.pullCacheItems()
	if err != nil {
		return err
	}

	if len(newItems) == 0 {
		return fmt.Errorf("A list of 0 items found. Not using it")
	}

	app.itemCache = newItems

	return nil
}

// Update the cache immediately and then once per interval
func (app *App) startCacheUpdateLoop() {
	app.cacheUpdateTick()
	for range time.NewTicker(app.refreshInterval).C {
		app.cacheUpdateTick()
	}
}

// Update the cache with consistent error handling
func (app *App) cacheUpdateTick() {
	err := app.UpdateCache()
	if err != nil {
		err = fmt.Errorf("Error Updating the Cache: %w", err)
		log.Print(err)
	}
}

func (app *App) pullCacheItems() ([]hackernews.Item, error) {
	storyIDs, err := app.hackerNewsAPI.TopStories()
	if err != nil {
		return nil, fmt.Errorf("Error getting top story IDs: %w", err)
	}

	items := []hackernews.Item{}
	for _, storyID := range storyIDs {
		item, err := app.hackerNewsAPI.Item(storyID)
		if err != nil {
			return nil, fmt.Errorf("Error getting story #%d: %w", storyID, err)
		}

		if item.URL == "" || item.Title == "" {
			continue
		}

		items = append(items, item)

		if len(items) >= app.pageSize {
			break
		}
	}

	return items, nil
}
