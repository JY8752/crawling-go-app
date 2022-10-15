package data

import (
	"JY8752/crawling_app_batch/ent"
	"JY8752/crawling_app_batch/ent/crawledurl"
	"context"
	"time"
)

type (
	LinkUrl interface {
		Save(ctx context.Context, url, referer string, time *time.Time, cu *ent.CrawledUrl) *ent.LinkUrl
		FindByCrawledUrlId(ctx context.Context, id int) []*ent.LinkUrl
	}

	linkUrl struct {
		client *ent.Client
	}
)

func NewLinkUrl(c *ent.Client) LinkUrl {
	return &linkUrl{client: c}
}

func (l *linkUrl) Save(ctx context.Context, url, referer string, time *time.Time, cu *ent.CrawledUrl) *ent.LinkUrl {
	return l.client.LinkUrl.
		Create().
		AddBaseURL(cu).
		SetURL(url).
		SetReferer(referer).
		SetNillableUpdatedAt(time).
		SetNillableCreatedAt(time).
		SaveX(ctx)
}

func (l *linkUrl) FindByCrawledUrlId(ctx context.Context, id int) []*ent.LinkUrl {
	return l.client.CrawledUrl.Query().
		Where(crawledurl.ID(id)).
		QueryLinkUrls().
		AllX(ctx)
}
