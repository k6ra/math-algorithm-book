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
	y := scan()

	for a := 1; a <= n; a++ {
		for b := a; b <= n; b++ {
			for c := b; c <= n; c++ {
				for d := c; d <= n; d++ {
					if a+b+c+d == x && a*b*c*d == y {
						fmt.Println("Yes")
						return
					}
				}
			}
		}
	}

	fmt.Println("No")
	return
}
