package url

import (
	"net/url"
	"strings"
)

func NormalizeHrefUrl(href, domain string) (*url.URL, error) {
	u := href
	if !strings.HasPrefix(href, "http") {
		u = domain + href
	}

	//トレイリングスラッシュは除去しておく
	u = strings.TrimSuffix(u, "/")

	pu, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	return pu, nil
}
