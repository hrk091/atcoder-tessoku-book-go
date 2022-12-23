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
	var n, q int
	fmt.Scanf("%d %d", &n, &q)
	as := scanSeq(n)
	ls := make([]int, q)
	rs := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d", &ls[i], &rs[i])
	}

	// main
	ts := make([]int, n+1, n+1)
	for i, a := range as {
		ts[i+1] = ts[i] + a
	}

	for i := 0; i < q; i++ {
		fmt.Println(ts[rs[i]] - ts[ls[i]-1])
	}
}
