package main

import (
	"fmt"
	"os"
)

func hello(msg string) {
	fmt.Println("Message:", msg)
}

func main() {
	msg := os.Getenv("MSG")
	hello(msg)
}
