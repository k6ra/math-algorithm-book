package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() (int, error) {
	sc.Scan()
	return strconv.Atoi(sc.Text())
}

func main() {
	sc.Split(bufio.ScanWords)

	n, err := scan()
	if err != nil {
		return
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], err = scan()
		if err != nil {
			return
		}
	}

	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i], err = scan()
		if err != nil {
			return
		}
	}

	var ans float64
	for i := 0; i < n; i++ {
		ans += float64(a[i]) * 2 / 6
		ans += float64(b[i]) * 4 / 6
	}

	fmt.Println(ans)
}
