//go:build blue

package buildcolor

import "log"

func init() {
	log.Printf("----- in blue.go -----")
}

func Color() string {
	return "blue"
}
