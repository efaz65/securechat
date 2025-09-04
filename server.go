package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	
	
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter 32-byte hex key (same as client): ")
	keyHex, _ := reader.ReadString('\n')
	keyHex = strings.TrimSpace(keyHex)
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		panic("Invalid key. Must be 64 hex characters (32 bytes).")
	}

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Secure server started on port 9000")

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Start reading from client
	go func() {
		for {
			buffer := make([]byte, 4096)
			n, err := conn.Read(buffer)
			if err != nil {
				return
			}
			if n > 0 {
				decrypted, err := decrypt(key, buffer[:n])
				if err == nil {
					fmt.Println("Client:", string(decrypted))
				}
			}
		}
	}()

	// Send to client
	for {
		fmt.Print("You: ")
		msg, _ := reader.ReadString('\n')
		ciphertext, _ := encrypt(key, []byte(msg))
		conn.Write(ciphertext)
	}
	
	

}
