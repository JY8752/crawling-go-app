package service

import (
	"JY8752/crawling_app_rest/domain/model"
	domain "JY8752/crawling_app_rest/domain/repository"
	"context"
)

type (
	CrawlingService interface {
		GetCrawledUrls(ctx context.Context) []*model.CrawledUrl
		GetLinkUrls(ctx context.Context, id int) []*model.LinkUrl
	}

	crawlingService struct {
		CrawlingRepository domain.CrawledUrl
		LinkUrlRepository  domain.LinkUrl
	}
)

func NewCrawlingService(crawlingRepository domain.CrawledUrl, linkUrlRepository domain.LinkUrl) CrawlingService {
	return &crawlingService{CrawlingRepository: crawlingRepository, LinkUrlRepository: linkUrlRepository}
}

func (cs *crawlingService) GetCrawledUrls(ctx context.Context) []*model.CrawledUrl {
	var list []*model.CrawledUrl
	for _, cu := range cs.CrawlingRepository.FindAll(ctx) {
		list = append(list, &model.CrawledUrl{
			Id:        cu.ID,
			Url:       cu.URL,
			CreatedAt: cu.CreatedAt,
		})
	}
	return list
}

func (cs *crawlingService) GetLinkUrls(ctx context.Context, id int) []*model.LinkUrl {
	var list []*model.LinkUrl
	for _, l := range cs.LinkUrlRepository.FindByCrawledUrlId(ctx, id) {
		list = append(list, &model.LinkUrl{
			Id:      l.ID,
			Url:     l.URL,
			Referer: l.Referer,
		})
	}
	return list
}
