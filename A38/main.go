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
	var d, n int
	fmt.Scanf("%d %d", &d, &n)
	ls := make([]int, n+1)
	rs := make([]int, n+1)
	hs := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scanf("%d %d %d", &ls[i], &rs[i], &hs[i])
	}

	maxs := make([]int, d+1)
	for i := 1; i <= d; i++ {
		maxs[i] = 24
	}

	for i := 1; i <= n; i++ {
		l, r, h := ls[i], rs[i], hs[i]
		for j := l; j <= r; j++ {
			maxs[j] = min(maxs[j], h)
		}
	}
	s := sum(maxs...)
	fmt.Println(s)
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

func sum(values ...int) int {
	var a int
	for _, v := range values {
		a += v
	}
	return a
}
