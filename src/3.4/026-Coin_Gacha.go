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
		fmt.Println(err)
	}

	var ans float64
	for i := 0.0; i < float64(n); i++ {
		ans += float64(1.0 / (1.0 - i/float64(n)))
	}
	fmt.Println(ans)
}
