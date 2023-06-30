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

	mod := 1000000007
	n := scan()

	ans := (n * (n + 1) / 2 % mod) * (n * (n + 1) / 2 % mod) % mod
	fmt.Println(ans)
}
