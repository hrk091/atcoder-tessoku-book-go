package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	for i := a; i <= b; i++ {
		if 100%i == 0 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
