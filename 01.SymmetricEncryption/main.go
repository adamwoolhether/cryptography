package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(generateRandomKey(16))
}

// generateRandomKey generates a random single-use hex key given the length.
// In production we should use crypto/rand, this is for demonstration purposes only.
func generateRandomKey(length int) (string, error) {
	rand.Seed(time.Now().Unix())

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", bytes), nil
}
