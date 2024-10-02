package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("failed to read data: %v", err)
		return
	}

	fmt.Printf("Input: %#v | len(%d)", string(data), len(data))
}
