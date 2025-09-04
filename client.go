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

	fmt.Print("Enter 32-byte hex key (same as server): ")
	keyHex, _ := reader.ReadString('\n')
	keyHex = strings.TrimSpace(keyHex)
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		panic("Invalid key. Must be 64 hex characters (32 bytes).")
	}

	fmt.Print("Enter server IP: ")
	serverIP, _ := reader.ReadString('\n')
	serverIP = strings.TrimSpace(serverIP)

	conn, err := net.Dial("tcp", serverIP+":9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to secure server.")

	// Start reading from server
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
					fmt.Println("Server:", string(decrypted))
				}
			}
		}
	}()

	// Send to server
	for {
		fmt.Print("You: ")
		msg, _ := reader.ReadString('\n')
		ciphertext, _ := encrypt(key, []byte(msg))
		conn.Write(ciphertext)
	}
	


	

}
