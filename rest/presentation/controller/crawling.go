package controller

import "net/http"

type (
	CrawlingController interface {
		Register()
	}

	crawlingController struct {
	}
)

func NewCrawlingController() CrawlingController {
	return &crawlingController{}
}

func (cc *crawlingController) Register() {
	http.HandleFunc("/crawling", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/crawling/details/{:id}", func(w http.ResponseWriter, r *http.Request) {})
}
