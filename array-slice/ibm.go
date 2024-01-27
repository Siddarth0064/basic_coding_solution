package main

import (
	"fmt"
	"math"
)

func getMinCost(n, k int, plans [][]int) int64 {
	const inf = math.MaxInt64
	dp := make([]int64, n+1)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0

	for _, plan := range plans {
		l, r, c, p := plan[0], plan[1], plan[2], plan[3]
		for j := n; j >= l; j-- {
			for used := 1; used <= c; used++ {
				if j-used >= 0 {
					dp[j] = min(dp[j], dp[j-used]+int64(used)*p)
				}
			}
		}
	}

	result := inf
	for i := n; i >= n-k+1; i-- {
		result = min(result, dp[i])
	}
	return result
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Sample testcase 1
	plans1 := [][]int{{1, 3, 5, 2}, {1, 4, 5, 3}, {2, 5, 10, 1}}
	output1 := getMinCost(5, 7, plans1)
	fmt.Println(output1) // Output: 44

	// Sample testcase 2
	plans2 := [][]int{{1, 1, 4, 5}}
	output2 := getMinCost(4, 4, plans2)
	fmt.Println(output2) // Output: 20
}
