package main

import (
	"bufio"
	"fmt"
	"math"
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
	x1 := float64(scan())
	y1 := float64(scan())
	r1 := float64(scan())

	x2 := float64(scan())
	y2 := float64(scan())
	r2 := float64(scan())

	dx := x1 - x2
	dy := y1 - y2
	v := math.Sqrt(dx*dx + dy*dy)
	if v < r1 || v < r2 {
		if v+r1 < r2 || v+r2 < r1 {
			fmt.Printf("%d", 1)
			return
		}

		if v+r1 == r2 || v+r2 == r1 {
			fmt.Printf("%d", 2)
			return
		}
	}

	if v < r1 || v < r2 || v < r1+r2 {
		fmt.Printf("%d", 3)
		return
	}

	if v == r1+r2 {
		fmt.Printf("%d", 4)
		return
	}

	fmt.Printf("%d", 5)
}
