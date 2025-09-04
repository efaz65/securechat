package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	// Generate 32 random bytes
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	// Print as hex string
	fmt.Println("Your 32-byte hex key (copy this into both server and client):")
	fmt.Println(hex.EncodeToString(key))
}
