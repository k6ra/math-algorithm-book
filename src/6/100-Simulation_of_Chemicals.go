package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanFloat() float64 {
	sc.Scan()
	n, _ := strconv.ParseFloat(sc.Text(), 64)
	return n
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func multiplication(a, b [3][3]float64) [3][3]float64 {
	var c [3][3]float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}

func power(a [3][3]float64, n int) [3][3]float64 {
	var q [3][3]float64
	p := a
	flag := false
	for i := 0; i < 30; i++ {
		if (n & (1 << i)) != 0 {
			if !flag {
				q = p
				flag = true
			} else {
				q = multiplication(q, p)
			}
		}
		p = multiplication(p, p)
	}
	return q
}

func main() {
	sc.Split(bufio.ScanWords)

	q := scanInt()
	for i := 1; i <= q; i++ {
		x := scanFloat()
		y := scanFloat()
		z := scanFloat()
		t := scanInt()

		var a [3][3]float64
		a[0][0] = 1.0 - x
		a[2][0] = x
		a[1][1] = 1.0 - y
		a[0][1] = y
		a[2][2] = 1.0 - z
		a[1][2] = z

		b := power(a, t)

		ansA := b[0][0] + b[0][1] + b[0][2]
		ansB := b[1][0] + b[1][1] + b[1][2]
		ansC := b[2][0] + b[2][1] + b[2][2]
		fmt.Printf("%g %g %g\n", ansA, ansB, ansC)
	}
}
