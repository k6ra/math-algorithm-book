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

	if c == 1 {
		fmt.Println("No")
		return
	}

	cb := 1
	for i := 1; i <= b; i++ {
		if a/c < cb {
			fmt.Println("Yes")
			return
		}
		cb *= c
	}

	fmt.Println("No")
}
