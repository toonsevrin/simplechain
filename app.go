package simplechain

import (
	"github.com/toonsevrin/simplechain/types"
)

var blockchain []types.Block = []types.Block{types.GetGenesis()}

func main(){

}

//func createNextBlock(data string) *types.Block {
//	previous := getLatestBlock()
//}

func getLatestBlock() types.Block {
	return blockchain[len(blockchain)-1]
}