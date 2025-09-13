package rand

import (
	"crypto/rand"
	"io"
	"math/big"
)

const (
	UPPER_CASE    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LOWER_CASE    = "abcdefghijklmnopqrstuvwxyz"
	NUMERIC       = "0123456789"
	ALPHA_NUMERIC = UPPER_CASE + LOWER_CASE + NUMERIC
)

// Returns a slice of bytes of the length provided using bytes uniformly selected from the provided dictionary
func Bytes(reader io.Reader, dictionary []byte, length int64) ([]byte, error) {
	b := make([]byte, length)

	for i := range length {
		k, err := rand.Int(reader, big.NewInt(int64(len(dictionary))))
		if err != nil {
			return b, err
		}

		b[i] = dictionary[k.Int64()]
	}

	return b, nil
}
