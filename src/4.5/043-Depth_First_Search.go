package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

var ab [][]int
var nmap []bool

func dfs(current int) {
	nmap[current] = true
	connect := ab[current]
	for _, v := range connect {
		if !nmap[v] {
			dfs(v)
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()
	m := scan()
	ab = make([][]int, n+1, n+1)
	for i := 0; i < m; i++ {
		a := scan()
		b := scan()
		ab[a] = append(ab[a], b)
		ab[b] = append(ab[b], a)
	}

	nmap = make([]bool, n+1, n+1)
	dfs(1)
	for i := 1; i <= n; i++ {
		if !nmap[i] {
			fmt.Println("The graph is not connected.")
			return
		}
	}

	fmt.Println("The graph is connected.")
}
