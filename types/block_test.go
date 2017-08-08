package types

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
	"encoding/json"
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
	block := Block{PreviousHash: [32]byte{0}, Timestamp: 1502089655, Hash: [32]byte{}, Data: "This is the genesis block!"}
	assert.True(t, bytes.Equal(genesis.Hash[:], block.GenerateHash()[:]))
	block2 := Block{Index: 1, PreviousHash: [32]byte{0}, Timestamp: 1502089655, Hash: [32]byte{}, Data: "This is the genesis block!"}
	assert.False(t, bytes.Equal(genesis.Hash[:], block2.GenerateHash()[:]))
	block3 := Block{PreviousHash: [32]byte{0}, Timestamp: 1502089656, Hash: [32]byte{}, Data: "This is the genesis block!"}
	assert.False(t, bytes.Equal(genesis.Hash[:], block3.GenerateHash()[:]))
	block4 := Block{PreviousHash: [32]byte{1}, Timestamp: 1502089655, Hash: [32]byte{}, Data: "This is the genesis block!"}
	assert.False(t, bytes.Equal(genesis.Hash[:], block4.GenerateHash()[:]))
	block5 := Block{PreviousHash: [32]byte{0}, Timestamp: 1502089655, Hash: [32]byte{}, Data: "This is the genesis block!!"}
	assert.False(t, bytes.Equal(genesis.Hash[:], block5.GenerateHash()[:]))
	assert.True(t, genesis.IsValid())
}

func TestBlock_IsValid(t *testing.T) {
	falseBlock := Block{0, [32]byte{}, 1502089565, [32]byte{}, "test"}
	assert.False(t, falseBlock.IsValid())
	falseBlock.Hash = [32]byte{8,3,4,6,1,2,3,1,8,7,1,8,0,11,1,2,3,5,6,1}
	assert.False(t, falseBlock.IsValid())
	falseBlock.Hash = *falseBlock.GenerateHash()
	assert.True(t, falseBlock.IsValid())
}

func TestJsonParsing(t *testing.T){
    genesis := GetGenesis()
	bytes, err := json.Marshal(genesis)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, `{"Index":0,"PreviousHash":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=","Timestamp":1502089655,"Hash":"3HrtpJqhKq2UPoM6lpYop/ZcUR8br3Etym5JcJJ+H1A=","Data":"This is the genesis block!"}`, string(bytes))
}