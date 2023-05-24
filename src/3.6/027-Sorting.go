package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() (int, error) {
	sc.Scan()
	return strconv.Atoi(sc.Text())
}

var a = make([]int, 200009)
var c = make([]int, 200009)

func mergeSort(l int, r int) {
	if r-l == 1 {
		return
	}

	m := (l + r) / 2
	mergeSort(l, m)
	mergeSort(m, r)

	c1 := l
	c2 := m
	cnt := 0
	for c1 != m || c2 != r {
		if c1 == m {
			c[cnt] = a[c2]
			c2++
		} else if c2 == r {
			c[cnt] = a[c1]
			c1++
		} else {
			if a[c1] < a[c2] {
				c[cnt] = a[c1]
				c1++
			} else {
				c[cnt] = a[c2]
				c2++
			}
		}
		cnt++
	}

	for i := 0; i < cnt; i++ {
		a[l+i] = c[i]
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n, err := scan()
	if err != nil {
		fmt.Printf("error n: %v", err)
		return
	}

	for i := 1; i <= n; i++ {
		v, err := scan()
		if err != nil {
			fmt.Printf("error a: %v", err)
			return
		}
		a[i] = v
	}

	mergeSort(1, n+1)
	for i := 1; i <= n; i++ {
		fmt.Println(a[i])
	}
}
