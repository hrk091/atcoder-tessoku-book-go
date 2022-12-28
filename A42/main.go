package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
	sc.Scan()
	n := atoi(sc.Text())
	sc.Scan()
	k := atoi(sc.Text())

	aa := make([]int, n+1)
	bb := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sc.Scan()
		aa[i] = atoi(sc.Text())
		sc.Scan()
		bb[i] = atoi(sc.Text())
	}

	// main
	max := 0
	for i := 1; i <= 100-k+1; i++ {
		for j := 0; j <= 100-k+1; j++ {
			count := 0
			for p := 1; p <= n; p++ {
				a, b := aa[p], bb[p]
				if i <= a && a <= i+k && j <= b && b <= j+k {
					count++
				}
			}
			if count > max {
				max = count
			}
		}
	}
	fmt.Println(max)

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
