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

func abs(a int) int {
	if a > 0 {
		return a
	}

	return a * -1
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()
	x := scan()
	y := scan()

	absXY := abs(x) + abs(y)
	if absXY <= n && absXY%2 == n%2 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
