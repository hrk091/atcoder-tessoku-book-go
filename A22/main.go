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
	bs := scanLineInt(sc, n, 1)

	// main
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

func scanLineInt(sc *bufio.Scanner, size, offset int) []int {
	items := make([]int, size+offset)
	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		items[i+offset], _ = strconv.Atoi(s)
	}
	return items
}
