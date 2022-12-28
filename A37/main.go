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
	var n, m, b int
	fmt.Scanf("%d %d %d", &n, &m, &b)
	as := scanLineInt(sc, n, 0)
	cs := scanLineInt(sc, m, 0)

	fmt.Println(b*n*m + sum(as...)*m + sum(cs...)*n)
}

func sum(values ...int) int {
	var a int
	for _, v := range values {
		a += v
	}
	return a
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
