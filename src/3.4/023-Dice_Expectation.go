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

func main() {
	sc.Split(bufio.ScanWords)

	n, err := scan()
	if err != nil {
		fmt.Printf("n err(%v)", err)
		return
	}

	b, err := calc(n)
	if err != nil {
		fmt.Printf("b err(%v)", err)
		return
	}

	r, err := calc(n)
	if err != nil {
		fmt.Println("r err(%v)", err)
		return
	}

	fmt.Println(b + r)
}

func calc(n int) (float64, error) {
	var result float64
	for i := 0; i < n; i++ {
		b, err := scan()
		if err != nil {
			return 0, err
		}

		result += float64(b) / float64(n)
	}

	return result, nil
}
