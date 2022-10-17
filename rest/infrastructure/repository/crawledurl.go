package repository

import (
	domain "JY8752/crawling_app_rest/domain/repository"
	"JY8752/crawling_app_rest/ent"
	"context"
	"time"
)

type crawledUrl struct {
	client *ent.Client
}

func NewCrawledUrl(c *ent.Client) domain.CrawledUrl {
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

func (cu *crawledUrl) FindAll(ctx context.Context) []*ent.CrawledUrl {
	return cu.client.CrawledUrl.Query().AllX(ctx)
}
