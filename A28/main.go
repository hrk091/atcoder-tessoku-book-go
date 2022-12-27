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
	var n int
	fmt.Scanf("%d", &n)
	ts := make([]byte, n+1)
	as := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var s string
		fmt.Scan(&s)
		ts[i] = s[0]
		fmt.Scan(&as[i])
	}
	v := 0
	for i := 1; i <= n; i++ {
		switch string(ts[i]) {
		case "+":
			v += as[i]
		case "-":
			v = v - as[i]
			if v < 0 {
				v += ((v/10000)*-1 + 1) * 10000
			}
		case "*":
			v = v * as[i]
		}
		v = v % 10000
		fmt.Println(v)
	}
}
