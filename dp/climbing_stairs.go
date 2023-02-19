package dp

// leetcode 70

// how many different ways can we get to stop

// this problem resembles the fibboncci  in reverse
func ClimbingStairs(stairs int) int {
	var (
		a = 1
		b = 1
	)
	for i := 0; i < stairs; i++ {
		a, b = b, a+b
	}

	return a
}
