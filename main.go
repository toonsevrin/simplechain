package main

import (
	"github.com/toonsevrin/simplechain/types"
	"github.com/toonsevrin/simplechain/app"
)
func main() {
	app := app.App{ []types.Block{types.GetGenesis()}}
	server := app.Server{app}
	server.Init()
}
