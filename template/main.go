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
	sc.Split(bufio.ScanWords)
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
	var mn int
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

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func scanLineInt(sc *bufio.Scanner, size, offset int) []int {
	items := make([]int, size+offset)
	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		items[i+offset], _ = strconv.Atoi(s)
	}
	return items
}
