package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	sc.Split(bufio.ScanWords)

	h := scan()
	w := scan()
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			a[i][j] = scan()
		}
	}

	row := make([]int, h)
	for i := 0; i < h; i++ {
		sum := 0
		for _, v := range a[i] {
			sum += v
		}
		row[i] = sum
	}

	col := make([]int, w)
	for i := 0; i < w; i++ {
		sum := 0
		for j := 0; j < h; j++ {
			sum += a[j][i]
		}
		col[i] = sum
	}

	for i := 0; i < h; i++ {
		ans := make([]string, w)
		for j := 0; j < w; j++ {
			ans[j] = strconv.Itoa(row[i] + col[j] - a[i][j])
		}
		fmt.Println(strings.Join(ans, " "))
	}
}
