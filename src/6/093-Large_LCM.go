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

	a := scan()
	b := scan()

	c := a
	d := b
	for c > 0 && d > 0 {
		if c > d {
			c = c % d
		} else {
			d = d % c
		}
	}

	gcd := c
	if c < d {
		gcd = d
	}

	if a/gcd > 1000000000000000000/b {
		fmt.Println("Large")
		return
	}
	fmt.Println(a / gcd * b)
}
