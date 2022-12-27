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
	var a int
	fmt.Scanf("%d", &a)

	// main
	var ok bool

	// output
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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

func bs(l, r int, fn func(int) bool) int {
	// fn must be the one that returns true only when the result is greater than the given value.
	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func scanLineInt(sc *bufio.Scanner, size, offset int) []int {
	items := make([]int, size+offset)
	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		items[i+offset] = atoi(s)
	}
	return items
}

func fillSlice(s []int, v int) {
	s[0] = v
	for p := 1; p < len(s); p *= 2 {
		copy(s[p:], s[:p])
	}
}

func fillMatrix(s [][]int, v int) {
	fillSlice(s[0], v)
	for p := 1; p < len(s); p++ {
		copy(s[p], s[0])
	}
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
