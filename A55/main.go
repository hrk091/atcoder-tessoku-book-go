package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
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

type empty struct{}

type set struct {
	m map[int]empty
}

func newset() *set {
	return &set{
		m: map[int]empty{},
	}
}

func (s *set) add(v int) {
	s.m[v] = empty{}
}

func (s *set) del(v int) {
	delete(s.m, v)
}

func (s *set) list(x int) []int {
	var l []int
	for v, _ := range s.m {
		if v >= x {
			l = append(l, v)
		}
	}
	return l
}

func (s *set) min(x int) int {
	l := s.list(x)
	if len(l) == 0 {
		return -1
	}
	return min(l...)
}

func main() {
	// input
	sc.Scan()
	q := atoi(sc.Text())

	// main
	s := newset()
	ans := make([]int, 100000)
	count := 0

	t := time.Now()
	for i := 1; i <= q; i++ {
		sc.Scan()
		act := sc.Text()
		switch act {
		case "1":
			sc.Scan()
			v := atoi(sc.Text())
			s.add(v)
		case "2":
			sc.Scan()
			v := atoi(sc.Text())
			s.del(v)
		case "3":
			sc.Scan()
			v := atoi(sc.Text())
			ans[count] = s.min(v)
			count++
		}
	}
	fmt.Println(time.Now().Sub(t))
	var res string
	for i := 0; i < count; i++ {
		res += fmt.Sprintf("%d\n", ans[i])
	}
	fmt.Print(res)
	fmt.Println(time.Now().Sub(t))
}

func abs(v int) int {
	return int(math.Abs(float64(v)))
}

func min(values ...int) int {
	mn := math.MaxInt64
	for _, v := range values {
		if v < mn {
			mn = v
		}
	}
	return mn
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
