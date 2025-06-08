package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	queueLen = 3
)

func TestNormalQueue(t *testing.T) {
	testCase := []struct {
		name   string
		queue  func(t *testing.T) *NormalQueue[string]
		res    func(t *testing.T) *NormalQueue[string]
		length int
		//after
	}{

		{
			name: "enqueue peek",
			queue: func(t *testing.T) *NormalQueue[string] {
				return nil
			},
			res:    nil,
			length: 0,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			queue := tc.queue(t)
			res := tc.res(t)
			assert.Equal(t, queue.len, res.len)
			assert.Equal(t, queue.data, res.data)
			for _ = range tc.length {
				q1, err1 := queue.Dequeue()
				r1, err2 := res.Dequeue()
				assert.NoError(t, err1)
				assert.NoError(t, err2)
				assert.Equal(t, q1, r1)
			}
		})
	}
}
