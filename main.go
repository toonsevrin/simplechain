package main

import (
	"github.com/toonsevrin/simplechain/types"
	appPkg "github.com/toonsevrin/simplechain/app"
)
func main() {
	app := appPkg.App{ []types.Block{types.GetGenesis()}}
	server := appPkg.Server{app}
	server.Init()
}
