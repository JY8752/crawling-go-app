package crawling

import (
	"JY8752/crawling_app_batch/tokenizer"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Crawling(u string) ([]string, error) {
	//baseURLのコンテンツ取得
	res, err := http.Get(u)
	if err != nil {
		log.Printf("failed HTTP request err: %v\n", err.Error())
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed read response body err: %v\n", err.Error())
		return nil, err
	}

	//取得したコンテンツのTokenizerを作成
	ht := tokenizer.NewHtmlTokenizer(b)

	pu, err := url.Parse(u)
	if err != nil {
		log.Printf("failed parse url. url: %v err: %v\n", u, err.Error())
		return nil, err
	}

	//一旦、リンクURLかき集める
	urls := ht.ExtractLinkUrl(fmt.Sprintf("%v://%v", pu.Scheme, pu.Host))

	fmt.Printf("extract urls %v\n", urls)

	var result []string
	for _, uu := range urls {
		puu, err := url.Parse(uu)
		if err != nil {
			log.Printf("failed parse url. url: %v err: %v\n", uu, err.Error())
			continue
		}
		if puu.Host == pu.Host {
			result = append(result, uu)
		}
	}

	return result, nil
}
