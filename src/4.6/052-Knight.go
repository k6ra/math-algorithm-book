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

	x := scan()
	y := scan()

	if y*2-x < 0 || 2*x-y < 0 {
		fmt.Println(0)
		return
	}

	if (2*y-x)%3 != 0 || (2*x-y)%3 != 0 {
		fmt.Println(0)
		return
	}

	const mod = 1000000007
	bunshi := 1
	bunbo := 1
	a := (2*y - x) / 3
	b := (2*x - y) / 3

	for i := 1; i <= a+b; i++ {
		bunshi *= i
		bunshi %= mod
	}

	for i := 1; i <= a; i++ {
		bunbo *= i
		bunbo %= mod
	}

	for i := 1; i <= b; i++ {
		bunbo *= i
		bunbo %= mod
	}

	p := bunbo
	bunboMSub2 := 1
	for i := 0; i < 30; i++ {
		if ((mod - 2) & (1 << i)) != 0 {
			bunboMSub2 *= p
			bunboMSub2 %= mod
		}
		p *= p
		p %= mod
	}

	fmt.Println((bunshi * bunboMSub2) % mod)
}
