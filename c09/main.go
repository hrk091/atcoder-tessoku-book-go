package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	debug int
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	flag.Parse()
	d := flag.Arg(0)
	if d != "" {
		debug = atoi(d)
	}
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())
	as := scanLineInt(sc, n, 1)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}

	for i := 1; i <= n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + as[i]
		if debug > 0 {
			fmt.Printf("dp: %+v\n", dp)
		}
	}

	fmt.Println(max(dp[n]...))
}

func max(values ...int) int {
	var mx int
	for _, v := range values {
		if v > mx {
			mx = v
		}
	}
	return mx
}

func scanLineInt(sc *bufio.Scanner, size, offset int) []int {
	items := make([]int, size+offset)
	for i := 0; i < size; i++ {
		sc.Scan()
		items[i+offset] = atoi(sc.Text())
	}
	return items
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	mustNil(err)
	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func btoi(b byte) int {
	if b < '0' || '9' < b {
		panic(fmt.Errorf("cannot convert %s to int", []byte{b}))
	}
	return atoi(string(b))
}

func itob(i int) byte {
	if i < 0 || i > 9 {
		panic(fmt.Errorf("cannot convert %d to byte", i))
	}
	return byte(i + '0')
}

func mustNil(err error) {
	if err != nil {
		panic(err)
	}
}
