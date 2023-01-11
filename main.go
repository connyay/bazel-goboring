package main

import (
	"fmt"

	"github.com/connyay/bazel-goboring/boring"
)

func main() {
	fmt.Printf("goboring enabled=%t\n", boring.Enabled())
}
