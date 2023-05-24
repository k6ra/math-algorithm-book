package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	count100 := 0
	count200 := 0
	count300 := 0
	count400 := 0
	for _, v := range a {
		switch v {
		case 100:
			count100++
		case 200:
			count200++
		case 300:
			count300++
		case 400:
			count400++
		}
	}

	fmt.Println(count100*count400 + count200*count300)
}
