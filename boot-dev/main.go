package main

func sum(nums ...int) int {
	sumVal := 0
	for i := 0; i < len(nums); i++ {
		sumVal += nums[i]
	}

	return sumVal
}
