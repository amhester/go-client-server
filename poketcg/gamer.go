package main

import (
	"fmt"

	"strings"

	"strconv"

	cli "github.com/amhester/go-client-server/client/interfaces"
)

type PlayerStats struct {
}

type Player struct {
	Name    string
	Stats   *PlayerStats
	Deck    *Deck
	Hand    *Deck
	Prizes  *Deck
	Discard *Deck
	Bench   *Deck
	Active  *Card
	Stadium *Card
}

type Game struct {
	ls      *LocalStorage
	cli     *cli.CLI
	Players []*Player
	Turn    int
}

func NewGame(ls *LocalStorage) *Game {
	return &Game{
		ls:  ls,
		cli: cli.NewCli(),
	}
}

func (g *Game) Start(args ...string) (string, error) {
	players := []*Player{}
	for i := 0; i < len(args)-1; i += 2 {
		name := args[i]
		if i+1 >= len(args) {
			return fmt.Sprintf("%s needs a deck", name), nil
		}
		deck, err := g.ls.GetDeck(args[i+1])
		if err != nil {
			return "", err
		}
		players = append(players, &Player{
			Name:    name,
			Deck:    deck,
			Hand:    &Deck{Id: "", Cards: []*Card{}, Name: "Hand"},
			Prizes:  &Deck{Id: "", Cards: []*Card{}, Name: "Prizes"},
			Discard: &Deck{Id: "", Cards: []*Card{}, Name: "Discard"},
			Bench:   &Deck{Id: "", Cards: []*Card{}, Name: "Bench"},
		})
	}
	if len(players) < 1 || len(players) > 2 {
		return "Too few or too many players", nil
	}
	g.Players = players
	g.Turn = 0
	done := make(chan error)
	g.start(done)
	err := <-done
	close(done)
	return "", err
}

func (g *Game) start(done chan error) {
	fmt.Println("Starting game...")
	for _, p := range g.Players {
		p.Deck.Shuffle()
		for i := 0; i < 7; i++ {
			p.Deck.Move(0, p.Hand)
		}
		for i := 0; i < 6; i++ {
			p.Deck.Move(0, p.Prizes)
		}
	}
	go func(done_ chan error) {
		for {
			d, err := g.turn()
			if err != nil {
				done_ <- err
				return
			}
			if d {
				done_ <- nil
				return
			}
		}
	}(done)
}

func (g *Game) turn() (bool, error) {
	g.Turn++
	pIdx := (g.Turn - 1) % len(g.Players)
	p := g.Players[pIdx]
	if len(p.Deck.Cards) == 0 {
		return true, nil
	}
	fmt.Println(fmt.Sprintf("Turn %d, %s's turn.", g.Turn, p.Name))
	var done bool
	for !done {
		commandText := g.cli.Readln()
		args := strings.Split(commandText, " ")
		if len(args) == 0 {
			fmt.Println("Too few arguments")
			continue
		}
		command := args[0]
		args = args[1:]
		switch command {
		case "show":
			if len(args) == 0 {
				fmt.Println("Not enough arguments.")
				continue
			}
			if len(args) == 1 {
				switch args[0] {
				case "deck":
					fmt.Println(p.Deck.List())
				case "prizes":
					fmt.Println(p.Prizes.List())
				case "discard":
					fmt.Println(p.Discard.List())
				case "hand":
					fmt.Println(p.Hand.List())
				case "bench":
					fmt.Println(p.Bench.List())
				}
				continue
			}
			if len(args) == 2 {
				var deck_ *Deck
				switch args[0] {
				case "deck":
					deck_ = p.Deck
				case "prizes":
					deck_ = p.Prizes
				case "discard":
					deck_ = p.Discard
				case "hand":
					deck_ = p.Hand
				case "bench":
					deck_ = p.Bench
				}
				which, err := strconv.ParseInt(args[1], 10, 64)
				if err != nil {
					fmt.Println("Invalid card index.")
					continue
				}
				if len(deck_.Cards) <= int(which) {
					fmt.Println("Card out of range.")
					continue
				}
				fmt.Println(deck_.Cards[int(which)].LongString())
			}
		case "move":
			swapDecks := make([]*Deck, 2)
			deckStrs := []string{
				args[0],
				args[2],
			}
			which, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			for idx, d := range deckStrs {
				switch d {
				case "deck":
					swapDecks[idx] = p.Deck
				case "prizes":
					swapDecks[idx] = p.Prizes
				case "discard":
					swapDecks[idx] = p.Discard
				case "hand":
					swapDecks[idx] = p.Hand
				case "bench":
					swapDecks[idx] = p.Bench
				}
			}
			if swapDecks[0] == nil || swapDecks[1] == nil {
				fmt.Println("Unknown pile given.")
				continue
			}
			swapDecks[0].Move(int(which), swapDecks[1])
		case "attach":
			source, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			target, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			if len(p.Hand.Cards) <= int(source) {
				fmt.Println("Card out of range")
				continue
			}
			if target == -1 {
				p.Hand.Move(int(source), p.Active.Attachments)
				continue
			}
			if len(p.Bench.Cards) <= int(target) {
				fmt.Println("Card out of range")
				continue
			}
			p.Hand.Move(int(source), p.Bench.Cards[int(target)].Attachments)
		case "done":
			done = true
		case "help":
			fmt.Println("Here is a list of commands that can be executed:")
			fmt.Println(``)
		default:
			fmt.Println("Unknown Command")
		}
	}
	for _, p_ := range g.Players {
		if len(p_.Prizes.Cards) == 0 {
			return true, nil
		}
	}
	return false, nil
}
