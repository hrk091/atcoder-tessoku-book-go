package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())

	type span struct {
		l int
		r int
	}
	ss := make([]span, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		ss[i].l = atoi(sc.Text())
		sc.Scan()
		ss[i].r = atoi(sc.Text())
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].r != ss[j].r {
			return ss[i].r < ss[j].r
		}
		return ss[i].l > ss[j].l
	})
	var maxL int
	for _, s := range ss {
		if maxL < s.l {
			maxL = s.l
		}
	}

	var count, cur int
	for i := 0; i < len(ss); i++ {
		if cur <= ss[i].l {
			count++
			cur = ss[i].r
		}
	}
	fmt.Println(count)
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
