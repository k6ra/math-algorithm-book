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

	n := scan()
	x := scan()

	ans := 0
	for a := 1; a <= n; a++ {
		for b := a + 1; b <= n; b++ {
			for c := b + 1; c <= n; c++ {
				if a+b+c == x {
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
}
