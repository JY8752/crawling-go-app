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
	"strconv"

	"github.com/gorilla/mux"
)

type (
	CrawlingController interface {
		Register(context.Context, *mux.Router)
	}

	crawlingController struct {
		CrawlingService service.CrawlingService
	}
)

func NewCrawlingController(cr domain.CrawledUrl, l domain.LinkUrl) CrawlingController {
	return &crawlingController{CrawlingService: service.NewCrawlingService(cr, l)}
}

func (cc *crawlingController) Register(ctx context.Context, r *mux.Router) {
	r.HandleFunc("/crawling/all", func(w http.ResponseWriter, r *http.Request) { CrawlingAllHandle(ctx, cc.CrawlingService, w, r) })
	r.HandleFunc("/crawling/details/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) { GetLinkUrlsHandle(ctx, cc.CrawlingService, w, r) })
}

func CrawlingAllHandle(ctx context.Context, service service.CrawlingService, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//クローリングしたサイトURLを全件取得する
		list := service.GetCrawledUrls(ctx)
		writeResponse(&form.GetCrawledUrls{CrawledUrls: list}, w)
	default:
		log.Printf("unkown http method %v\n", r.Method)
	}
}

func GetLinkUrlsHandle(ctx context.Context, service service.CrawlingService, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vars := mux.Vars(r)

		param, ok := vars["id"]
		if !ok {
			log.Printf("id is missing in path parameters.\n")
			return
		}

		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("parameter is not int value.err: %v\n", err.Error())
		}

		list := service.GetLinkUrls(ctx, id)
		writeResponse(&form.GetLinkUrls{LinkUrls: list}, w)
	default:
		log.Printf("unkown http method %v\n", r.Method)
	}
}

func writeResponse(f any, w http.ResponseWriter) string {
	var buf bytes.Buffer
	if err := json.JsonEncode(&buf, f); err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, buf.String())

	return buf.String()
}
