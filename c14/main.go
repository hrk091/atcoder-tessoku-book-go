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

type edge struct {
	farend int
	len    int
}

type dp struct {
	total int
	fixed int
}

func main() {
	// input
	sc.Scan()
	n := atoi(sc.Text())
	sc.Scan()
	m := atoi(sc.Text())

	g := newGraph(n)
	for i := 1; i <= m; i++ {
		sc.Scan()
		a := atoi(sc.Text())
		sc.Scan()
		b := atoi(sc.Text())
		sc.Scan()
		c := atoi(sc.Text())

		g.addEdge(a, b, c)
		g.addEdge(b, a, c)
	}

	// main
	dps := make([]dp, n+1)

	var ok bool

	// output
	if debug == 0 {
		if ok {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func abs(v int) int {
	return int(math.Abs(float64(v)))
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

func min(values ...int) int {
	mn := math.MaxInt64
	for _, v := range values {
		if v < mn {
			mn = v
		}
	}
	return mn
}

func sum(values ...int) int {
	var a int
	for _, v := range values {
		a += v
	}
	return a
}

func bitToList(b int, len int) ([]int, int) {
	// 63, 8 => [1,1,1,1,1,1,0,0]
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

func binarySearch(l, r int, fn func(int) int) int {
	// fn must be the one that returns true only when the result is greater than the given value.
	for l < r {
		mid := (l + r) / 2
		if ret := fn(mid); ret > 0 {
			l = mid + 1
		} else if ret < 0 {
			r = mid
		} else {
			return mid
		}
	}
	return l
}

type graph struct {
	nodeSize int
	data     map[int][]edge
}

func newGraph(nodeSize int) *graph {
	data := make(map[int][]edge, nodeSize+1)
	return &graph{
		nodeSize: nodeSize,
		data:     data,
	}
}

func (g *graph) addEdge(a, b, c int) {
	g.data[a] = append(g.data[a], edge{farend: b, len: c})
}

func (g *graph) newVisited() []bool {
	visited := make([]bool, g.nodeSize+1)
	visited[0] = true
	return visited
}

func (g *graph) isCompleted(visited []bool) bool {
	completed := true
	for _, v := range visited {
		if !v {
			completed = false
			break
		}
	}
	if debug > 0 {
		var visitedP []int
		for i, v := range visited {
			if i != 0 && v {
				visitedP = append(visitedP, i)
			}
		}
		fmt.Printf("visited: %+v\n", visitedP)
		fmt.Printf("completed: %+v\n", completed)
	}
	return completed
}

type queue struct {
	data []int
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) push(v int) {
	q.data = append(q.data, v)
}

func (q *queue) pop() int {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *queue) empty() bool {
	return len(q.data) == 0
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
