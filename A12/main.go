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

	var min int = 10e9
	for _, a := range as {
		if a < min {
			min = a
		}
	}

	// 答えがxより大きいか判定
	check := func(x int) bool {
		sum := 0
		for i := 0; i < n; i++ {
			sum += x / as[i]
		}
		return sum < k
	}

	l := 0
	r := k * min
	if len(as) == 1 {
		fmt.Println(r)
		return
	}

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	fmt.Println(l)
}
