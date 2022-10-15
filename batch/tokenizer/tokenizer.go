package tokenizer

import (
	"bytes"
	"fmt"

	"JY8752/crawling_app_batch/url"

	"golang.org/x/net/html"
)

type HtmlTokenizer struct {
	html []byte
}

func NewHtmlTokenizer(h []byte) *HtmlTokenizer {
	return &HtmlTokenizer{html: h}
}

func (ht *HtmlTokenizer) GetStringHtml() string {
	return string(ht.html)
}

func (ht *HtmlTokenizer) ExtractLinkUrl(host string) (urls []string) {
	t := ht.getTokenizer()
	for {
		if t.Next() == html.ErrorToken {
			return urls
		}
		tk := t.Token()
		if tk.Data == "a" && tk.Type.String() == "StartTag" {
			for _, attr := range tk.Attr {
				if attr.Key == "href" {
					u := attr.Val
					nu, err := url.NormalizeHrefUrl(u, host)

					if err != nil {
						fmt.Printf("remove failed normalize url. url: %v\n", u)
						continue
					}

					//TOP画面除外
					if nu.String() == host {
						continue
					}

					urls = append(urls, nu.String())
				}
			}
		}
	}
}

func (ht *HtmlTokenizer) getTokenizer() *html.Tokenizer {
	return html.NewTokenizer(bytes.NewReader(ht.html))
}
