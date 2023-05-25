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
	a := make([]int, n+1)
	diff := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = scan()
		diff[i] = diff[i-1] + a[i]
	}

	for i := 0; i < q; i++ {
		l := scan()
		r := scan()

		fmt.Printf("%d\n", diff[r]-diff[l-1])
	}
}
