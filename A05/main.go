package main

import (
	"fmt"
)

func main() {
	// input
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	// main logic
	var count int
	for a := 1; a <= n; a++ {
		if k-a > n*2 || k-a < 2 {
			continue
		}
		for b := 1; b <= n; b++ {
			if k-a-b > n || k-a-b < 1 {
				continue
			}
			count++
		}
	}
	fmt.Println(count)
}
