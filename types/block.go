package types

import (
	"crypto/sha256"
	"encoding/binary"
	"bytes"
)

type Block struct {
	index uint32//32bit
	previousHash *[32]byte//256bit
	timestamp uint64//64 bit
	hash *[32]byte//256bit
	data string//unlimited
}

func (block *Block) Hash() *[32]byte{
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, block.index)
	binary.Write(buf, binary.LittleEndian, block.previousHash)
	binary.Write(buf, binary.LittleEndian, block.timestamp)
	binary.Write(buf, binary.LittleEndian, []byte(block.data))

	res := sha256.Sum256(buf.Bytes())
	return &res
}

func GetGenesis() Block {
	block := Block{index: 0, previousHash: &[32]byte{}, timestamp: 1502089655, hash: &[32]byte{}, data: "This is the genesis block!"}
	block.hash = block.Hash()
	return block
}