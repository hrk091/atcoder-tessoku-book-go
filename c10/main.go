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
	w := atoi(sc.Text())

	if w == 1 {
		fmt.Println(12)
		return
	}

	cache := map[int]int{}
	cache[1] = 7
	max := 1
	for i := 1; i <= w/2; i = i * 2 {
		cache[i*2] = cache[i] * cache[i] % 1000000007
		max = i * 2
	}
	if debug > 0 {
		fmt.Printf("cache: %+v\n", cache)
		fmt.Printf("max: %+v\n", max)
	}

	sum := 1
	rest := w - 1
	for i := max; i > 0; i = i / 2 {
		if i <= rest {
			rest = rest - i
			sum = (sum * cache[i]) % 1000000007
		}
	}
	fmt.Println(sum * 12 % 1000000007)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	mustNil(err)
	return i
}

func mustNil(err error) {
	if err != nil {
		panic(err)
	}
}
