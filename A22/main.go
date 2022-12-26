package main

import (
	"fmt"
	"time"
)

func max(values ...int) int {
	var mx int
	for _, v := range values {
		if v > mx {
			mx = v
		}
	}
	return mx
}

func main() {
	// input
	var n int
	fmt.Scanf("%d", &n)
	as := make([]int, n+1)
	bs := make([]int, n+1)
	for i := 1; i <= n-1; i++ {
		fmt.Scan(&as[i])
	}
	for i := 1; i <= n-1; i++ {
		fmt.Scan(&bs[i])
	}

	// main
	t1 := time.Now()
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i != 1 && dp[i] == 0 {
			continue
		}
		//fmt.Printf("%d %v\n", i, dp)
		if i == n {
			break
		}
		a := as[i]
		dp[a] = max(dp[a], dp[i]+100)
		b := bs[i]
		dp[b] = max(dp[b], dp[i]+150)
	}
	fmt.Println(dp[n])
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
