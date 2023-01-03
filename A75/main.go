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

type exam struct {
	required int
	end      int
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())

	exams := make([]exam, n+1)
	for i := 1; i <= n; i++ {
		sc.Scan()
		exams[i].required = atoi(sc.Text())
		sc.Scan()
		exams[i].end = atoi(sc.Text())
	}

	sort.Slice(exams, func(i, j int) bool {
		return exams[i].end < exams[j].end
	})
	if debug > 0 {
		fmt.Printf("exams: %+v\n", exams)
	}

	dpmatrix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dpmatrix[i] = make([]int, 1449)
	}

	for i := 1; i <= n; i++ {
		copy(dpmatrix[i], dpmatrix[i-1])
		ex := exams[i]
		next := ex.required
		if ex.required <= ex.end {
			dpmatrix[i][next] = max(1, dpmatrix[i][next])
		}
		for j := 1; j <= 1440; j++ {
			if dpmatrix[i-1][j] == 0 {
				continue
			}
			next := j + ex.required
			if next <= ex.end {
				dpmatrix[i][next] = max(dpmatrix[i-1][j]+1, dpmatrix[i][next])
			}
		}

		if debug > 0 {
			fmt.Printf("--- i=%d\n", i)
			for i := 0; i <= n; i++ {
				fmt.Printf("%+v\n", dpmatrix[i])
			}
		}
	}

	fmt.Println(max(dpmatrix[n]...))
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
