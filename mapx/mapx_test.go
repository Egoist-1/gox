package mapx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	testCase := []struct {
		Name string
		src  map[string]string
		fn   func(key, val string) string
		res  []string
	}{
		{
			Name: "success",

			src: map[string]string{
				"1": "1",
				"2": "2",
			},
			fn: func(key, val string) string {
				return key + val
			},
			res: []string{
				"11", "22",
			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			slice := ToSlice[map[string]string, string](tc.src, tc.fn)
			assert.Equal(t, tc.res, slice)
		})
	}
}
