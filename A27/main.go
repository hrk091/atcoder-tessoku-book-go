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
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	var ans int
	for {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
		if a == 0 {
			ans = b
			break
		}
		if b == 0 {
			ans = a
			break
		}
	}
	fmt.Println(ans)
}
