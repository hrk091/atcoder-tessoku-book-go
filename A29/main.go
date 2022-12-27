package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	// main
	type vv struct {
		idx   int
		power int
		val   int
	}
	var res []vv
	count := 0
	res = append(res, vv{
		power: 1,
		val:   a,
	})
	for i := 2; i <= b; i *= 2 {
		count++
		res = append(res, vv{
			idx:   count,
			power: i,
			val:   res[count-1].val * res[count-1].val % 1000000007,
		})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].idx > res[j].idx
	})

	pSum, vSum := 0, 1
	for _, v := range res {
		if pSum+v.power <= b {
			pSum += v.power
			vSum = vSum * v.val % 1000000007
		}
		if pSum == b {
			break
		}
	}
	fmt.Println(vSum)
}
