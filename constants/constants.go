package main

import (
	"fmt"
)

const (
	first = iota
	second
	third
)

const (
	fourth = iota
)

func main() {
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
}
