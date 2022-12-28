package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
}

func main() {
	// input
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)
	dp := make([]bool, 100009)

	// main
	for i := 0; i <= n; i++ {
		if i < a && i < b {
			dp[i] = false
			continue
		}
		var winnable bool
		if i < b {
			winnable = !dp[i-a]
		} else {
			winnable = !dp[i-a] || !dp[i-b]
		}
		if winnable {
			dp[i] = true
		} else {
			dp[i] = false
		}
	}
	if dp[n] {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
