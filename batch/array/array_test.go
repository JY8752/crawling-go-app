package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	cases := map[string]struct {
		arr    []string
		item   string
		expect bool
	}{
		fmt.Sprintf("when arr: %v item: %v, return %v\n", []string{"item1", "item2"}, "item1", true):  {[]string{"item1", "item2"}, "item1", true},
		fmt.Sprintf("when arr: %v item: %v, return %v\n", []string{"item1", "item2"}, "item3", false): {[]string{"item1", "item2"}, "item3", false},
		fmt.Sprintf("when arr: %v item: %v, return %v\n", []string{}, "item3", false):                 {[]string{}, "item3", false},
	}

	for name, tt := range cases {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := Contains(tt.arr, tt.item)
			if r != tt.expect {
				t.Errorf("expect %v, but %v", tt.expect, r)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	cases := map[string]struct {
		arr   []string
		item  string
		resut []string
	}{
		"when remove item2 from [item1, item2], result [item1]":        {[]string{"item1", "item2"}, "item2", []string{"item1"}},
		"when remove item3 from [item1, item2], result [item1, item2]": {[]string{"item1", "item2"}, "item3", []string{"item1", "item2"}},
		"when remove item3 from [], result []":                         {nil, "item3", nil},
	}

	for name, tt := range cases {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := Remove(tt.arr, tt.item)
			if !reflect.DeepEqual(tt.resut, r) {
				t.Errorf("expecte %v, but %v\n", tt.resut, r)
			}
		})
	}
}
