package main

import (
	"fmt"
)

func main() {
	length, capacity, y := 1000, 2000, 3
	AllResults := make([]int, length, capacity)
	for i := 0; i < length; i++ {
		result := i / y // Hi there
		fmt.Println(result)
		AllResults[i] = result
	}
	j := 0
	fmt.Println(j)

	fmt.Println()

	const query string = `INSERT INTO table VALUES($1, $2, $3)`
	_ = query
}
