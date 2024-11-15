package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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
func writeFile(osDir string) (string, error) {
	token := randomToken()

	if token == "" {
		fmt.Println("Failed to create a random value token")
		return "", fmt.Errorf("failed to create token")
	}

	filename := fmt.Sprintf("%s/d1_%d", osDir, time.Now().UnixNano())

	err := os.WriteFile(filename, []byte(token), 0644)
	if err != nil {
		fmt.Println("Error creating file", err)
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	fmt.Println("File created: ", filename)
	return filename, nil
}

// handler for creating files with api endpoint
func api_createfile(c *gin.Context) {
	filename, err := writeFile("/tmp")
	if err != nil {
		c.String(500, fmt.Sprintf("Error: %v", err))
	}
	c.String(200, fmt.Sprintf("File created at: %s", filename))
}

func main() {
	route := gin.Default()
	route.GET("/file", api_createfile)
	route.Run(":8080")
	fmt.Println("Starting server on 8080")
}
