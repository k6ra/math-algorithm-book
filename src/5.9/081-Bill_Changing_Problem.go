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

	ans := n / 10000
	remain := n % 10000

	ans += remain / 5000
	remain = remain % 5000

	ans += remain / 1000

	fmt.Println(ans)
}
