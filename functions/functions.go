package main

import "fmt"

func main() {
	sayHello("hello", "world")
	message := "Message"
	sayHelloToPointer(&message)
	fmt.Println(message)

	numberOfTerms, sum := returnMultipleValues(1, 2, 3, 4, 5)
	fmt.Println(numberOfTerms)
	fmt.Println(sum)
}

func sayHello(messages ...string) {
	for _, message := range messages {
		fmt.Println(message)
	}
}

func sayHelloToPointer(message *string) {
	fmt.Println(*message)
	*message = "Messagezzz"
}

func returnMultipleValues(terms ...int) (numberOfTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numberOfTerms = len(terms)
	return
}
