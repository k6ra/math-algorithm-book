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
	c := scan()
	d := scan()

	max := a * c
	if max < a*d {
		max = a * d
	}
	if max < b*c {
		max = b * c
	}
	if max < b*d {
		max = b * d
	}

	fmt.Println(max)
}
