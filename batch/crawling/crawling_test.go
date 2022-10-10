package crawling

import (
	"reflect"
	"testing"
)

//mock
type MockClient struct {
	MockGet func(u string) ([]byte, error)
}

func (mc MockClient) Get(u string) ([]byte, error) {
	return mc.MockGet(u)
}

func TestCrawling(t *testing.T) {
	cases := map[string]struct {
		mockClient MockClient
		expected   []string
	}{
		"success": {MockClient{MockGet: func(u string) ([]byte, error) {
			return []byte("<html><body><a href=\"/user\">link</a></body></html>"), nil
		}}, []string{"https://test.com/user"}},
		"other domain": {MockClient{MockGet: func(u string) ([]byte, error) {
			return []byte("<html><body><a href=\"https://test.co.jp\">link</a></body></html>"), nil
		}}, nil},
	}

	for name, tt := range cases {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			c := NewClawler(tt.mockClient)
			urls, err := c.Crawling("https://test.com")
			if err != nil {
				t.Errorf("failed crawling. error: %v\n", err.Error())
			}
			if !reflect.DeepEqual(urls, tt.expected) {
				t.Errorf("expected %v, but %v\n", tt.expected, urls)
			}
		})
	}
}
