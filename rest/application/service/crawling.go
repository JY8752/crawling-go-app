package service

type (
	CrawlingService interface{}
	crawlingService struct{}
)

func NewCrawlingService() CrawlingService {
	return &crawlingService{}
}
