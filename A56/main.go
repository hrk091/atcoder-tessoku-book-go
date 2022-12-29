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
	_ = atoi(sc.Text())
	sc.Scan()
	q := atoi(sc.Text())
	sc.Scan()
	s := sc.Bytes()

	for i := 1; i <= q; i++ {
		sc.Scan()
		a := atoi(sc.Text())
		sc.Scan()
		b := atoi(sc.Text())
		sc.Scan()
		c := atoi(sc.Text())
		sc.Scan()
		d := atoi(sc.Text())
		fmt.Printf("%s %s\n", s[a-1:b], s[c-1:d])
		if string(s[a-1:b]) == string(s[c-1:d]) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
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
