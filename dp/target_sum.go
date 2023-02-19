package dp

// leetcode 494

// you are given an integer array nums and an integer target

// build and expression out of nums by adding one of the symbols + and - beofre each integer

// in nums and then conctenate all the integers

// eg if nums = [2,1], you can add a + before 2 and a - before 1 and
// concatenate them to build the expression +2-1

// return the number of different expressions that you can build which evaluate to the target

// Input nums = [1,1,1,1,1]  Targe= 3

// Output = 5

// the are five ways to assign symbols to make the sum of nums be target 3

// -1 + 1 + 1 + 1 + 1 = 3

// +1 - 1 + 1 + 1 + 1 = 3

// +1 + 1 - 1 + 1 + 1 = 3

// +1 + 1 + 1 - 1 + 1 = 3

// +1 + 1 + 1 + 1 - 1 = 3

// initialize a two-d array rows = len(inputs) columns = (target +1 )

// this function calculates the number of subarrays that can add upto a certain target

func FindTargetSum(arr []int, target int) int {
	// handle the edge cases
	// if target sum is greater than total sum or theri difference(total-target) is not a mutlple of two just return

	var total int
	for i := range arr {
		total += arr[i]
	}

	if (total < int(target)) || (total-int(target))%2 != 0 {
		return 0
	}

	return FindSubsetsWithSumTarget(arr, (total-target)/2)
}
