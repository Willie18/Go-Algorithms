package dp

import "testing"

func TestFindSubsets(t *testing.T) {

	var tableTest = []test{
		{[]int{1, 2, 3, 4, 5}, 6, 3},
	}

	for _, tc := range tableTest {
		result := FindSubsetsWithSumTarget(tc.arrray, tc.target)
		if result != tc.output {
			t.Errorf("Expected number of subsets to be %d, but got %d\n", tc.output, result)
		}
	}
}
