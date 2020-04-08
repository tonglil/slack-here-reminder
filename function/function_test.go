package function

import (
	"fmt"
	"testing"
)

func TestInSet(t *testing.T) {
	var tests = []struct {
		item string
		set  []string
		want bool
	}{
		{"a", []string{}, false},
		{"a", []string{"a"}, true},
		{"a", []string{"b"}, false},
		{"a", []string{"a", "b"}, true},
		{"a-b", []string{"a-b"}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.item, tt.set)
		t.Run(testname, func(t *testing.T) {
			res := inSet(tt.item, tt.set)
			if res != tt.want {
				t.Errorf("got %v, want %v", res, tt.want)
			}
		})
	}
}
