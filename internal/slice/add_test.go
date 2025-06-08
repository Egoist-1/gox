package slice

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SliceTestSuite struct {
	suite.Suite
}

func (suite *SliceTestSuite) TestAdd() {
	testCase := []struct {
		name string
		src  []int
		ele  int
		idx  int
		res  []int
	}{
		{
			name: "添加到下标0",
			src:  []int{0, 1, 2, 3, 4},
			ele:  99,
			idx:  0,
			res:  []int{99, 0, 1, 2, 3, 4},
		},
		{
			name: "添加到中间",
			src:  []int{0, 1, 2, 3, 4},
			ele:  99,
			idx:  2,
			res:  []int{0, 1, 99, 2, 3, 4},
		},
	}
	for _, tc := range testCase {
		suite.Run(tc.name, func() {
			res, err := Add[int](tc.src, tc.ele, tc.idx)
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), tc.res, res)
		})
	}
}

func TestName(t *testing.T) {
}
func TestSliceTestSuite(t *testing.T) {
	suite.Run(t, new(SliceTestSuite))
}
