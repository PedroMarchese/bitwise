package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// Ok, I'll create a new variable to operate
	var num uint = 7

	// Do you wanna count bits? Ok, here's the easiest way
	numberOfBits := bits.OnesCount(num)

	fmt.Println("We have:", numberOfBits, "bits!")
}
