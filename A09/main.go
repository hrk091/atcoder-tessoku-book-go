package main

import (
	"fmt"
)

func main() {
	// input
	var h, w, n int
	fmt.Scanf("%d %d %d", &h, &w, &n)

	var as, bs, cs, ds [100009]int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d %d", &as[i], &bs[i], &cs[i], &ds[i])
	}

	var x [1509][1509]int
	for i := 0; i < n; i++ {
		a, b, c, d := as[i]-1, bs[i]-1, cs[i]-1, ds[i]-1
		x[a][b] += 1
		x[c+1][d+1] += 1
		x[a][d+1] -= 1
		x[c+1][b] -= 1
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

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Printf("%d ", x[i][j])
		}
		fmt.Println()
	}

}
