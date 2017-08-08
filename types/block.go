package types

import (
	"crypto/sha256"
	"encoding/binary"
	"bytes"
)

type Block struct {
	Index uint32//32bit
	PreviousHash Hash//256bit
	Timestamp int64//64 bit
	Hash Hash//256bit
	Data string//unlimited
}


func (block *Block) GenerateHash() *Hash{
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, block.Index)
	binary.Write(buf, binary.LittleEndian, block.PreviousHash)
	binary.Write(buf, binary.LittleEndian, block.Timestamp)
	binary.Write(buf, binary.LittleEndian, []byte(block.Data))


	var res Hash = sha256.Sum256(buf.Bytes())
	return &res
}

func GetGenesis() Block {
	block := Block{Index: 0, PreviousHash: [32]byte{}, Timestamp: 1502089655, Hash: [32]byte{}, Data: "This is the genesis block!"}

	block.Hash = *block.GenerateHash()
	return block
}

func (block *Block) IsValid() bool{
	return bytes.Equal(block.GenerateHash()[:], block.Hash[:])
}


