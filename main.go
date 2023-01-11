package main

import (
	"fmt"
	"os"

	"crypto/boring"
)

func main() {
	fmt.Printf("goboring enabled=%t\n", boring.Enabled())
	if !boring.Enabled() {
		os.Exit(1)
	}
}
