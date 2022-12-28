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

type PQueue []int

func (q PQueue) show() {
	_, v := min(q...)
	fmt.Println(v)
}

func (q *PQueue) push(v int) {
	*q = append(*q, v)
}

func (q *PQueue) pop() {
	if len(*q) == 0 {
		return
	}
	pos, _ := min(*q...)
	*q = append((*q)[:pos], (*q)[pos+1:]...)
}

func main() {
	// input
	sc.Scan()
	q := atoi(sc.Text())

	// main
	var queue PQueue
	for i := 1; i <= q; i++ {
		sc.Scan()
		act := sc.Text()
		switch act {
		case "1":
			sc.Scan()
			val := sc.Text()
			queue.push(atoi(val))
		case "2":
			queue.show()
		case "3":
			queue.pop()
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

func min(values ...int) (int, int) {
	pos := 0
	mn := math.MaxInt64
	for i, v := range values {
		if v < mn {
			mn = v
			pos = i
		}
	}
	return pos, mn
}
