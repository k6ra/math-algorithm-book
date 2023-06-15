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

	for i := 1; i < 60; i++ {
		if n == (1<<i)-1 {
			fmt.Println("Second")
			return
		}
	}

	fmt.Println("First")
}
