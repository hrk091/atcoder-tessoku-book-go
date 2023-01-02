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

type pathStatus struct {
	len  int
	tree int
}

type stat struct {
	minLen    int
	maxTree   int
	confirmed bool
}

func newStat() *stat {
	return &stat{
		minLen:    math.MaxInt64,
		maxTree:   0,
		confirmed: false,
	}
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())
	sc.Scan()
	m := atoi(sc.Text())

	as := make([]int, m+1)
	bs := make([]int, m+1)
	cs := make([]int, m+1)
	ds := make([]int, m+1)
	for i := 1; i <= m; i++ {
		sc.Scan()
		as[i] = atoi(sc.Text())
		sc.Scan()
		bs[i] = atoi(sc.Text())
		sc.Scan()
		cs[i] = atoi(sc.Text())
		sc.Scan()
		ds[i] = atoi(sc.Text())
	}

	graph := make([]map[int]pathStatus, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = map[int]pathStatus{}
	}
	for i := 1; i <= m; i++ {
		a, b, c, d := as[i], bs[i], cs[i], ds[i]
		graph[a][b] = pathStatus{
			len:  c,
			tree: d,
		}
	}
	if debug > 1 {
		fmt.Printf("graph: %+v\n", graph)
	}

	statAr := make([]*stat, n+1)
	for i := 0; i <= n; i++ {
		statAr[i] = newStat()
	}
	statAr[0].confirmed = true
	curr := 1
	for {
		if debug > 0 {
			fmt.Printf("curr: %d, len=%d, tree=%d\n", curr, statAr[curr].minLen, statAr[curr].maxTree)
		}
		if debug > 1 {
			for i := 1; i < n+1; i++ {
				fmt.Printf("statAr[%d]: %+v\n", i, statAr[i])
			}
		}
		if curr == n {
			break
		}

		statAr[curr].confirmed = true
		for i, v := range graph[curr] {
			if statAr[i].confirmed {
				continue
			}
			if debug > 1 {
				fmt.Printf("graph[%d]: %+v\n", i, v)
				for i := 1; i < n+1; i++ {
					fmt.Printf("statAr[%d]: %+v\n", i, statAr[i])
				}
			}
			currLen := 0
			if statAr[curr].minLen != math.MaxInt64 {
				currLen = statAr[curr].minLen
			}
			if currLen+v.len < statAr[i].minLen {
				statAr[i].minLen = currLen + v.len
				statAr[i].maxTree = statAr[curr].maxTree + v.tree
				continue
			}
			if currLen+v.len == statAr[i].minLen {
				if statAr[curr].maxTree+v.tree > statAr[i].maxTree {
					statAr[i].maxTree = statAr[curr].maxTree + v.tree
					continue
				}
			}
		}
		minLen := math.MaxInt64
		next := 0
		for i, st := range statAr {
			if st.confirmed {
				continue
			}
			if minLen > st.minLen {
				minLen = st.minLen
				next = i
			}
		}
		curr = next
	}
	fmt.Printf("%d %d\n", statAr[n].minLen, statAr[n].maxTree)
}

func fillSlice(s []int, v int) {
	s[0] = v
	for p := 1; p < len(s); p *= 2 {
		copy(s[p:], s[:p])
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
