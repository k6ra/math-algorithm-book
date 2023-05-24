package main

import "fmt"

func main() {
	var n, r int
	fmt.Scanf("%d %d", &n, &r)

	ans := fact(n) / (fact(r) * fact(n-r))
	fmt.Println(ans)
}

func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * fact(n-1)
}
