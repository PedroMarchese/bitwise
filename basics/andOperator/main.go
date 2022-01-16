package main

import "fmt"

func main() {
	/*
		& operator is knowed as AND operator
		It returns the interception between two bit sequence
	*/

	var firstNum uint16 = 2
	var secondNum uint16 = 3

	// Here we are getting the intersection between 10 & 11
	intersection := firstNum & secondNum

	// Let's see the result
	// 10 & 11 = 10
	fmt.Printf("%08b\n", intersection)
}
