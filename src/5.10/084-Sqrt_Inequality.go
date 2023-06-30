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

	a := scan()
	b := scan()
	c := scan()

	if c-a-b < 0 {
		fmt.Println("No")
		return
	}

	if 4*a*b < (c-a-b)*(c-a-b) {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
