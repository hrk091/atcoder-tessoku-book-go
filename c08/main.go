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

type ticket struct {
	num   int
	class int
}

func main() {
	failMsg := "Can't Solve"
	// input
	sc.Scan()
	n := atoi(sc.Text())
	tickets := make([]ticket, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		tickets[i].num = atoi(sc.Text())
		sc.Scan()
		tickets[i].class = atoi(sc.Text())
	}
	if debug > 0 {
		fmt.Printf("tickets: %+v\n", tickets)
	}

	sort.Slice(tickets, func(i, j int) bool {
		return tickets[i].class < tickets[j].class
	})

	cand := map[int]bool{}
	t := tickets[0]
	if t.class == 1 {
		fmt.Println(t.num)
		return
	}
	if t.class == 3 {
		fmt.Println(failMsg)
		return
	}
	if t.class == 2 {
		for j := 0; j <= 9; j++ {
			numStr := []byte(itoa(t.num))
			numStr[0] = byte(j + '0')
			cand[atoi(string(numStr))] = true
			numStr = []byte(itoa(t.num))
			numStr[1] = byte(j + '0')
			cand[atoi(string(numStr))] = true
			numStr = []byte(itoa(t.num))
			numStr[2] = byte(j + '0')
			cand[atoi(string(numStr))] = true
			numStr = []byte(itoa(t.num))
			numStr[3] = byte(j + '0')
			cand[atoi(string(numStr))] = true
		}
		if debug > 0 {
			fmt.Printf("=== %+v\n", t)
			for i, ok := range cand {
				if ok {
					fmt.Printf("%d ", i)
				}
			}
			fmt.Println()
		}
	}

	for i := 1; i < n; i++ {
		t := tickets[i]
		if t.class == 2 {
			for i, v := range cand {
				if !v {
					continue
				}
				count := matchCount(i, t.num)
				if count != 3 {
					cand[i] = false
				}
			}
		}
		if t.class == 3 {
			for i, v := range cand {
				if !v {
					continue
				}
				count := matchCount(i, t.num)
				if count > 2 {
					cand[i] = false
				}
			}
		}
		if debug > 0 {
			fmt.Printf("=== %+v\n", t)
			for i, ok := range cand {
				if ok {
					fmt.Printf("%d ", i)
				}
			}
			fmt.Println()
		}
	}
	var ans []int

	for i, ok := range cand {
		if ok {
			ans = append(ans, i)
		}
	}
	if len(ans) != 1 {
		fmt.Println(failMsg)
		return
	}
	fmt.Println(ans[0])
}

func matchCount(a, b int) int {
	count := 0
	if a/1000 == b/1000 {
		count++
	}
	if (a%1000)/100 == (b%1000)/100 {
		count++
	}
	if (a%100)/10 == (b%100)/10 {
		count++
	}
	if a%10 == b%10 {
		count++
	}
	return count
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
