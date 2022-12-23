package main

import (
	"fmt"
)

func main() {
	var n, x int

	fmt.Scanf("%d %d", &n, &x)

	for c := 0; c < n; c++ {
		var a int
		fmt.Scan(&a)
		if a == x {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
