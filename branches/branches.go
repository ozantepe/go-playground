package main

import "fmt"

func main() {
	foo := 1

	if foo == 1 {
		fmt.Println("One")
	} else {
		fmt.Println("Not one")
	}

	bar := 2

	switch {
	case bar > 5:
		fmt.Println("Greater than five")
	case bar == 3:
		fmt.Println("Three")
	case bar > 1:
		fmt.Println("Greater than one")
	}
}
