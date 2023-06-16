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
	k := scan()

	sum := 0
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scan()
		sum += a[i]
	}

	if sum%2 != k%2 || sum > k {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
