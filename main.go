package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

// will just create a random value like [1234, 3124, 4312, 4321]
func randomToken() string {
	token := make([]byte, 4)
	_, err := rand.Read(token)
	if err != nil {
		fmt.Println("Error creating random token", err)
		return ""
	}
	return hex.EncodeToString(token)
}

// will create a new file in /tmp with prefix d1_TIMESTAMP with a random value in each file
func writeFile() {
	token := randomToken()

	if token == "" {
		fmt.Println("Failed to create a random value token")
		return
	}

	filename := fmt.Sprintf("/tmp/d1_%d", time.Now().UnixNano())

	err := os.WriteFile(filename, []byte(token), 0644)
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	fmt.Println("File creted: ", filename)
}

func main() {
	writeFile()
}
