package main

import (
	"github.com/toonsevrin/simplechain/types"
	appPkg "github.com/toonsevrin/simplechain/app"
)
func main() {
	app := appPkg.App{Blockchain:[]types.Block{types.GetGenesis()},PeerAddresses:map[string]bool{}, Peers:map[string]*appPkg.Peer{}}
	server := appPkg.Server{app}
	server.Init()
}
