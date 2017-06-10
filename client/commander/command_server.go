package commander

import (
	"fmt"
	"regexp"
	"time"

	"github.com/amhester/go-client-server/client/game"
)

var passStars = regexp.MustCompile(".")
var Game *game.Game

type CommandServer struct {
	Exit chan error
}

func NewCommandServer(g *game.Game) *CommandServer {
	Game = g
	return &CommandServer{
		Exit: make(chan error),
	}
}

func (server *CommandServer) Start() chan error {
	go server.init()
	return server.Exit
}

func (server *CommandServer) init() {
	var name string
	var pass string
	fmt.Println("Starting server...")
	time.Sleep(time.Second)
	fmt.Println("Enter username: ")
	fmt.Scanln(&name)
	fmt.Println("Enter pasword: ")
	fmt.Scanln(&pass)
	fmt.Printf("Successfully logged in as %s. Password %s\n", name, passStars.ReplaceAllString(pass, "*"))
	Game.Start(name, pass)
	server.prompt("helloworld> ")
}

func (server *CommandServer) prompt(p string) {
	fmt.Print(p)
	var command string
	fmt.Scanln(&command)
	res := server.handleCommand(command)
	if res != "" {
		fmt.Println(res)
	}
	server.prompt(p)
}

func (server *CommandServer) handleCommand(input string) string {
	fmt.Println(input)
	command, args := parseCommand(input)
	fmt.Println(args)
	if fn, ok := handlerMap[command]; ok {
		res, err := fn(args...)
		if err != nil {
			return err.Error()
		}
		return res
	}
	return fmt.Sprintf("You entered the command: %s", command)
}
