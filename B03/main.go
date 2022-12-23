package main

import (
	"fmt"
)

func scanSeq(n int) []int {
	var items []int
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		items = append(items, a)
	}
	return items
}

func main() {
	// input
	var n int
	fmt.Scanf("%d", &n)

	as := scanSeq(n)

	var ok bool
	for _, a := range as {
		for _, b := range as {
			if b == a {
				continue
			}
			for _, c := range as {
				if c == a || c == b {
					continue
				}
				if a+b+c == 1000 {
					ok = true
				}
			}
		}
	}

	// output
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
