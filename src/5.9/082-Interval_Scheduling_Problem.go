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

type movie struct {
	start int
	end   int
}

func main() {
	sc.Split(bufio.ScanWords)

	n := scan()

	movies := make([]movie, n)
	for i := 0; i < n; i++ {
		l := scan()
		r := scan()
		movies[i] = movie{
			start: l,
			end:   r,
		}
	}

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].end < movies[j].end
	})

	cnt := 1
	current := movies[0].end
	for i := 1; i < n; i++ {
		if movies[i].start < current {
			continue
		}

		current = movies[i].end
		cnt++
	}

	fmt.Println(cnt)
}
