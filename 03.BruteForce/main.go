package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	fmt.Println(alphabetSize(4))
}

/*
The number of total characters in an alphabet depends on how many bits each character in an encoded string can represent.

1 bit -> 2 chars (0 and 1, binary)
2 bits -> 4 chars
4 bits -> 16 chars (i.e. hexadecimal)
7 bits -> 128 chars (i.e. ASCII)
8 bits -> 256 chars
*/

// alphabetSize takes returns the size of alphabet that can be represented
// based on the number of bits per a given character.
func alphabetSize(numBits int) float64 {
	return math.Pow(2, float64(numBits))
}

/*
A brute-force attack is simple: the attacker systematically guesses every possible combination.
 A 256-bit key is not twice as hard to guess as a 128-bit key, it is 2^128 times harder to guess!
*/

// findKey will attempt to brute force a given message, returning
// the key used to encrypt the password if found.
func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	for i := 0; i < int(math.Pow(2, 24)); i++ {
		b := intToBytes(i)

		if decrypt := crypt(encrypted, b); string(decrypt) == decrypted {
			return b, nil
		}
	}

	return nil, fmt.Errorf("unable to decrypt")
}

func crypt(dat, key []byte) []byte {
	final := []byte{}
	for i, d := range dat {
		final = append(final, d^key[i])
	}
	return final
}

func intToBytes(num int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int64(num))
	if err != nil {
		fmt.Println("Error in intToBytes:", err)
		return nil
	}
	bs := buf.Bytes()
	if len(bs) > 3 {
		return bs[:3]
	}
	return bs
}
