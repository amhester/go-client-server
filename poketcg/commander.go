package main

import "errors"
import "fmt"

type Commander struct {
	ls   *LocalStorage
	game *Game
}

func NewCommander(ls *LocalStorage, game *Game) *Commander {
	return &Commander{
		ls:   ls,
		game: game,
	}
}

func (c *Commander) ResolveCommand(args ...string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("No args given")
	}
	key := args[0]
	args = args[1:]
	switch key {
	case "deck":
		return c.deckCommands(args...)
	case "game":
		return c.gameCommands(args...)
	case "help":
		return "The following commands are available:\ndeck\tgame", nil
	default:
		return fmt.Sprintf("Unrecognized command: %s", key), nil
	}
}

func (c *Commander) deckCommands(args ...string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("No args given")
	}
	subCommand := args[0]
	args = args[1:]
	switch subCommand {
	case "add":
		err := c.ls.ImportDeck(args[0])
		if err != nil {
			return "", err
		}
		return "Successfully imported deck.", nil
	case "remove":
	case "list":
	case "describe":
		deck, err := c.ls.GetDeck(args[0])
		if err != nil {
			return "", err
		}
		return deck.String(), nil
	default:
		return fmt.Sprintf("Unrecognized subcommand: %s", subCommand), nil
	}
	return "", nil
}

func (c *Commander) gameCommands(args ...string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("No args given")
	}
	subCommand := args[0]
	args = args[1:]
	switch subCommand {
	case "start":
		return c.game.Start(args...)
	default:
		return fmt.Sprintf("Unrecognized subcommand: %s", subCommand), nil
	}
	return "", nil
}
