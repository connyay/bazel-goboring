//go:build red

package buildcolor

import "log"

func init() {
	log.Printf("----- in red.go -----")
}

func Color() string {
	return "red"
}
