package types

import (
	"crypto/sha256"
	"encoding/binary"
	"bytes"
)

type Block struct {
	index uint//32bit
	previousHash [32]byte//256bit
	timestamp uint64//64 bit
	hash [32]byte//256bit
	data string//unlimited
}

func (block *Block) Hash() [32]byte{
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, block.index)
	binary.Write(buf, binary.LittleEndian, block.previousHash)
	binary.Write(buf, binary.LittleEndian, block.timestamp)
	binary.Write(buf, binary.LittleEndian, block.hash)
	binary.Write(buf, binary.LittleEndian, block.data)

	res := sha256.Sum256(buf.Bytes())
	return res
}

func GetGenesis() Block {
	block := Block{0, nil, 1502089655, nil, "This is the genesis block!"}
	block.hash = block.Hash()
	return block
}