package main

import (
	"bufio"
	"fmt"
	"math"
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

	a := float64(scan())
	b := float64(scan())

	ans := 0
	for i := 1.0; i <= float64(b); i++ {
		if math.Floor(b/i)-math.Ceil(a/i) >= 1 {
			ans = int(i)
		}
	}

	fmt.Println(ans)
}
