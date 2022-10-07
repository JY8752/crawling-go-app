package tokenizer

import (
	"bytes"

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

func (ht *HtmlTokenizer) ExtractLinkUrl() (urls []string) {
	t := ht.getTokenizer()
	for {
		if t.Next() == html.ErrorToken {
			return urls
		}
		tk := t.Token()
		if tk.Data == "a" && tk.Type.String() == "StartTag" {
			for _, attr := range tk.Attr {
				if attr.Key == "href" {
					urls = append(urls, attr.Val)
				}
			}
		}
	}
}

func (ht *HtmlTokenizer) getTokenizer() *html.Tokenizer {
	return html.NewTokenizer(bytes.NewReader(ht.html))
}
