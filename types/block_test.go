package types

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, 0, genesis.index)
	assert.Equal(t, nil, genesis.previousHash)
	assert.Equal(t, genesis.Hash(), genesis.hash)



}