package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())
	ss := make([]string, n+1)
	sc.Scan()
	for i, b := range sc.Bytes() {
		ss[i+1] = string(b)
	}

	// main
	for i := 1; i <= n-2; i++ {
		if ss[i] == ss[i+1] && ss[i+1] == ss[i+2] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	mustNil(err)
	return i
}

func mustNil(err error) {
	if err != nil {
		panic(err)
	}
}
