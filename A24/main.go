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
	dp := map[int]int{}
	for i := 1; i <= n; i++ {
		var maxV int
		for k, v := range dp {
			if as[k] < as[i] {
				if maxV < v {
					maxV = v
				}
			}
		}
		dp[i] = maxV + 1
		//fmt.Println("---")
		//for k, v := range dp {
		//	fmt.Printf("k=%d, a=%d, len=%d\n", k, as[k], v)
		//}
	}

	var max int
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
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
