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
	var ans float64
	for i := 0; i < n; i++ {
		p := scan()
		q := scan()

		ans += float64(q) / float64(p)
	}

	fmt.Printf("%g", ans)
}
