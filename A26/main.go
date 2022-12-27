package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
}

func main() {
	// input
	var q int
	fmt.Scanf("%d", &q)

	xs := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&xs[i])
	}

	// main
	for _, x := range xs {
		check(x)
	}
}

func check(x int) {
	if x == 2 {
		fmt.Println("Yes")
		return
	}
	max := int(math.Sqrt(float64(x))) + 1
	for i := 2; i <= max; i++ {
		if x%i == 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
