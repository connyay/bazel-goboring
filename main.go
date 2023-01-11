package main

import (
	"fmt"
	"os"

	"github.com/connyay/bazel-goboring/boring"
	"github.com/connyay/bazel-goboring/buildcolor"
)

func main() {
	fmt.Printf("buildcolor=%s", buildcolor.Color())
	fmt.Printf("goboring enabled=%t\n", boring.Enabled())
	if !boring.Enabled() {
		os.Exit(1)
	}
}
