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
	d := atoi(sc.Text())
	sc.Scan()
	_ = atoi(sc.Text())
	as := make([]int, d+1)
	for i := 2; i <= d; i++ {
		sc.Scan()
		as[i] = atoi(sc.Text())
	}
	sumAr := make([]int, d+1)
	sumAr[2] = as[2]
	for i := 3; i <= d; i++ {
		sumAr[i] = sumAr[i-1] + as[i]
	}

	sc.Scan()
	q := atoi(sc.Text())

	for i := 0; i < q; i++ {
		sc.Scan()
		s := atoi(sc.Text())
		sc.Scan()
		t := atoi(sc.Text())
		if sumAr[t] > sumAr[s] {
			fmt.Println(t)
		} else if sumAr[s] > sumAr[t] {
			fmt.Println(s)
		} else {
			fmt.Println("Same")
		}
	}
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
