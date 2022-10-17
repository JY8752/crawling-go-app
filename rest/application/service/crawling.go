package service

import (
	"JY8752/crawling_app_rest/domain/model"
	domain "JY8752/crawling_app_rest/domain/repository"
	"context"
)

type (
	CrawlingService interface {
		GetCrawledUrls(ctx context.Context) []*model.CrawledUrl
	}

	crawlingService struct {
		CrawlingRepository domain.CrawledUrl
	}
)

func NewCrawlingService(crawlingRepository domain.CrawledUrl) CrawlingService {
	return &crawlingService{CrawlingRepository: crawlingRepository}
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
