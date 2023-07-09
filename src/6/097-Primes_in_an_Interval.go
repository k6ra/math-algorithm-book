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

	l := scan()
	r := scan()

	prime := make([]bool, r-l+1)
	for i := range prime {
		prime[i] = true
	}
	if l == 1 {
		prime[0] = false
	}

	for i := 2; i*i <= r; i++ {
		min := ((l + i - 1) / i) * i
		for j := min; j <= r; j += i {
			if j == i {
				continue
			}
			prime[j-l] = false
		}
	}

	ans := 0
	for i := range prime {
		if prime[i] {
			ans++
		}
	}

	fmt.Println(ans)
}
