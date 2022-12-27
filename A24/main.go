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
	dp := make([]int, n+1)
	var LEN int
	for i := 1; i <= n; i++ {
		if as[i] <= dp[1] {
			dp[1] = as[i]
		} else if as[i] > dp[LEN] {
			LEN++
			dp[LEN] = as[i]
		} else {
			j := bs(1, LEN, func(pos int) int {
				if dp[pos+1] < as[i] {
					return 1
				}
				if as[i] <= dp[pos] {
					return -1
				}
				return 0
			})
			dp[j+1] = as[i]
		}
	}

	fmt.Println(LEN)
}

func bs(l, r int, fn func(int) int) int {
	// fn must be the one that returns true only when the result is greater than the given value.
	for l < r {
		mid := (l + r) / 2
		if ret := fn(mid); ret > 0 {
			l = mid + 1
		} else if ret < 0 {
			r = mid
		} else {
			return mid
		}
	}
	return l
}

func fillSlice(s []int, v int) {
	s[0] = v
	for p := 1; p < len(s); p *= 2 {
		copy(s[p:], s[:p])
	}
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
