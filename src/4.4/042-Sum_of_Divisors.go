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
	n := scan()

	list := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; i*j <= n; j++ {
			list[i*j]++
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += i * list[i]
	}

	fmt.Println(ans)
}
