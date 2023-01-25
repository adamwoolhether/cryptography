package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

func main() {
	helloWorldBytes := []byte("Hello world")
	fmt.Println(helloWorldBytes)

	hexString := getHexString(helloWorldBytes)
	fmt.Println(hexString)

	binaryString := getBinaryString(helloWorldBytes)
	fmt.Println(binaryString)

	hexDecoded, err := getHexBytes(hexString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hexDecoded)
}

// getHexString will take a byte slice and return it's hexadecimal form
// separated by a colon.
func getHexString(b []byte) string {
	result := make([]string, 0)
	for _, bin := range b {
		result = append(result, fmt.Sprintf("%02x", bin))
	}

	return strings.Join(result, ":")
}

// getBinaryString will take a byte slice and return it's binary form
// separated by a colon.
func getBinaryString(b []byte) string {
	result := make([]string, 0)
	for _, bin := range b {
		result = append(result, fmt.Sprintf("%08b", bin))
	}

	return strings.Join(result, ":")
}

// getHexBytes will take hex encoded string, and return the bytes decoded.
func getHexBytes(s string) ([]byte, error) {
	pieces := strings.Split(s, ":")
	result := make([]byte, 0)
	for _, p := range pieces {
		dec, err := hex.DecodeString(p)
		if err != nil {
			return nil, err
		}
		result = append(result, dec...)
	}

	return result, nil
}
