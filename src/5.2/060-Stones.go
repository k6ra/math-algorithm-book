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

	if n%4 == 0 {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
	}
}
