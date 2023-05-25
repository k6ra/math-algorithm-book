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

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()
	q := scan()

	area := make([]int, n+1)
	for i := 0; i < q; i++ {
		l := scan()
		r := scan()
		x := scan()

		area[l] += x
		if r < n {
			area[r+1] -= x
		}
	}

	for i := 2; i <= n; i++ {
		if area[i] == 0 {
			fmt.Print("=")
		} else if area[i] > 0 {
			fmt.Print("<")
		} else {
			fmt.Print(">")
		}
	}
}
