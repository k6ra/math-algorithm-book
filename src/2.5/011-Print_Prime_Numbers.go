package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	ans := []string{}
	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	fmt.Println(strings.Join(ans, " "))
}
