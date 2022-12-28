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
	q := atoi(sc.Text())

	// main
	var queue []string
	for i := 1; i <= q; i++ {
		sc.Scan()
		act := sc.Text()
		switch act {
		case "1":
			sc.Scan()
			title := sc.Text()
			queue = append(queue, title)
		case "2":
			title := queue[0]
			fmt.Println(title)
		case "3":
			if len(queue) > 0 {
				queue = queue[1:]
			}
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
