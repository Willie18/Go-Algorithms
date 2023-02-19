package dp

import "math"

// Find minimum number of coins that make a given value
// O(n * m) time complexity
// (m) space complexity
func MinimumCoins(coins []int, value int) int {
	if value <= 0 || len(coins) == 0 {
		return -1
	}

	var (
		dp = []int{0}
	)

	for i := 0; i < value+1; i++ {
		dp = append(dp, math.MaxInt)
	}

	for amount := 1; amount < value+1; amount++ {
		for _, coin := range coins {
			if amount-coin >= 0 {
				dp[amount] = int(math.Min(float64(dp[amount]), float64(1+dp[amount-coin])))
			}
		}
	}

	if dp[value] < math.MaxInt {
		return dp[value]
	}
	return -1
}
