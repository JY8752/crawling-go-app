package repository

import (
	domain "JY8752/crawling_app_rest/domain/repository"
	"JY8752/crawling_app_rest/ent"
	"JY8752/crawling_app_rest/ent/crawledurl"
	"context"
	"time"
)

type linkUrl struct {
	client *ent.Client
}

func NewLinkUrl(c *ent.Client) domain.LinkUrl {
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
