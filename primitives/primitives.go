package main

import (
	"fmt"
)

func main() {
	myInt := 2
	var myFloat float32 = 2
	myOtherFloat := 3.
	myStr := "Ozan"

	fmt.Println(myInt)
	fmt.Printf("Float = %f\n", myFloat)
	fmt.Printf("OtherFloat = %.2f\n", myOtherFloat)
	fmt.Println(myStr)
}
