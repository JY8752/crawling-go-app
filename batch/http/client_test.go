package http

import "testing"

func TestGet(t *testing.T) {
	cases := map[string]struct {
		url    string
		expect bool
	}{
		"success": {"https://example.com", true},
		"failed":  {"https://xxxxxx", false},
	}

	for name, tt := range cases {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			_, err := NewClient().Get(tt.url)
			if err != nil && tt.expect {
				t.Errorf("failed get request error: %v\n", err.Error())
			}
		})
	}
}
