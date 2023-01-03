package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
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

	// main
	var ans []int
	m := int(math.Sqrt(float64(n)))
	for i := 1; i <= m; i++ {
		if n%i == 0 {
			ans = append(ans, i)
			if n/i != i {
				ans = append(ans, n/i)
			}
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	for _, a := range ans {
		fmt.Println(a)
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	mustNil(err)
	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func mustNil(err error) {
	if err != nil {
		panic(err)
	}
}
