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

func modpow(a int, b int, mod int) int {
	p := a
	ans := 1
	for i := 0; i < 60; i++ {
		if b&(1<<i) != 0 {
			ans *= p
			ans %= mod
		}
		p *= p
		p %= mod
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()

	const mod = 1000000007
	v := modpow(4, n+1, mod) - 1
	ans := v * modpow(3, mod-2, mod) % mod

	fmt.Println(ans)
}
