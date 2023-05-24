package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	for a > 0 && b > 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}

	ans := a
	if b > 0 {
		ans = b
	}
	fmt.Println(ans)

}
