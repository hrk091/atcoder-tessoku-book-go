package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

type data struct {
	size     int
	reversed bool
	array    []int
}

func newData(s int) *data {
	ar := make([]int, s+1)
	for i := 1; i <= s; i++ {
		ar[i] = i
	}
	return &data{
		size:     s,
		reversed: false,
		array:    ar,
	}
}

func (d *data) reverse() {
	d.reversed = !d.reversed
}
func (d *data) set(x, y int) {
	if !d.reversed {
		d.array[x] = y
	} else {
		d.array[d.size-x+1] = y
	}
}

func (d *data) show(x int) {
	if !d.reversed {
		fmt.Println(d.array[x])
	} else {
		fmt.Println(d.array[d.size-x+1])
	}
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())
	sc.Scan()
	q := atoi(sc.Text())

	d := newData(n)
	for i := 0; i < q; i++ {
		sc.Scan()
		act := sc.Text()
		switch act {
		case "1":
			sc.Scan()
			x := atoi(sc.Text())
			sc.Scan()
			y := atoi(sc.Text())
			d.set(x, y)
		case "2":
			d.reverse()
		case "3":
			sc.Scan()
			x := atoi(sc.Text())
			d.show(x)
		}
	}
}

func max(values ...int) int {
	var mx int
	for _, v := range values {
		if v > mx {
			mx = v
		}
	}
	return mx
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

func sum(values ...int) int {
	var a int
	for _, v := range values {
		a += v
	}
	return a
}

func bs(l, r int, fn func(int) int) int {
	// fn must be the one that returns true only when the result is greater than the given value.
	for l < r {
		mid := (l + r) / 2
		if ret := fn(mid); ret > 0 {
			l = mid + 1
		} else if ret < 0 {
			r = mid
		} else {
			return mid
		}
	}
	return l
}

func scanLineInt(sc *bufio.Scanner, size, offset int) []int {
	items := make([]int, size+offset)
	for i := 0; i < size; i++ {
		sc.Scan()
		items[i+offset] = atoi(sc.Text())
	}
	return items
}

func fillSlice(s []int, v int) {
	s[0] = v
	for p := 1; p < len(s); p *= 2 {
		copy(s[p:], s[:p])
	}
}

func fillMatrix(s [][]int, v int) {
	fillSlice(s[0], v)
	for p := 1; p < len(s); p++ {
		copy(s[p], s[0])
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
