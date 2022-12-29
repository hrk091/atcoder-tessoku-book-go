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

type resolver struct {
	len       int
	next      []int
	cache     [][]int
	totalStep int
}

func newResolver(l int, next []int) *resolver {
	return &resolver{
		len:       l,
		next:      next,
		cache:     make([][]int, 100),
		totalStep: 0,
	}
}

func (r *resolver) addStep() {
	r.totalStep++
	r.cache[r.totalStep] = make([]int, r.len+1)
	if r.totalStep == 1 {
		copy(r.cache[r.totalStep], r.next)
	} else {
		for i := 1; i <= r.len; i++ {
			pos := r.cache[r.totalStep-1][i]
			pos = r.cache[r.totalStep-1][pos]
			r.cache[r.totalStep][i] = pos
		}
	}
}

func (r *resolver) getAfter(step int) []int {
	if res := r.cache[step]; res == nil {
		r.addStep()
	} else {
		return res
	}
	if res := r.cache[step]; res == nil {
		panic(fmt.Sprintf("previous step cache is not generated yet: %+v\n", r))
	} else {
		return res
	}
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())
	sc.Scan()
	q := atoi(sc.Text())

	as := scanLineInt(sc, n, 1)
	rslv := newResolver(n, as)

	for i := 1; i <= q; i++ {
		sc.Scan()
		x := atoi(sc.Text())
		sc.Scan()
		y := atoi(sc.Text())

		step := 1
		pos := x
		for y > 0 {
			cache := rslv.getAfter(step)
			if y%2 == 1 {
				pos = cache[pos]
			}
			y = y / 2
			step++
		}
		fmt.Println(pos)
	}
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
