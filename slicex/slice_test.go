package slicex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	testCase := []struct {
		Name string
		fn   func(src string) string
		src  []string
		dst  []string
	}{
		{
			Name: "success",
			fn: func(src string) string {
				return src + "dst"
			},
			src: []string{"1", "2", "3"},
			dst: []string{"1dst", "2dst", "3dst"},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			result := Map[string, string](tc.src, tc.fn)
			assert.Equal(t, tc.dst, result)
		})
	}
}
