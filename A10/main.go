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
	var n int
	fmt.Scanf("%d", &n)
	as := scanSeq(n)

	var d int
	fmt.Scanf("%d", &d)
	var ls, rs [100009]int
	for i := 0; i < d; i++ {
		fmt.Scanf("%d %d", &ls[i], &rs[i])
	}

	// main
	var bs, cs [100009]int
	bs[0] = as[0]
	for i := 1; i < n; i++ {
		if as[i] > bs[i-1] {
			bs[i] = as[i]
		} else {
			bs[i] = bs[i-1]
		}
	}
	cs[n-1] = as[n-1]
	for i := n - 2; i > -1; i-- {
		if as[i] > cs[i+1] {
			cs[i] = as[i]
		} else {
			cs[i] = cs[i+1]
		}
	}

	for i := 0; i < d; i++ {
		l, r := ls[i]-1, rs[i]-1
		var max int
		if bs[l-1] > cs[r+1] {
			max = bs[l-1]
		} else {
			max = cs[r+1]
		}
		fmt.Println(max)
	}

}
