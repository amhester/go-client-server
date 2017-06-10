package main

import (
	"fmt"

	"github.com/amhester/go-client-server/client/commander"
	"github.com/amhester/go-client-server/client/game"
)

func main() {
	g := game.NewGame()
	server := commander.NewCommandServer(g)
	exit := server.Start()
	fmt.Println(<-exit)
}
