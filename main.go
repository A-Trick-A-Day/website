package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println(slices.Contains([]int64{1, 2, 3}, 1))
	fmt.Println(slices.Contains([]float64{1, 2, 3}, 1))
}
