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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return a * -1
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()
	k := scan()

	yojisho := 0
	for a := 1; a <= n; a++ {
		for b := max(1, a-(k-1)); b <= min(n, a+k-1); b++ {
			for c := max(1, a-(k-1)); c <= min(n, a+k-1); c++ {
				if abs(b-c) <= k-1 {
					yojisho++
				}
			}
		}
	}

	fmt.Println(n*n*n - yojisho)
}
