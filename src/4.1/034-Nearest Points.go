package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() float64 {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return float64(n)
}

func main() {
	sc.Split(bufio.ScanWords)
	n := int(scan())
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = scan()
		y[i] = scan()
	}

	ans := math.MaxFloat64
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := x[i] - x[j]
			dy := y[i] - y[j]
			v := math.Sqrt(dx*dx + dy*dy)
			if ans > v {
				ans = v
			}
		}
	}

	fmt.Printf("%g", ans)
}
