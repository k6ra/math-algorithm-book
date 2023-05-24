package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	ans := lcm(a[0], a[1])
	for i := 2; i < n; i++ {
		ans = lcm(ans, a[i])
	}

	fmt.Println(ans)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func gcd(a, b int) int {
	for a > 0 && b > 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	if a > 0 {
		return a
	}
	return b
}
