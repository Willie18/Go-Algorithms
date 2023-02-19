package dp

// calculates the number of subsets with a given target
func FindSubsetsWithSumTarget(numbers []int, target int) int {

	var n = len(numbers)

	var array = make([][]int, n)

	for row := range array {
		array[row] = make([]int, target+1)
	}

	if numbers[0] <= target {
		array[0][numbers[0]] = 1
	}

	if numbers[0] == 0 {
		array[0][0] = 2
	} else {
		array[0][0] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < target+1; j++ {
			// if we do not take the current value
			notTaken := array[i-1][j]

			taken := 0

			// if we decide to take this value
			// j is our new target for each iteration
			if numbers[i] <= j {
				taken = array[i-1][j-numbers[i]]
			}

			array[i][j] = taken + notTaken
		}
	}

	return array[n-1][target]
}
