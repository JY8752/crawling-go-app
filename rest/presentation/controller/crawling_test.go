package controller

import (
	"JY8752/crawling_app_rest/domain/model"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type CrawlingServiceMock struct {
	FakeGetCrawledUrls func(context.Context) []*model.CrawledUrl
}

func (m *CrawlingServiceMock) GetCrawledUrls(ctx context.Context) []*model.CrawledUrl {
	return m.FakeGetCrawledUrls(ctx)
}

type CrawlingAllHandleResult struct {
	CrawledUrls []*model.CrawledUrl `json:"crawled_urls"`
}

func TestCrawlingAllHandler(t *testing.T) {
	testTime := time.Date(2022, 10, 15, 0, 0, 0, 0, time.Local)

	cases := map[string]struct {
		service CrawlingServiceMock
		expect  CrawlingAllHandleResult
	}{
		"return two record": {
			service: CrawlingServiceMock{
				FakeGetCrawledUrls: func(ctx context.Context) []*model.CrawledUrl {
					return []*model.CrawledUrl{
						{Id: 1, Url: "https://test.com1", CreatedAt: testTime},
						{Id: 2, Url: "https://test.com2", CreatedAt: testTime},
					}
				},
			},
			expect: CrawlingAllHandleResult{CrawledUrls: []*model.CrawledUrl{
				{Id: 1, Url: "https://test.com1", CreatedAt: testTime},
				{Id: 2, Url: "https://test.com2", CreatedAt: testTime},
			}},
		},
		"return empty": {
			service: CrawlingServiceMock{
				FakeGetCrawledUrls: func(ctx context.Context) []*model.CrawledUrl {
					return []*model.CrawledUrl{}
				},
			},
			expect: CrawlingAllHandleResult{CrawledUrls: []*model.CrawledUrl{}},
		},
	}

	for name, tt := range cases {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			//given
			ctx := context.Background()
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "https://test", &bytes.Buffer{})

			//when
			CrawlingAllHandle(ctx, &tt.service, w, r)

			//then
			if w.Code != http.StatusOK {
				t.Errorf("expect %v, but %v\n", http.StatusOK, w.Code)
			}

			var b CrawlingAllHandleResult
			dec := json.NewDecoder(w.Result().Body)
			if err := dec.Decode(&b); err != nil {
				t.Errorf("failed parse body. err: %v\n", err.Error())
			}

			//ポインタのスライスをプロパティに持つ構造体の比較ができない...
			// if !reflect.DeepEqual(&b, &tt.expect) {
			// 	t.Errorf("expect %v, but %v\n", tt.expect, b)
			// }
			assertCrawlingAllHandleResult(t, tt.expect, b)
		})
	}
}

func assertCrawlingAllHandleResult(t *testing.T, expect, act CrawlingAllHandleResult) {
	t.Helper()

	if len(expect.CrawledUrls) != len(act.CrawledUrls) {
		t.Errorf("expect crawled_urls length %v, act %v\n", len(expect.CrawledUrls), len(act.CrawledUrls))
	}

	for i := 0; i < len(expect.CrawledUrls); i++ {
		expect := expect.CrawledUrls[i]
		act := act.CrawledUrls[i]

		if expect.Id != act.Id {
			t.Errorf("[index %d]expect id %v, but %v\n", i, expect.Id, act.Id)
		}

		if expect.Url != act.Url {
			t.Errorf("[index %d]expect url %v, but %v\n", i, expect.Url, act.Url)
		}

		if expect.CreatedAt != act.CreatedAt {
			t.Errorf("[index %d]expect created_at %v, but %v\n", i, expect.CreatedAt, act.CreatedAt)
		}
	}
}
