package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
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
	sc.Scan()
	w := atoi(sc.Text())
	sc.Scan()
	l := atoi(sc.Text())
	sc.Scan()
	r := atoi(sc.Text())

	divisor := 1000000007

	xs := scanLineInt(sc, n, 1)
	xs = append(xs, w)
	if debug > 0 {
		fmt.Printf("xs: %+v\n", xs)
	}

	routes := make([]int, n+2)
	routes[0] = 1
	for i := 0; i <= n; i++ {
		x := xs[i]
		j := binarySearch(0, n, func(k int) int {
			if xs[k] < x+l {
				return 1
			}
			if x+r < xs[k] {
				return -1
			}
			return 0
		})
		for j2 := j; j2 <= n+1; j2++ {
			if xs[j2] < x+l {
				continue
			}
			if xs[j2] > x+r {
				break
			}
			routes[j2] += routes[i] % divisor
		}
		for j2 := j - 1; j2 >= 1; j2-- {
			if xs[j2] > x+r {
				continue
			}
			if xs[j2] < x+l {
				break
			}
			routes[j2] += routes[i] % divisor
		}
		if debug > 0 {
			fmt.Printf("routes: %+v\n", routes)
		}
	}
	fmt.Println(routes[n+1] % divisor)
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
