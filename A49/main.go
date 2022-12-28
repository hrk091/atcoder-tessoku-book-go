package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	debug int
	k     = 1000
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

func (s *status) eval() int {
	var sum int
	for _, v := range s.array {
		sum += abs(v)
	}
	return sum
}

func (s *status) eval2() int {
	return s.totalScore
}

func (s *status) show() {
	fmt.Printf("score: %d, array: %v\n", s.totalScore, s.array)
}

func selectTop(n int, sts []*status) []*status {
	sort.Slice(sts, func(i, j int) bool {
		return sts[i].eval() < sts[j].eval()
	})
	if debug > 1 {
		fmt.Println("---")
		for _, st := range sts {
			st.show()
		}
	}
	if len(sts) < n {
		return sts
	}
	return sts[:n]
}

func selectTop2(n int, sts []*status) []*status {
	sort.Slice(sts, func(i, j int) bool {
		return sts[i].eval2() > sts[j].eval2()
	})
	if debug > 1 {
		fmt.Println("---")
		for _, st := range sts {
			st.show()
		}
	}
	if len(sts) < n {
		return sts
	}
	return sts[:n]
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
	sts := []*status{newStatus()}
	for i := 1; i <= t-10; i++ {
		p, q, r := ps[i], qs[i], rs[i]
		var cands []*status
		for _, st := range sts {
			cands = append(cands, st.apply(true, p, q, r), st.apply(false, p, q, r))
		}
		sts = selectTop(k, cands)
	}
	for i := t - 9; i <= t; i++ {
		p, q, r := ps[i], qs[i], rs[i]
		var cands []*status
		for _, st := range sts {
			cands = append(cands, st.apply(true, p, q, r), st.apply(false, p, q, r))
		}
		sts = selectTop2(k, cands)
	}

	// output
	st := selectTop2(1, sts)[0]
	if debug > 0 {
		fmt.Println("---")
		st.show()
	}
	if debug == 0 {
		for i := 0; i < t; i++ {
			if st.history[i] {
				fmt.Println("A")
			} else {
				fmt.Println("B")
			}
		}
	}
}

func abs(v int) int {
	return int(math.Abs(float64(v)))
}

func atoi(s string) int {
	if s == "" {
		return 0
	}
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
