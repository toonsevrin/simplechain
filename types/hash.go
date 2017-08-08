package types

import (
	"encoding/json"
)

type Hash [32]byte


func (hash Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(hash[:])
}