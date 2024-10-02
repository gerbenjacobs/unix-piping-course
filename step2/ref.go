package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Output struct {
	Message string `json:"message"`
	Length  int    `json:"len"`
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("failed to read data: %v", err)
		return
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(Output{
		Message: string(data),
		Length:  len(data),
	})
	if err != nil {
		fmt.Printf("failed to encode json: %v", err)
	}
}
