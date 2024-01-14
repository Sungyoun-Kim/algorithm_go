package main

import (
	"bufio"
	"fmt"
	"os"
)

var tree [][]int
var visited []bool
var parents []int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)

	var tempTree = make([][]int, N+1)
	for i := 0; i < int(N)-1; i++ {
		var u, v int
		fmt.Fscanln(reader, &u, &v)
		tempTree[u] = append(tempTree[u], v)
		tempTree[v] = append(tempTree[v], u)
	}
	tree = tempTree
	visited = make([]bool, N+1)
	parents = make([]int, N+1)

	visited[1] = true
	findParent(1, 0)
	for _, v := range parents {
		if v == 0 {
			continue
		}

		fmt.Fprintln(writer, v)
	}
}

func findParent(start int, parent int) {
	parents[start] = parent
	for _, v := range tree[start] {
		if !visited[v] {
			visited[v] = true
			findParent(v, start)
		}
	}

}
