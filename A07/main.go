package main

import (
	"fmt"
)

func main() {
	// input
	var d, n int
	fmt.Scan(&d)
	fmt.Scan(&n)
	ls := make([]int, n)
	rs := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&ls[i])
		fmt.Scan(&rs[i])
	}

	// main
	dd := d + 10
	diffs := make([]int, dd)
	for _, l := range ls {
		diffs[l-1] += 1
	}
	for _, r := range rs {
		diffs[(r-1)+1] -= 1
	}

	answers := make([]int, dd)
	answers[0] = diffs[0]
	for i := 1; i < d; i++ {
		answers[i] = answers[i-1] + diffs[i]
	}

	for i := 0; i < d; i++ {
		fmt.Println(answers[i])
	}
}
