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

	mod := 998244353
	a := scan()
	b := scan()
	c := scan()

	x := a * (a + 1) / 2 % mod
	y := b * (b + 1) / 2 % mod
	z := c * (c + 1) / 2 % mod

	ans := (x * y % mod) * z % mod
	fmt.Println(ans)
}
