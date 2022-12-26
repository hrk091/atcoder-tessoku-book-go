package main

import (
	"fmt"
	"sort"
)

func scanSeq(n int) []int {
	items := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&items[i])
	}
	return items
}

func main() {
	// input
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	as := scanSeq(n)
	bs := scanSeq(n)
	cs := scanSeq(n)
	ds := scanSeq(n)
	ps := make([]int, n*n)
	qs := make([]int, n*n)

	// main
	count := 0
	for _, a := range as {
		for _, b := range bs {
			ps[count] = a + b
			count++
		}
	}
	count = 0
	for _, c := range cs {
		for _, d := range ds {
			qs[count] = c + d
			count++
		}
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i] < ps[j]
	})

	check := func(ps []int, q int) bool {
		posL, posR := 0, len(ps)
		for posL < posR {
			posMid := (posL + posR) / 2
			//fmt.Printf("l=%v, posR=%v, k-q=%v, value=%v\n", posL, posR, k-q, ps[posMid])
			if ps[posMid] == k-q {
				return true
			} else if ps[posMid] < k-q {
				posL = posMid + 1
			} else {
				posR = posMid
			}
		}
		return false
	}

	for _, q := range qs {
		if check(ps, q) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
