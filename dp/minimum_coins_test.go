package dp

import "testing"

func TestMinimumCoins(t *testing.T) {
	var testcases = []test{
		{[]int{10, 5, 15}, 10, 1},
		{[]int{5, 5, 15, 8, 1, 1}, 10, 2},
		{[]int{5, 5, 15, 8, 4}, 3, -1},
		{[]int{}, 3, -1},
	}

	for _, tc := range testcases {
		result := MinimumCoins(tc.arrray, tc.target)
		if result != tc.output {
			t.Errorf("Expected number of coints to be %d, but got %d\n", tc.output, result)
		}

	}

}
