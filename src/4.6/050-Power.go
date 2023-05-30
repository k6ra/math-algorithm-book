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

	a := scan()
	b := scan()

	p := a
	ans := 1
	for i := 0; i < 30; i++ {
		if (b & (1 << i)) != 0 {
			ans *= p
			ans %= 1000000007
		}
		p *= p
		p %= 1000000007
	}

	fmt.Println(ans)
}
