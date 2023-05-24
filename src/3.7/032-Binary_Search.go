package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func binarySearch(arr []int, left int, right int, value int) bool {
	if left > right {
		return false
	}

	mid := left + (right-left)/2
	if arr[mid] == value {
		return true
	}

	if arr[mid] < value {
		return binarySearch(arr, mid+1, right, value)
	}
	return binarySearch(arr, left, mid-1, value)
}

func main() {
	sc.Split(bufio.ScanWords)
	n := scan()
	x := scan()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scan()
	}

	sort.Slice(a, func(i int, j int) bool {
		return a[i] < a[j]
	})

	if binarySearch(a, 0, n-1, x) {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
