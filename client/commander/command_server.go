package commander

import (
	"fmt"
	"regexp"
	"time"

	"github.com/amhester/go-client-server/client/game"
	cli "github.com/amhester/go-client-server/client/interfaces"
)

var passStars = regexp.MustCompile(".")
var Game *game.Game

type CommandServer struct {
	Exit chan error
	Face *cli.CLI
}

func NewCommandServer(g *game.Game) *CommandServer {
	Game = g
	c := cli.NewCli()
	c.Start()
	return &CommandServer{
		Exit: make(chan error),
		Face: c,
	}
}

func (server *CommandServer) Start() chan error {
	go server.init()
	return server.Exit
}

func (server *CommandServer) init() {
	var name string
	var pass string
	server.Face.Println("Starting server...")
	time.Sleep(time.Second)
	server.Face.Println("Enter username: ")
	name = server.Face.Readln()
	server.Face.Println("Enter pasword: ")
	pass = server.Face.Readln()
	server.Face.Println(fmt.Sprintf("Successfully logged in as %s. Password %s", name, passStars.ReplaceAllString(pass, "*")))
	Game.Start(name, pass)
	Game.StartEnemies()
	go func() {
		for {
			select {
			case stuff := <-Game.BackgroundStuff:
				server.Face.Print("\n" + stuff + "\nhelloworld> ")
			}
		}
	}()
	server.prompt("helloworld> ")
}

func (server *CommandServer) prompt(p string) {
	server.Face.Print(p)
	command := server.Face.Readln()
	res := server.handleCommand(command)
	if res != "" {
		server.Face.Println(res)
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
