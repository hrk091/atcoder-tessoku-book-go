package main

import (
	"fmt"
)

func main() {
	// input
	var h, w int
	fmt.Scanf("%d %d", &h, &w)
	var x [1509][1509]int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Scan(&x[i][j])
		}
	}

	var q int
	fmt.Scan(&q)
	var as, bs, cs, ds [100009]int
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d %d %d", &as[i], &bs[i], &cs[i], &ds[i])
	}

	for i := 0; i < h; i++ {
		for j := 1; j < w; j++ {
			x[i][j] = x[i][j-1] + x[i][j]
		}
	}
	for j := 0; j < w; j++ {
		for i := 1; i < h; i++ {
			x[i][j] = x[i-1][j] + x[i][j]
		}
	}

	for i := 0; i < q; i++ {
		a, b, c, d := as[i]-1, bs[i]-1, cs[i]-1, ds[i]-1

		var lt, lb, rt, rb int
		if a > 0 && b > 0 {
			lt = x[a-1][b-1]
		}
		if a > 0 {
			rt = x[a-1][d]
		}
		if b > 0 {
			lb = x[c][b-1]
		}
		rb = x[c][d]

		ans := rb + lt - lb - rt
		fmt.Println(ans)
	}
}
