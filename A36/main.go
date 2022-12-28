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
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	// main
	if k < 2*(n-1) {
		fmt.Println("No")
		return
	}

	if k%2 == 1 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
