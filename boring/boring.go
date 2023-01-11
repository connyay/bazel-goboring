//go:build boringcrypto

package boring

import (
	"log"

	"crypto/boring"
)

func init() {
	log.Printf("----- in boring.go -----")
}

// Enabled reports whether BoringCrypto handles supported crypto operations.
func Enabled() bool {
	return boring.Enabled()
}
