package main

import (
	"fmt"
	"os"
)

func main() {
	ls, err := NewLocalStorage("poketcg.db")
	if err != nil {
		panic(err)
	}
	InitSeed()
	game := NewGame(ls)
	server := NewCommander(ls, game)
	res, err := server.ResolveCommand(os.Args[1:]...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		fmt.Println(res)
	}
}
