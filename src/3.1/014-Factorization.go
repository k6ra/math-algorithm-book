package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, i int
	fmt.Scanf("%d", &n)
	var ans []string
	for i = 2; i*i <= n; {
		if n%i == 0 {
			ans = append(ans, strconv.Itoa(i))
			n = n / i
			i = 2
		} else {
			i++
		}
	}
	if n != 1 {
		ans = append(ans, strconv.Itoa(n))
	}
	fmt.Println(strings.Join(ans, " "))
}
