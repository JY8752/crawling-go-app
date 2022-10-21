package domain

import (
	"JY8752/crawling_app_rest/ent"
	"context"
	"time"
)

type LinkUrl interface {
	Save(ctx context.Context, url, referer string, time *time.Time, cu *ent.CrawledUrl) *ent.LinkUrl
	FindByCrawledUrlId(ctx context.Context, id int) []*ent.LinkUrl
}
