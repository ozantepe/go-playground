package main

import (
	"fmt"
)

func main() {

	myArray := [...]int{1, 2, -1, 0}

	fmt.Println(myArray)
	fmt.Println(len(myArray))

	mySlice := myArray[2:]
	fmt.Println(mySlice)

	myMap := make(map[int]string)
	fmt.Println(myMap)

	myMap[3] = "Third"
	myMap[7] = "Seven"

	fmt.Println(myMap)
	fmt.Println(myMap[3])
	fmt.Println(myMap[99])
}
