package dp

import "testing"

type test struct {
	arrray []int
	target int
	output int
}

func TestFindTargetSum(t *testing.T) {
	var tableTest = []test{
		{[]int{0, 0, 0, 0, 0}, 0, 32},
		{[]int{8, 1, 2, 0}, 5, 2},
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{1}, 1, 1},
	}
	for _, tc := range tableTest {
		result := FindTargetSum(tc.arrray, tc.target)
		if result != tc.output {
			t.Errorf("Expected number of ways to be %d, but got %d\n", tc.output, result)
		}
	}
}
