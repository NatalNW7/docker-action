package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting action...")
	fmt.Printf("Hello, %s!\n", os.Args[1])
}
