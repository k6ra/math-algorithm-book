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

const mod = 1000000007

var fact [200001]int

func modpow(a, b, m int) int {
	p := a
	ans := 1
	for i := 0; i < 30; i++ {
		if (b & (1 << i)) != 0 {
			ans *= p
			ans %= m
		}
		p *= p
		p %= m
	}
	return ans
}

func division(a, b, m int) int {
	return (a * modpow(b, m-2, m)) % m
}

func ncr(n, r int) int {
	return division(fact[n], fact[r]*fact[n-r]%mod, mod)
}

func main() {
	sc.Split(bufio.ScanWords)

	fact[0] = 1
	for i := 1; i <= 200000; i++ {
		fact[i] = fact[i-1] * i % mod
	}

	n := scan()
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = scan()
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += a[i] * ncr(n-1, i-1)
		ans %= mod
	}

	fmt.Println(ans)
}
