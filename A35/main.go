package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
}

func main() {
	// input
	var n int
	fmt.Scanf("%d", &n)
	as := scanLineInt(sc, n, 1)

	// main
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, i+1)
	}
	copy(dp[n], as)

	// main
	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			if i%2 == 1 {
				dp[i][j] = max(dp[i+1][j], dp[i+1][j+1])
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1])
			}
		}
	}
	fmt.Println(dp[1][1])
}

func max(values ...int) int {
	var mx int
	for _, v := range values {
		if v > mx {
			mx = v
		}
	}
	return mx
}

func min(values ...int) int {
	mn := math.MaxInt64
	for _, v := range values {
		if v < mn {
			mn = v
		}
	}
	return mn
}

func scanLineInt(sc *bufio.Scanner, size, offset int) []int {
	items := make([]int, size+offset)
	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		items[i+offset] = atoi(s)
	}
	return items
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	mustNil(err)
	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func mustNil(err error) {
	if err != nil {
		panic(err)
	}
}
