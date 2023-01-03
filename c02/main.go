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

	as := scanLineInt(sc, n, 0)
	bs := make([]int, n)
	copy(bs, as)

	first := max(as...)

	for i, a := range bs {
		if a == first {
			bs[i] = 0
			break
		}
	}
	second := max(bs...)

	fmt.Println(first + second)
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
