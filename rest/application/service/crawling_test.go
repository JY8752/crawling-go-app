package service

import (
	"JY8752/crawling_app_rest/ent"
	"context"
	"testing"
	"time"
)

type CrawlingRepositoryMock struct {
	FakeFindAll func(ctx context.Context) []*ent.CrawledUrl
}

func (m CrawlingRepositoryMock) Save(ctx context.Context, url string, time *time.Time) (*ent.CrawledUrl, error) {
	return nil, nil
}

func (m CrawlingRepositoryMock) FindById(ctx context.Context, id int) *ent.CrawledUrl {
	return nil
}

func (m CrawlingRepositoryMock) FindAll(ctx context.Context) []*ent.CrawledUrl {
	return m.FakeFindAll(ctx)
}

type LinkUrlRepositoryMock struct {
	FakeFindByCrawledUrlId func(ctx context.Context, id int) []*ent.LinkUrl
}

func (m LinkUrlRepositoryMock) Save(ctx context.Context, url, referer string, time *time.Time, cu *ent.CrawledUrl) *ent.LinkUrl {
	return nil
}

func (m LinkUrlRepositoryMock) FindByCrawledUrlId(ctx context.Context, id int) []*ent.LinkUrl {
	return m.FakeFindByCrawledUrlId(ctx, id)
}

func TestGetCrawledUrls(t *testing.T) {
	//given
	ctx := context.Background()
	testTime := time.Date(2022, 10, 19, 0, 0, 0, 0, time.Local)
	service := NewCrawlingService(
		CrawlingRepositoryMock{
			FakeFindAll: func(ctx context.Context) []*ent.CrawledUrl {
				return []*ent.CrawledUrl{{ID: 1, URL: "https://test.com", CreatedAt: testTime}}
			}},
		LinkUrlRepositoryMock{
			FakeFindByCrawledUrlId: func(ctx context.Context, id int) []*ent.LinkUrl {
				return nil
			},
		},
	)

	//when
	r := service.GetCrawledUrls(ctx)

	//then
	if len(r) != 1 {
		t.Errorf("expect result length 1, but %v\n", len(r))
	}
	if r[0].Id != 1 {
		t.Errorf("expected crawled_url id is 1, but %v\n", r[0].Id)
	}
	if r[0].Url != "https://test.com" {
		t.Errorf("expected crawled_url url is https://test.com, but %v\n", r[0].Url)
	}
	if r[0].CreatedAt != testTime {
		t.Errorf("expected crawled_url created_at is %v, but %v\n", testTime, r[0].CreatedAt)
	}
}

func TestGetLinkUrls(t *testing.T) {
	//given
	ctx := context.Background()
	s := NewCrawlingService(
		CrawlingRepositoryMock{},
		LinkUrlRepositoryMock{
			FakeFindByCrawledUrlId: func(ctx context.Context, id int) []*ent.LinkUrl {
				return []*ent.LinkUrl{
					{ID: 1, URL: "https://test.com/page1", Referer: "https://test.com"},
				}
			},
		},
	)

	//when
	r := s.GetLinkUrls(ctx, 1)

	//then
	if len(r) != 1 {
		t.Errorf("expect result length 1, but %v\n", len(r))
	}
	if r[0].Id != 1 {
		t.Errorf("expected id is 1, but %v\n", r[0].Id)
	}
	if r[0].Url != "https://test.com/page1" {
		t.Errorf("expected url is https://test.com/page1, but %v\n", r[0].Url)
	}
	if r[0].Referer != "https://test.com" {
		t.Errorf("expected referer is https://test.com, but %v\n", r[0].Referer)
	}
}
