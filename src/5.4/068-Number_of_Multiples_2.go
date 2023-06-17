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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()
	k := scan()
	v := make([]int, k+1)
	for i := 1; i <= k; i++ {
		v[i] = scan()
	}

	ans := 0
	for i := 1; i < (1 << k); i++ {
		count := 0
		l := 1
		for j := 0; j < k; j++ {
			if (i & (1 << j)) != 0 {
				count++
				l = lcm(l, v[j+1])
			}
		}
		num := n / l
		if count%2 == 1 {
			ans += num
		} else {
			ans -= num
		}
	}

	fmt.Println(ans)
}
