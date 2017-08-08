package types

import (
	"encoding/json"
)

type Hash [32]byte


func (hash Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(hash[:])
}

func (h *Hash) UnmarshalJSON(b []byte) error {
	var s []byte
	err := json.Unmarshal(b, &s)
	copy(h[:], s)
	return err
}