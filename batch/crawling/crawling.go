package crawling

import (
	"JY8752/crawling_app_batch/tokenizer"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Crawling(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("failed HTTP request err: %v", err.Error())
		return
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed read response body err: %v", err.Error())
		return
	}

	ht := tokenizer.NewHtmlTokenizer(b)
	fmt.Println(ht.GetStringHtml())
}
