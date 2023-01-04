package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	debug int
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	flag.Parse()
	d := flag.Arg(0)
	if d != "" {
		debug = atoi(d)
	}
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())
	cs := scanLineInt(sc, n, 1)

	sumAr := make([]int, n+1)
	sort.Slice(cs, func(i, j int) bool {
		return cs[i] < cs[j]
	})

	for i := 1; i <= n; i++ {
		sumAr[i] = sumAr[i-1] + cs[i]
	}

	sc.Scan()
	q := atoi(sc.Text())
	for i := 0; i < q; i++ {
		sc.Scan()
		x := atoi(sc.Text())
		if x > sumAr[n] {
			fmt.Println(n)
			continue
		}
		j := binarySearch(0, n-1, func(j int) int {
			if sumAr[j] <= x && x < sumAr[j+1] {
				return 0
			}
			if x < sumAr[j] {
				return -1
			}
			return 1
		})
		fmt.Println(j)
	}
}

func binarySearch(l, r int, fn func(int) int) int {
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

func scanLineInt(sc *bufio.Scanner, size, offset int) []int {
	items := make([]int, size+offset)
	for i := 0; i < size; i++ {
		sc.Scan()
		items[i+offset] = atoi(sc.Text())
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
