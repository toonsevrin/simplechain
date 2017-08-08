package types

import (
	"crypto/sha256"
	"encoding/binary"
	"bytes"
	"encoding/base64"
)

type Block struct {
	Index uint32//32bit
	PreviousHash Hash//256bit
	Timestamp int64//64 bit
	Hash Hash//256bit
	Data string//unlimited
}
type Hash [32]byte

func (hash Hash) MarshalJSON() ([]byte, error) {
	dest := make([]byte, len(hash))
	base64.StdEncoding.Encode(dest, hash[:])
	return dest, nil
}

func (block *Block) GenerateHash() *[32]byte{
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, block.Index)
	binary.Write(buf, binary.LittleEndian, block.PreviousHash)
	binary.Write(buf, binary.LittleEndian, block.Timestamp)
	binary.Write(buf, binary.LittleEndian, []byte(block.Data))

	res := sha256.Sum256(buf.Bytes())
	return &res
}

func GetGenesis() Block {
	block := Block{Index: 0, PreviousHash: &[32]byte{}, Timestamp: 1502089655, Hash: &[32]byte{}, Data: "This is the genesis block!"}
	block.Hash = block.GenerateHash()
	return block
}

func (block *Block) IsValid() bool{
	return(bytes.Equal(block.GenerateHash()[:], block.Hash[:]))
}


