package main

import (
	"fmt"
)

func scanSeq(n int, start int) []int {
	items := make([]int, n)
	for i := start; i < n; i++ {
		fmt.Scan(&items[i])
	}
	return items
}

func main() {
	// input
	var n int
	fmt.Scanf("%d", &n)
	as := scanSeq(n+1, 2)
	bs := scanSeq(n+1, 3)

	sts := make([]int, n+1)
	sts[1] = 0
	sts[2] = as[2]
	for i := 3; i <= n; i++ {
		c1 := sts[i-1] + as[i]
		c2 := sts[i-2] + bs[i]
		if c1 < c2 {
			sts[i] = c1
		} else {
			sts[i] = c2
		}
	}

	routes := make([]int, n+1)
	routes[1] = n
	count := 2
	cur := n
	for {
		if sts[cur] == sts[cur-1]+as[cur] {
			cur = cur - 1
		} else if sts[cur] == sts[cur-2]+bs[cur] {
			cur = cur - 2
		}
		routes[count] = cur
		if cur == 1 {
			break
		}

		count++
	}
	fmt.Println(count)
	for i := count; i > 0; i-- {
		fmt.Printf("%d ", routes[i])
	}
	fmt.Println()
}
