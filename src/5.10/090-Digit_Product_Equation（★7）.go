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

var set map[int]struct{}

func f(digit, m int) {
	if digit == 11 {
		set[product(m)] = struct{}{}
		return
	}

	minVal := m % 10
	for i := minVal; i <= 9; i++ {
		f(digit+1, 10*m+i)
	}
}

func product(m int) int {
	if m == 0 {
		return 0
	}

	ans := 1
	for m >= 1 {
		ans *= m % 10
		m /= 10
	}

	return ans
}

func main() {
	sc.Split(bufio.ScanWords)

	set = make(map[int]struct{}, 300000)
	f(0, 0)

	n := scan()
	b := scan()

	ans := 0
	for k := range set {
		m := k + b
		prodM := product(m)
		if m-prodM == b && m <= n {
			ans++
		}
	}

	fmt.Println(ans)
}
