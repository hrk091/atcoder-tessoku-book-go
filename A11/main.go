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
	var n, x int
	fmt.Scanf("%d %d", &n, &x)
	as := scanSeq(n)

	pos := 0
	for {
		c := len(as) / 2
		//fmt.Printf("%v, tgt=%d:%d, pos=%d\n", as, c, as[c], pos)
		if as[c] < x {
			pos += len(as[:c+1])
			as = as[c+1:]
		} else if as[c] > x {
			as = as[:c]
		} else {
			pos += len(as[:c])
			break
		}
	}
	fmt.Println(pos + 1)
}
