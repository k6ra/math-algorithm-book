package main

import (
	"bufio"
	"fmt"
	"os"
)

func ncr(x, y int) int {
	if x < 3 && y < 3 {
		if x < y {
			return 0
		}
		if x == 2 && y == 1 {
			return 2
		}
		return 1
	}
	return ncr(x/3, y/3) * ncr(x%3, y%3) % 3
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	var c string
	fmt.Fscan(in, &n, &c)

	ans := 0
	for i, v := range c {
		code := 0
		switch v {
		case 'B':
			code = 0
		case 'W':
			code = 1
		case 'R':
			code = 2
		}
		ans += code * ncr(n-1, i)
		ans %= 3
	}

	if n%2 == 0 {
		ans = (3 - ans) % 3
	}

	fmt.Println(string("BWR"[ans]))
}
