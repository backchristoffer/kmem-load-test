package main

import (
	"crypto/rand"
	"fmt"
)

func randomToken() {
	token := make([]byte, 4)
	_, err := rand.Read(token)
	if err != nil {
		fmt.Println("Error creating token", err)
		return
	}
	fmt.Println(token)
}

func main() {
	randomToken()
}
