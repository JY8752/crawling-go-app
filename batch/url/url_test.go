package url

import (
	"fmt"
	"testing"
)

func TestNormalizeHrefUrl(t *testing.T) {
	top := "https://test.com"
	cases := map[string]struct {
		href   string
		result string
	}{
		normalizeHrefUrlTestName("https://test.com", "https://test.com"):  {"https://test.com", "https://test.com"},  //トップページへのリンク
		normalizeHrefUrlTestName("https://test.com/", "https://test.com"): {"https://test.com/", "https://test.com"}, //トレイリングスラッシュ付きのページ
		normalizeHrefUrlTestName("/user", "https://test.com/user"):        {"/user", "https://test.com/user"},        //別ページへのリンク
		normalizeHrefUrlTestName("#", "https://test.com"):                 {"#", "https://test.com"},                 //#
	}

	for name, tt := range cases {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r, err := NormalizeHrefUrl(tt.href, top)
			if err != nil {
				t.Errorf("err: %v\n", err.Error())
			}
			if r.String() != tt.result {
				t.Errorf("expect %v, but %v\n", tt.result, r.String())
			}
		})
	}
}

func normalizeHrefUrlTestName(href, result string) string {
	return fmt.Sprintf("when href: %v, result is %v\n", href, result)
}
