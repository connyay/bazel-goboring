package main

import (
	"fmt"
	"os"

	"github.com/connyay/bazel-goboring/boring"
)

func main() {
	fmt.Printf("goboring enabled=%t\n", boring.Enabled())
	if !boring.Enabled() {
		os.Exit(1)
	}
}
