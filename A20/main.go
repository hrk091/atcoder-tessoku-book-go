package main

import (
	"fmt"
)

func scanSeq(n int) []int {
	items := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&items[i])
	}
	return items
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

func main() {
	// input
	var s, t []byte
	fmt.Scan(&s)
	fmt.Scan(&t)

	// main
	dpmatrix := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dpmatrix[i] = make([]int, len(t)+1)
	}
	dpmatrix[0][0] = 0

	matched := func(posI, posJ int) bool {
		if posI == 0 || posJ == 0 {
			return false
		}
		ss := s[posI-1 : posI][0]
		tt := t[posJ-1 : posJ][0]
		return ss == tt
	}

	for posI := 0; posI <= len(s); posI++ {
		for posJ := 0; posJ <= len(t); posJ++ {
			if posI == 0 || posJ == 0 {
				dpmatrix[posI][posJ] = 0
				continue
			}
			candidates := []int{dpmatrix[posI][posJ-1], dpmatrix[posI-1][posJ]}
			if matched(posI, posJ) {
				candidates = append(candidates, dpmatrix[posI-1][posJ-1]+1)
			}
			m := max(candidates...)
			dpmatrix[posI][posJ] = m
			//fmt.Printf("%d, %d, %d\n", posI, posJ, m)
		}
	}
	fmt.Println(dpmatrix[len(s)][len(t)])

}
