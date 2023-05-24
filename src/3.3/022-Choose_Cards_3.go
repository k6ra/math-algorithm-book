package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		return
	}

	count := make([]int, 100000)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, err := strconv.Atoi(sc.Text())
		if err != nil {
			return
		}

		count[a]++
	}

	ans := 0
	for i := 1; i < 50000; i++ {
		ans += count[i] * count[100000-i]
	}

	ans += count[50000] * (count[50000] - 1) / 2

	fmt.Println(ans)
}
