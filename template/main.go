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
	var a int
	fmt.Scanf("%d", &a)

	// main
	var ok bool

	// output
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
