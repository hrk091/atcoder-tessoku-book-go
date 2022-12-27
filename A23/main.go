package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
}

func toBits(as []int) int {
	var b int
	for i, a := range as {
		if a == 1 {
			b += 1 << i
		}
	}
	return b
}

func main() {
	// input
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	as := make([][]int, m+1)
	for i := 1; i <= m; i++ {
		as[i] = scanLineInt(sc, n, 1)
	}

	// main
	LEN := 1 << n
	matrix := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		matrix[i] = make([]int, LEN)
	}
	fillMatrix(matrix, -1)
	for i := 0; i <= m; i++ {
		matrix[i][0] = 0
	}

	for i := 1; i <= m; i++ {
		b := toBits(as[i][1:])
		for j := 0; j < LEN; j++ {
			matrix[i][j] = matrix[i-1][j]
		}
		for j := 0; j < LEN; j++ {
			if matrix[i-1][j] != -1 {
				matrix[i][j|b] = min(matrix[i][j|b], matrix[i-1][j]+1)
			}
		}

		//fmt.Println("---")
		//for j := 0; j <= m; j++ {
		//	for k := 0; k < LEN; k++ {
		//		v := matrix[j][k]
		//		if v == -1 {
		//			fmt.Print("* ")
		//		} else {
		//			fmt.Printf("%d ", v)
		//		}
		//	}
		//	fmt.Println()
		//}
	}

	ans := matrix[m][LEN-1]
	if ans == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}

}

func min(values ...int) int {
	mn := math.MaxInt64
	for _, v := range values {
		if v == -1 {
			continue
		}
		if v < mn {
			mn = v
		}
	}
	return mn
}

func scanLineInt(sc *bufio.Scanner, size, offset int) []int {
	items := make([]int, size+offset)
	sc.Scan()
	for i, s := range strings.Split(sc.Text(), " ") {
		items[i+offset] = atoi(s)
	}
	return items
}

func fillSlice(s []int, v int) {
	s[0] = v
	for p := 1; p < len(s); p *= 2 {
		copy(s[p:], s[:p])
	}
}

func fillMatrix(s [][]int, v int) {
	fillSlice(s[0], v)
	for p := 1; p < len(s); p++ {
		copy(s[p], s[0])
	}
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
