package main

import "fmt"

func main() {
	/*
		<< and >> operators are knowed as shift operators
	*/

	// Here i'm declaring an integer composed of 8 bits
	var binaryNumber uint8 = 1

	// Now the number is being formated and printed as a 8bit sequence
	fmt.Printf("%08b\n", binaryNumber)

	// 8 bits = 1 byte! Now we gonna have a shifted 8 bit sequence
	shiftedByte := binaryNumber << 2

	// Let's print it as we already learned
	fmt.Printf("%08b\n", shiftedByte)
}
