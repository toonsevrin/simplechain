package types

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
)

//func TestBlock_Hash(t *testing.T) {
//	genesis := getGenesis()
//	block := Block{1,genesis.hash, 1502106792, nil, "Second block!" }
//	hash := block.hash
//
//
//}

func TestGetGenesis(t *testing.T) {
	genesis := GetGenesis()
	block := Block{0, &[32]byte{0}, 1502089655, &[32]byte{}, "This is the genesis block!"}
	assert.True(t, bytes.Equal(genesis.hash[:], block.Hash()[:]))
	block2 := Block{1, &[32]byte{0}, 1502089655, &[32]byte{}, "This is the genesis block!"}
	assert.False(t, bytes.Equal(genesis.hash[:], block2.Hash()[:]))
	block3 := Block{0, &[32]byte{0}, 1502089656, &[32]byte{}, "This is the genesis block!"}
	assert.False(t, bytes.Equal(genesis.hash[:], block3.Hash()[:]))
	block4 := Block{0, &[32]byte{1}, 1502089655, &[32]byte{}, "This is the genesis block!"}
	assert.False(t, bytes.Equal(genesis.hash[:], block4.Hash()[:]))
	block5 := Block{0, &[32]byte{0}, 1502089655, &[32]byte{}, "This is the genesis block!!"}
	assert.False(t, bytes.Equal(genesis.hash[:], block5.Hash()[:]))
}