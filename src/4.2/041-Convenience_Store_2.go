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

	t := scan()
	n := scan()
	diff := make([]int, t+1)
	for i := 0; i < n; i++ {
		l := scan()
		r := scan()

		diff[l]++
		diff[r]--
	}

	sum := 0
	for i := 0; i < t; i++ {
		fmt.Println(sum + diff[i])
		sum += diff[i]
	}
}
