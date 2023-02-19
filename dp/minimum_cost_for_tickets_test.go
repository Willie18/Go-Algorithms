package dp

import "testing"

type ticketsTest struct {
	days   []int
	costs  [3]int
	output int
}

func TestMinimumCostForTickets(t *testing.T) {
	var tb = []ticketsTest{
		{[]int{1, 4, 6, 7, 8, 20}, [3]int{2, 7, 15}, 11},
	}

	for _, tc := range tb {
		result := MinimumCostForTickets(tc.days, tc.costs)
		if result != tc.output {
			t.Errorf("MinimumCostForTickets() returned %d, expected %d", result, tc.output)
		}
	}

}
