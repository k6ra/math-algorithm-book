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
	a := make([]float64, n)
	b := make([]float64, n)
	c := make([]float64, n)
	for i := 0; i < n; i++ {
		a[i] = float64(scan())
		b[i] = float64(scan())
		c[i] = float64(scan())
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i]*b[j] == a[j]*b[i] {
				continue
			}
			x := (c[i]*b[j] - c[j]*b[i]) / (a[i]*b[j] - a[j]*b[i])
			y := (c[i]*a[j] - c[j]*a[i]) / (b[i]*a[j] - b[j]*a[i])

			cond := true
			for k := 0; k < n; k++ {
				if a[k]*x+b[k]*y > c[k] {
					cond = false
					break
				}
			}

			if cond && (x+y > ans) {
				ans = x + y
			}
		}
	}

	fmt.Println(ans)
}
