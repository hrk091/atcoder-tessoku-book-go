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
	var h, w int
	fmt.Scanf("%d %d", &h, &w)
	c := make([][]byte, h+1)
	c[0] = make([]byte, w+1)
	for i := 1; i <= h; i++ {
		var cc []byte
		fmt.Scan(&cc)
		c[i] = append([]byte{0}, cc...)
	}

	dp := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		dp[i] = make([]int, w+1)
	}

	for i := 1; i < h+1; i++ {
		for j := 1; j < w+1; j++ {
			if i == 1 && j == 1 {
				dp[1][1] = 1
				continue
			}
			if c[i-1][j] == 0 || string(c[i-1][j]) == "." {
				dp[i][j] += dp[i-1][j]
			}
			if c[i][j-1] == 0 || string(c[i][j-1]) == "." {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	fmt.Println(dp[h][w])
}
