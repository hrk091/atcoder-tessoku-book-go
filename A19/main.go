package main

import (
	"fmt"
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
	var n, wmax int
	fmt.Scanf("%d %d", &n, &wmax)
	ws := make([]int, n+1)
	vs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d %d", &ws[i], &vs[i])
	}

	// main
	dpmatrix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dpmatrix[i] = make([]int, wmax+1)
	}
	for i := 1; i <= n; i++ {
		newW, newV := ws[i], vs[i]
		prev := dpmatrix[i-1]
		curr := dpmatrix[i]
		curr[newW] = newV
		for ww, vv := range prev {
			if vv == 0 {
				continue
			}
			if curr[ww] < vv {
				curr[ww] = vv
			}
		}
		for ww, vv := range prev {
			if ww+newW > wmax {
				continue
			}
			if vv == 0 {
				continue
			}
			if curr[ww+newW] < vv+newV {
				curr[ww+newW] = vv + newV
			}
		}
		// fmt.Printf("%v: %v\n", i, curr)
	}

	max := 0
	for _, v := range dpmatrix[len(dpmatrix)-1] {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
