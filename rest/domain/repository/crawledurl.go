package domain

import (
	"JY8752/crawling_app_rest/ent"
	"context"
	"time"
)

type CrawledUrl interface {
	Save(ctx context.Context, url string, time *time.Time) (*ent.CrawledUrl, error)
	FindById(ctx context.Context, id int) *ent.CrawledUrl
}
