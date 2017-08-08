package types

import (
	"encoding/base64"
)

type Hash [32]byte


func (hash Hash) MarshalJSON() ([]byte, error) {
	dest := make([]byte, len(hash))
	base64.StdEncoding.Encode(dest, hash[:])
	return dest, nil
}