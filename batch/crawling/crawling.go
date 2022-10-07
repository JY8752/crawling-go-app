package crawling

import (
	"JY8752/crawling_app_batch/tokenizer"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Crawling(url string) ([]string, error) {
	//baseURLのコンテンツ取得
	res, err := http.Get(url)
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

	//一旦、リンクURLかき集める
	urls := ht.ExtractLinkUrl()

	d, err := getDomain(url)
	if err != nil {
		log.Printf("failed parse url url: %v err: %v\n", url, err.Error())
		return nil, err
	}

	var result []string
	for _, url := range urls {
		dd, err := getDomain(url)
		if err != nil {
			log.Printf("failed parse url url: %v err: %v\n", url, err.Error())
			continue
		}
		if dd == d {
			result = append(result, url)
		}
	}

	return result, nil
}

func getDomain(u string) (string, error) {
	uu, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	return uu.Host, nil
}
