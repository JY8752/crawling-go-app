package tokenizer

import (
	"reflect"
	"testing"
)

func TestGetStringHtml(t *testing.T) {
	h := "<html><body><h1>test</h1></body></html>"
	tz := NewHtmlTokenizer([]byte(h))
	if r := tz.GetStringHtml(); r != h {
		t.Errorf("expecte %v, but %v", h, r)
	}
}

func TestExtractLinkUrl(t *testing.T) {
	cases := map[string]struct {
		html   string
		expect []string
	}{
		"when href=/user, result=[https://test.com/user]": {
			"<html><a href=\"/user\">link</a></html>",
			[]string{"https://test.com/user"}},
		"when href=[/user, /about], result=[https://test.com/user, https://test.com/about]": {
			"<html><a href=\"/user\">link</a><a href=\"/about\">link2</a></html>",
			[]string{"https://test.com/user", "https://test.com/about"}},
		"when href=[], result=[]": {"<html></html>", nil},
		"when inner a tag href=/user, result=[https://test.com/user]": {
			"<html><div><a href=\"/user\">link</a></div></html>",
			[]string{"https://test.com/user"}},
	}

	h := "https://test.com"

	for name, tt := range cases {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			tz := NewHtmlTokenizer([]byte(tt.html))
			r := tz.ExtractLinkUrl(h)
			if !reflect.DeepEqual(r, tt.expect) {
				t.Errorf("expect %v, but %v", tt.expect, r)
			}
		})
	}
}
