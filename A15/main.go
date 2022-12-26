package main

import (
	"fmt"
	"sort"
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
	bs := make([]int, len(as))
	for i, a := range as {
		bs[i] = a
	}

	sort.Slice(as, func(i, j int) bool {
		return as[i] < as[j]
	})

	count := 1
	m := map[int]int{}
	var now int
	for _, a := range as {
		if a != now {
			m[a] = count
			now = a
			count++
		}
	}

	for _, b := range bs {
		fmt.Printf("%d ", m[b])
	}
	fmt.Println()
}
