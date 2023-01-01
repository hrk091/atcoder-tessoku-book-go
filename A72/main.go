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

func bitToList(b int, len int) ([]int, int) {
	var ret []int
	count := 0
	for i := 0; i < len; i++ {
		v := b >> i & 1
		if v == 1 {
			count++
		}
		ret = append(ret, v)
	}
	return ret, count
}

func bitGen(len int, upper int) [][]int {
	var ret [][]int
	for i := 0; i < int(math.Pow(2, float64(len))); i++ {
		l, count := bitToList(i, len)
		if count <= upper {
			ret = append(ret, l)
		}
	}
	return ret
}

func main() {
	// input
	sc.Scan()
	h := atoi(sc.Text())
	sc.Scan()
	w := atoi(sc.Text())
	sc.Scan()
	k := atoi(sc.Text())

	cm := make([][]int, h+1)
	cm[0] = make([]int, w+1)
	for i := 1; i <= h; i++ {
		cm[i] = make([]int, w+1)
		sc.Scan()
		for j, v := range sc.Bytes() {
			if string(v) == "#" {
				cm[i][j+1] = 1
			}
		}
	}

	getCopy := func() [][]int {
		newCm := make([][]int, h+1)
		for i := 1; i <= h; i++ {
			newCm[i] = make([]int, w+1)
			copy(newCm[i], cm[i])
		}
		return newCm
	}

	ans := 0
	for _, b := range bitGen(h, k) {
		cm2 := getCopy()
		if debug > 0 {
			fmt.Println("===")
			fmt.Println(b)
		}
		if debug > 1 {
			fmt.Printf("before: %v\n", cm2)
		}
		consumed := 0
		for i, tgt := range b {
			pos := i + 1
			if tgt == 1 {
				for p := 1; p <= w; p++ {
					cm2[pos][p] = 1
				}
				consumed++
			}
		}
		if debug > 1 {
			fmt.Printf("after: %v\n", cm2)
		}

		totalBlack := make([]int, w+1)
		for j := 1; j <= w; j++ {
			for i := 1; i <= h; i++ {
				if cm2[i][j] == 1 {
					totalBlack[j]++
				}
			}
		}
		sort.Slice(totalBlack, func(i, j int) bool {
			return totalBlack[i] < totalBlack[j]
		})

		cand := 0
		for j := 1; j <= w; j++ {
			if j <= k-consumed {
				cand += h
			} else {
				cand += totalBlack[j]
			}
		}
		if debug > 0 {
			fmt.Printf("totalBlack: %v, %v -> %v\n", totalBlack, sum(totalBlack...), cand)
		}
		ans = max(ans, cand)
	}
	fmt.Println(ans)
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

func sum(values ...int) int {
	var a int
	for _, v := range values {
		a += v
	}
	return a
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
