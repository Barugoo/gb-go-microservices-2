package main

import "fmt"

func Average(elems ...int) int {
	var sum int
	for _, elem := range elems {
		sum += elem
	}
	return sum / len(elems)
}

func main() {
	fmt.Println("Hello world!")
}
