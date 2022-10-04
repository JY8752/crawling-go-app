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

func (ht *HtmlTokenizer) GetTokenizer() *html.Tokenizer {
	return html.NewTokenizer(bytes.NewReader(ht.html))
}

func (ht *HtmlTokenizer) GetStringHtml() string {
	return string(ht.html)
}
