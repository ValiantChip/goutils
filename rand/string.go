package rand

import (
	"crypto/rand"
	"io"
	"math/big"
)

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
