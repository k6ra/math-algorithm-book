package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	ans := a[0]
	for i := 1; i < n; i++ {
		for ans > 0 && a[i] > 0 {
			if ans > a[i] {
				ans = ans % a[i]
			} else {
				a[i] = a[i] % ans
			}
		}

		if a[i] > 0 {
			ans = a[i]
		}
	}

	fmt.Println(ans)
}
