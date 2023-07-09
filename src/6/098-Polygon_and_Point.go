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
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = scan()
		y[i] = scan()
	}

	a := scan()
	b := scan()

	count := 0
	for i := 0; i < n; i++ {
		xa := x[i] - a
		ya := y[i] - b
		xb := x[(i+1)%n] - a
		yb := y[(i+1)%n] - b
		if ya > yb {
			xa, xb = xb, xa
			ya, yb = yb, ya
		}

		if ya <= 0 && 0 < yb && xa*yb-xb*ya < 0 {
			count++
		}
	}

	if count%2 == 1 {
		fmt.Println("INSIDE")
	} else {
		fmt.Println("OUTSIDE")
	}
}
