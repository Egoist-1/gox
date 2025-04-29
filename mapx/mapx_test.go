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

func TestContain(t *testing.T) {
	type User struct {
		Name string
	}
	testCase := []struct {
		Name string
		src  map[string]User
		fn   func(val User) bool
		res  string
		res2 bool
	}{
		{
			Name: "success",
			src: map[string]User{
				"AA": User{Name: "aa"},
				"BB": User{Name: "bb"},
			},
			fn: func(val User) bool {
				return val.Name == "bb"
			},
			res:  "BB",
			res2: true,
		},
		{
			Name: "失败",
			src: map[string]User{
				"AA": User{Name: "aa"},
				"BB": User{Name: "bb"},
			},
			fn: func(val User) bool {
				return val.Name == "cc"
			},
			res:  "",
			res2: false,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			key, ok := Contain[map[string]User](tc.src, tc.fn)
			assert.Equal(t, tc.res, key)
			assert.Equal(t, tc.res2, ok)
		})
	}
}
