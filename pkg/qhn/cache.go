package qhn

import (
	"context"
	"fmt"

	"github.com/aaronellington/quiet-hacker-news/pkg/hackernews"
)

func (app *App) updateCacheTick(ctx context.Context) {
	app.logger.Info(ctx, "Updating the Cache", nil)

	if err := app.updateCache(ctx); err != nil {
		app.logger.Error(ctx, "Error Updating the Cache", map[string]interface{}{
			"error": err.Error(),
		})
	}
}

func (app *App) updateCache(ctx context.Context) error {
	hackernewsItems, err := app.fetchHackerNewsItems(ctx)
	if err != nil {
		return err
	}

	app.hackernewsItems = hackernewsItems

	return nil
}

func (app *App) fetchHackerNewsItems(ctx context.Context) ([]hackernews.Item, error) {
	storyIDs, err := app.hackerNewsAPI.TopStories(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting top story IDs: %w", err)
	}

	items := []hackernews.Item{}

	for _, storyID := range storyIDs {
		item, err := app.hackerNewsAPI.Item(ctx, storyID)
		if err != nil {
			return nil, fmt.Errorf("error getting story #%d: %w", storyID, err)
		}

		if item.URL == "" || item.Title == "" {
			continue
		}

		items = append(items, item)

		if len(items) >= app.config.PageSize {
			break
		}
	}

	return items, nil
}
