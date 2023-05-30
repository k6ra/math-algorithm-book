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
	m := 1000000007

	fact := make([]int, 200001)
	fact[0] = 1
	for i := 1; i <= 200000; i++ {
		fact[i] = i * fact[i-1] % m
	}

	x := scan()
	y := scan()

	p := fact[y] * fact[x] % m
	ans := 1
	for i := 0; i < 30; i++ {
		if ((m - 2) & (1 << i)) != 0 {
			ans *= p
			ans %= m
		}
		p *= p
		p %= m
	}

	fmt.Println(fact[x+y] * ans % m)
}
