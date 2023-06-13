package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type matrix [2][2]int

const mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func multiplication(a matrix, b matrix) matrix {
	c := matrix{
		{0, 0},
		{0, 0},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				c[i][k] += a[i][j] * b[j][k]
				c[i][k] %= mod
			}
		}
	}
	return c
}

func power(m matrix, n int) matrix {
	var answer matrix
	p := m
	flag := false
	for i := 0; i < 60; i++ {
		if (n & (1 << i)) != 0 {
			if !flag {
				answer = p
				flag = true
			} else {
				answer = multiplication(answer, p)
			}
		}
		p = multiplication(p, p)
	}
	return answer
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()

	a := matrix{
		{2, 1},
		{1, 0},
	}
	ans := power(a, n-1)
	fmt.Println((ans[1][0] + ans[1][1]) % mod)
}
