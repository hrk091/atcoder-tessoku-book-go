package main

import (
	"fmt"
)

func scanSeq(n int) []int {
	items := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&items[i])
	}
	return items
}

func main() {
	// input
	var n, s int
	fmt.Scanf("%d %d", &n, &s)
	as := scanSeq(n)

	partialSum := func(cur []bool, val int) []bool {
		next := make([]bool, s+1)
		for i := 0; i <= s; i++ {
			v := cur[i]
			if !v {
				continue
			}
			next[i] = true
			if i+val <= s {
				next[i+val] = true
			}
		}
		return next
	}

	dp := make([][]bool, n+1)
	dp[0] = make([]bool, s+1)
	dp[0][0] = true

	var ok bool
	for i := 1; i <= n; i++ {
		dp[i] = partialSum(dp[i-1], as[i-1])
		if dp[i][s] {
			ok = true
			break
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
