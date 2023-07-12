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

var fact [100001]int

func modpow(a, b, m int) int {
	p := a
	ans := 1
	for i := 0; i < 30; i++ {
		if (b & (1 << i)) != 0 {
			ans *= p
			ans %= mod
		}
		p *= p
		p %= m
	}
	return ans
}

func division(a, b, m int) int {
	return a * modpow(b, m-2, m) % m
}

func ncr(n, r int) int {
	return division(fact[n], fact[r]*fact[n-r]%mod, mod)
}

func main() {
	sc.Split(bufio.ScanWords)

	fact[0] = 1
	for i := 1; i <= 100000; i++ {
		fact[i] = 1 * i * fact[i-1] % mod
	}

	n := scan()
	for i := 1; i <= n; i++ {
		ans := 0
		for j := 1; j <= (n+i-1)/i; j++ {
			ans += ncr(n-(i-1)*(j-1), j)
			ans %= mod
		}
		fmt.Println(ans)
	}
}
