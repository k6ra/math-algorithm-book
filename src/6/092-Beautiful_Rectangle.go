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

	ans := 1 << 60
	for h := 1; h*h <= n; h++ {
		if n%h != 0 {
			continue
		}
		w := n / h
		size := h*2 + w*2
		if ans > size {
			ans = size
		}
	}

	fmt.Println(ans)
}
