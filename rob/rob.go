package rob

import "math"

func rob(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dp[i] = Max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[length-1]
}

func Max(arr ...int) int {
	res := math.MinInt64
	for i := range arr {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}

func Min(arr ...int) int {
	res := math.MaxInt64
	for i := range arr {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}

func rob_2(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return Max(nums[1], nums[0])
	}
	return Max(help(nums[:len(nums)-1]), help(nums[1:]))
}

func help(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = Max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}
