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

	if n%4 == 1 {
		fmt.Println(2)
		return
	}
	if n%4 == 2 {
		fmt.Println(4)
		return
	}
	if n%4 == 3 {
		fmt.Println(8)
		return
	}
	fmt.Println(6)
}
