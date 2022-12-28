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

type record struct {
	len       int
	direction string
}

func (r *record) rest(total int) int {
	if r.direction == "E" {
		return total - r.len
	} else {
		return r.len
	}
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())
	sc.Scan()
	l := atoi(sc.Text())

	records := make([]record, n+1)

	for i := 1; i <= n; i++ {
		sc.Scan()
		records[i].len = atoi(sc.Text())
		sc.Scan()
		records[i].direction = sc.Text()
	}
	var max int
	for i := 1; i <= n; i++ {
		ll := records[i].rest(l)
		if ll > max {
			max = ll
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
