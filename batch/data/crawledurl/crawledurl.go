package data

import (
	"JY8752/crawling_app_batch/ent"
	"context"
	"time"
)

type (
	CrawledUrl interface {
		Save(ctx context.Context, url string, time *time.Time) (*ent.CrawledUrl, error)
		FindById(ctx context.Context, id int) *ent.CrawledUrl
	}

	crawledUrl struct {
		client *ent.Client
	}
)

func NewCrawledUrl(c *ent.Client) CrawledUrl {
	return &crawledUrl{client: c}
}

func (cu *crawledUrl) Save(ctx context.Context, url string, time *time.Time) (*ent.CrawledUrl, error) {
	return cu.client.CrawledUrl.Create().
		SetURL(url).
		SetNillableUpdatedAt(time).
		SetNillableCreatedAt(time).
		Save(ctx)
}

func (cu *crawledUrl) FindById(ctx context.Context, id int) *ent.CrawledUrl {
	return cu.client.CrawledUrl.GetX(ctx, id)
}
