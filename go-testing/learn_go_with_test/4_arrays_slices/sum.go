package main

func Sum(nums [5]int) int {
	sum := 0
	// for i := 0; i < len(nums); i++ {
	// 	sum += nums[i]
	// }

	for _, value := range nums {
		sum += value
	}
	return sum
}

func Sum2(nums []int) int {
	sum := 0

	for _, value := range nums {
		sum += value
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	// lengthOfNums := len(numsToSum)
	// sums := make([]int, lengthOfNums)

	var sums []int

	// for i, nums := range numsToSum {
	// 	sums[i] = Sum2(nums)
	// }

	// Refactor

	for _, nums := range numsToSum {
		sums = append(sums, Sum2(nums))
	}

	return sums
}

func SumAllTails(numsToSum ...[]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tails := nums[1:]
			sums = append(sums, Sum2(tails))
		}
	}

	return sums
}
