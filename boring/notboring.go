//go:build !boringcrypto

// Package boring is a very thin wrapper around crypto/boring that has the negative build tag to be
// IDE friendly.
package boring

import "log"

func init() {
	log.Printf("----- in boring.go -----")
}

// Enabled reports whether BoringCrypto handles supported crypto operations.
func Enabled() bool {
	return false
}
