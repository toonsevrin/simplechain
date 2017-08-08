package app

import (
	"github.com/toonsevrin/simplechain/types"
	"time"
	"bytes"
	"fmt"
)
type App struct {
	Blockchain []types.Block
}

func (app *App) createNextBlock(data string) types.Block {
	previous := app.getLatestBlock()
	next := types.Block{Index: previous.Index, PreviousHash: previous.Hash, Timestamp: time.Now().Unix(), Hash: &[32]byte{}, Data: data}
	next.Hash = next.GenerateHash();
	return next
}

func (app *App) isValidNewBlock(block types.Block) bool{
	return block.IsValid() && bytes.Equal(block.PreviousHash[:], app.getLatestBlock().Hash[:]) && app.getLatestBlock().Index + 1 == block.Index
}

func (app *App) getLatestBlock() types.Block {
	return app.Blockchain[len(app.Blockchain)-1]
}

func (app *App) isValidChain(chain []types.Block) bool{
	if chain[0].GenerateHash() != app.Blockchain[0].Hash {
		return false
	}
	for i, block := range chain {
		if(i == 0){
			continue
		}
		if block.PreviousHash != chain[i - 1].Hash {
			return false
		}
		if !block.IsValid() {
			return false
		}
	}
	return true
}

func (app *App) pickLongestChain(newChain []types.Block)  {
	if len(newChain) > len(app.Blockchain) && isValidChain(newChain) {
		fmt.Println("Received longer chain of length ", len(newChain))
		app.Blockchain = newChain
		//broadcast new chain
	}else{
		fmt.Println("Received invalid chain")
	}
}