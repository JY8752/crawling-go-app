package controller

import (
	"JY8752/crawling_app_rest/application/service"
	domain "JY8752/crawling_app_rest/domain/repository"
	form "JY8752/crawling_app_rest/presentation/form/crawling"
	"JY8752/crawling_app_rest/utils/json"
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
)

type (
	CrawlingController interface {
		Register(context.Context)
	}

	crawlingController struct {
		CrawlingService service.CrawlingService
	}
)

func NewCrawlingController(cr domain.CrawledUrl) CrawlingController {
	return &crawlingController{CrawlingService: service.NewCrawlingService(cr)}
}

func (cc *crawlingController) Register(ctx context.Context) {
	http.HandleFunc("/crawling", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/crawling/all", func(w http.ResponseWriter, r *http.Request) { crawlingAllHandle(ctx, cc.CrawlingService, w, r) })
	http.HandleFunc("/crawling/details/{:id}", func(w http.ResponseWriter, r *http.Request) {})
}

func crawlingAllHandle(ctx context.Context, service service.CrawlingService, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//クローリングしたサイトURLを全件取得する
		list := service.GetCrawledUrls(ctx)
		writeResponse(&form.GetCrawledUrls{CrawledUrls: list}, w)
	default:
		log.Printf("unkown http method %v\n", r.Method)
	}
}

func writeResponse(f *form.GetCrawledUrls, w http.ResponseWriter) string {
	var buf bytes.Buffer
	if err := json.JsonEncode(&buf, f); err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, buf.String())

	return buf.String()
}
