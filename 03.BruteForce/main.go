package main

import (
	"fmt"
	"math"
)

/*
The number of total characters in an alphabet depends on how many bits each character in an encoded string can represent.

1 bit -> 2 chars (0 and 1, binary)
2 bits -> 4 chars
4 bits -> 16 chars (i.e. hexadecimal)
7 bits -> 128 chars (i.e. ASCII)
8 bits -> 256 chars
*/

func main() {
	fmt.Println(alphabetSize(4))
}

// alphabetSize takes returns the size of alphabet that can be represented
// based on the number of bits per a given character.
func alphabetSize(numBits int) float64 {
	return math.Pow(2, float64(numBits))
}
