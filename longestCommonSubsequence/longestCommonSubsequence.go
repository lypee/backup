package main

import (
	"log"
	"math"
)

func main() {
	//log.Println(max(1,2))
	str1 := "abcde"
	str2 := "ace"
	log.Println(longestCommonSubsequence(str1, str2))
}

func max(arr ...int) int {
	res := math.MinInt64
	for _, num := range arr {
		if num > res {
			res = num
		}
	}
	return res
}


func longestCommonSubsequence2(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	dp := make([][]int, len1)
	for idx := range dp {
		dp[idx] = make([]int, len2)
	}

	for idx1, c1 := range text1 {
		for idx2, c2 := range text2 {
			if c1 == c2 {
				dp[idx1+1][idx2+1] = dp[idx1][idx2] + 1
			} else {
				dp[idx1+1][idx2+1] = max(dp[idx1+1][idx2], dp[idx1][idx2+1])
			}
		}
	}
	return dp[len1][len2]
}


func longestCommonSubsequence(text1 string, text2 string) int {
	t1Len, t2Len := len(text1), len(text2)
	dp := make([][]int, t1Len+1)
	for index := range dp {
		dp[index] = make([]int, t2Len+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[t1Len][t2Len]
}
