package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
			fmt.Println(n / i)
		}
	}
}
