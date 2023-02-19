package dp

import "math"

// leetcode 983

func MinimumCostForTickets(days []int, costs [3]int) int {
	var dp = make([]int, len(days)+1)
	var ticketdays = [3]int{1, 7, 30}
	var n = len(days)

	for i := n - 1; i >= 0; i-- {
		dp[i] = math.MaxInt
		for d := 0; d < 3; d++ {
			j := i
			for j < n && days[j] < days[i]+ticketdays[d] {
				j += 1
			}

			value := 0
			if j <= n {
				value = dp[j]
			}

			dp[i] = min(dp[i], costs[d]+value)
		}
	}

	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
