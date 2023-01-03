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

func bitToList(b int, len int) ([]int, int) {
	var ret []int
	count := 0
	for i := 0; i < len; i++ {
		v := b >> i & 1
		if v == 1 {
			count++
		}
		ret = append(ret, v)
	}
	return ret, count
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())

	// main
	l, _ := bitToList(n-1, 10)
	ans := 0
	for i, v := range l {
		if v == 1 {
			ans += int(math.Pow(10, float64(i))) * 7
		} else {
			ans += int(math.Pow(10, float64(i))) * 4
		}
	}
	fmt.Println(ans)
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
