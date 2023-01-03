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

func countSortOp(list []int) int {
	count := 0
	LEN := len(list) - 1
	for i := 1; i <= LEN-1; i++ {
		for j := LEN - 1; j >= i; j-- {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				count++
			}
		}
	}
	return count
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())

	ps := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		ps[i] = scanLineInt(sc, n, 1)
	}

	horiList := make([]int, n+1)
	vertList := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if ps[i][j] != 0 {
				horiList[i] = ps[i][j]
				vertList[j] = ps[i][j]
			}
		}
	}
	vertCount := countSortOp(vertList)
	horiCount := countSortOp(horiList)
	fmt.Println(vertCount + horiCount)
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
