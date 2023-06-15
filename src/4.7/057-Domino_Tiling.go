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

var matrix2 = [4][4]int{
	{0, 0, 0, 1},
	{0, 0, 1, 0},
	{0, 1, 0, 0},
	{1, 0, 0, 1},
}

var matrix3 = [8][8]int{
	{0, 0, 0, 0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 1, 0},
	{0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 1},
	{0, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 1, 0, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 1, 0, 0, 1, 0},
}

var matrix4 = [16][16]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
	{0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
	{1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1},
}

type matrix struct {
	size int
	p    [16][16]int
}

func multiplication(a matrix, b matrix) matrix {
	var c matrix

	c.size = a.size
	for i := 0; i < a.size; i++ {
		for j := 0; j < a.size; j++ {
			for k := 0; k < a.size; k++ {
				c.p[i][k] += a.p[i][j] * b.p[j][k]
				c.p[i][k] %= 1000000007
			}
		}
	}

	return c
}

func power(a matrix, n int) matrix {
	p := a
	var q matrix
	flag := false
	for i := 0; i < 60; i++ {
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

	k := scan()
	n := scan()

	var a matrix
	a.size = 1 << k
	for i := 0; i < (1 << k); i++ {
		for j := 0; j < (1 << k); j++ {
			if k == 2 {
				a.p[i][j] = matrix2[i][j]
			} else if k == 3 {
				a.p[i][j] = matrix3[i][j]
			} else {
				a.p[i][j] = matrix4[i][j]
			}
		}
	}

	b := power(a, n)

	fmt.Println(b.p[(1<<k)-1][(1<<k)-1])
}
