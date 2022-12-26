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

func max(values ...int) int {
	var mx int
	for _, v := range values {
		if v > mx {
			mx = v
		}
	}
	return mx
}

var (
	n        int
	ps       []int
	as       []int
	dpmatrix [][]int
)

func main() {
	// input
	fmt.Scanf("%d", &n)
	ps, as = make([]int, n+9), make([]int, n+9)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d %d", &ps[i], &as[i])
	}

	// main
	dpmatrix = make([][]int, n+9)
	for i := 0; i <= n; i++ {
		dpmatrix[i] = make([]int, n+9)
	}
	dpmatrix[1][n] = 0

	for LEN := n - 2; LEN >= 0; LEN-- {
		for l := 1; l <= n-LEN; l++ {
			r := l + LEN

			var lScore, rScore int
			if l <= ps[l-1] && ps[l-1] <= r {
				lScore = as[l-1]
			}
			if l <= ps[r+1] && ps[r+1] <= r {
				rScore = as[r+1]
			}
			if l == 1 {
				dpmatrix[l][r] = dpmatrix[l][r+1] + rScore
			} else if r == n {
				dpmatrix[l][r] = dpmatrix[l-1][r] + lScore
			} else {
				dpmatrix[l][r] = max(dpmatrix[l][r+1]+rScore, dpmatrix[l-1][r]+lScore)
			}
			//fmt.Printf("LEN=%d, l=%d r=%d %v\n", LEN, l, r, dpmatrix[l][r])
		}
	}

	var finalPoints []int
	for i := 1; i < n; i++ {
		finalPoints = append(finalPoints, dpmatrix[i][i])
	}
	fmt.Println(max(finalPoints...))
}
