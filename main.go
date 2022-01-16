package main

import (
	"fmt"

	"github.com/quinn/algorithm-practice/quicksort"
)

func main() {
	ints := []int{1, 8, 3, 47, 2, 5, 7}

	quicksort.Sort(ints)

	fmt.Println(ints)
}
