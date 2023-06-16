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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scan()
	}

	first := make([]int, n)
	second := make([]int, n)
	for i := 0; i < n; i++ {
		first[i] = -1
		second[i] = -1
	}

	current := 0
	count := 0
	for {
		if first[current] == -1 {
			first[current] = count
		} else if second[current] == -1 {
			second[current] = count
		}

		if count == k {
			fmt.Println(current + 1)
			return
		} else if second[current] != -1 {
			if (k-first[current])%(second[current]-first[current]) == 0 {
				fmt.Println(current + 1)
				return
			}
		}

		current = a[current] - 1
		count++
	}
}
