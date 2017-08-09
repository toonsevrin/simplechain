package app

import (
	"github.com/toonsevrin/simplechain/types"
	"time"
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
)

type App struct {
	Blockchain    []types.Block
	Peers         map[string]*Peer //defined by url
	PeerAddresses map[string]string //defined by ip, reference to url
}

func (app *App) createAndAddNextBlock(data string) types.Block {
	previous := app.getLatestBlock()
	next := types.Block{Index: previous.Index + 1, PreviousHash: previous.Hash, Timestamp: time.Now().Unix(), Hash: [32]byte{}, Data: data}
	next.Hash = *next.GenerateHash();
	app.Blockchain = append(app.Blockchain, next)
	return next
}
func (app *App) HasBlock(block types.Block) bool {
	return uint32(len(app.Blockchain)) > block.Index && app.Blockchain[block.Index].Hash == block.Hash
}

func (app *App) AddBlock(block types.Block) {
	app.Blockchain = append(app.Blockchain, block)
}
func (app *App) isValidNewBlock(block types.Block) bool {
	return block.IsValid() && bytes.Equal(block.PreviousHash[:], app.getLatestBlock().Hash[:]) && app.getLatestBlock().Index+1 == block.Index
}

func (app *App) getLatestBlock() *types.Block {
	return &app.Blockchain[len(app.Blockchain)-1]
}

func (app *App) isValidChain(chain []types.Block) bool {
	if *chain[0].GenerateHash() != app.Blockchain[0].Hash {
		return false
	}
	for i, block := range chain {
		if (i == 0) {
			continue
		}
		if block.PreviousHash != chain[i-1].Hash {
			return false
		}
		if !block.IsValid() {
			return false
		}
	}
	return true
}

func (app *App) pickLongestChain(newChain []types.Block) bool {
	if len(newChain) > len(app.Blockchain) && app.isValidChain(newChain) {
		fmt.Println("Received longer chain of length ", len(newChain))
		app.Blockchain = newChain
		return true
	} else {
		fmt.Println("Received invalid chain")
		return false
	}
}

func (app *App) broadcast(block types.Block) {
	marshalled, err := json.Marshal(block)
	if err != nil {
		panic(err.Error())
	}
	for _, peer := range app.Peers {
		req, err := http.NewRequest("POST", peer.getUrl()+"/addBlock", bytes.NewReader(marshalled))
		if err != nil {
			log.Println(err.Error())
			return
		}
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			log.Println(err.Error())
			return
		}

		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("Response when sharing new block: " + string(bytes))
	}
}
