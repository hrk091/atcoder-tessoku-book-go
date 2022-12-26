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
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	as := scanSeq(n)

	l := 0
	r := 1
	rs := make([]int, n)
	for l < n-1 {
		diff := as[r] - as[l]
		if diff <= k {
			if r >= n-1 {
				rs[l] = r
				l++
			} else {
				r++
			}
		} else {
			rs[l] = r - 1
			l++
		}
	}

	ans := 0
	for i := 0; i < len(rs)-1; i++ {
		ans += rs[i] - i
	}

	fmt.Println(ans)
}
