package form

import "JY8752/crawling_app_rest/domain/model"

type GetCrawledUrls struct {
	CrawledUrls []*model.CrawledUrl `json:"crawled_urls"`
}

type GetLinkUrls struct {
	LinkUrls []*model.LinkUrl `json:"link_urls"`
}
