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

	h := scan()
	w := scan()

	if h == 1 || w == 1 {
		fmt.Println(1)
		return
	}

	ans := h * w / 2
	if h*w%2 != 0 {
		ans++
	}

	fmt.Println(ans)
}
