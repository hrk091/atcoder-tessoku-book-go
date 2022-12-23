package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// input
	var n []byte
	fmt.Scanf("%s", &n)

	// main logic
	var v float64
	for i := 0; i < len(n); i++ {
		r := n[len(n)-1-i]
		b, _ := strconv.Atoi(string(r))
		v += float64(b) * math.Pow(2, float64(i))
	}
	fmt.Println(v)
}
