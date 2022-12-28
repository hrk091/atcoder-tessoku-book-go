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
	debug bool
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	flag.Parse()
	if flag.Arg(0) == "d" {
		debug = true
	}
}

type status struct {
	history    []bool
	totalScore int
	array      []int
}

func newStatus() *status {
	return &status{
		array: make([]int, 20),
	}
}

func (s *status) copy() *status {
	ns := status{
		history:    make([]bool, len(s.history)),
		totalScore: s.totalScore,
		array:      make([]int, 20),
	}
	copy(ns.history, s.history)
	copy(ns.array, s.array)
	return &ns
}

func (s *status) apply(add bool, p, q, r int) *status {
	cs := s.copy()
	if add {
		cs.array[p-1]++
		cs.array[q-1]++
		cs.array[r-1]++
	} else {
		cs.array[p-1]--
		cs.array[q-1]--
		cs.array[r-1]--
	}
	score := 0
	for _, v := range cs.array {
		if v == 0 {
			score++
		}
	}
	cs.totalScore += score
	cs.history = append(cs.history, add)
	return cs
}

func (s *status) eval1() int {
	return s.totalScore
}

func main() {
	// input
	sc.Scan()
	t := atoi(sc.Text())

	ps := make([]int, t+1)
	qs := make([]int, t+1)
	rs := make([]int, t+1)
	for i := 1; i < t+1; i++ {
		sc.Scan()
		ps[i] = atoi(sc.Text())
		sc.Scan()
		qs[i] = atoi(sc.Text())
		sc.Scan()
		rs[i] = atoi(sc.Text())
	}

	// main
	st := newStatus()
	for i := 1; i <= t; i++ {
		p, q, r := ps[i], qs[i], rs[i]
		c1 := st.apply(true, p, q, r)
		c2 := st.apply(false, p, q, r)
		if debug {
			fmt.Printf("%d: (%d, %d, %d) %d, %d\n", i, p, q, r, c1.eval1(), c2.eval1())
		}
		if c1.eval1() > c2.eval1() {
			st = c1
		} else {
			st = c2
		}
		if debug {
			fmt.Printf("st: %d, %+v\n", len(st.history), st.array)
		}
	}
	for i := 0; i < t; i++ {
		if st.history[i] {
			fmt.Println("A")
		} else {
			fmt.Println("B")
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
