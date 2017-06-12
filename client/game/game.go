package game

import (
	"time"

	"fmt"

	t "github.com/amhester/go-client-server/types"
)

type PlayerStats struct {
	Attacks     int32
	Successfull int32
	TotalDamage int32
}

type Game struct {
	PC              *Player
	Enemies         []*Enemy
	Stats           *PlayerStats
	BackgroundStuff chan string
}

func NewGame() *Game {
	return &Game{
		Stats: &PlayerStats{},
		Enemies: []*Enemy{
			NewEnemy("1", t.Vector{1, 0, 0}),
			NewEnemy("2", t.Vector{3, 0, 0}),
			NewEnemy("3", t.Vector{1, 4, 0}),
			NewEnemy("4", t.Vector{-2, 6, 0}),
			NewEnemy("5", t.Vector{-1, -2, 0}),
			NewEnemy("6", t.Vector{4, -2, 0}),
			NewEnemy("7", t.Vector{0, 8, 0}),
			NewEnemy("8", t.Vector{-5, -5, 0}),
		},
		BackgroundStuff: make(chan string),
	}
}

func (g *Game) StartEnemies() {
	go func() {
		for {
			select {
			case <-time.After(time.Second * 6):
			inner:
				for _, e := range g.Enemies {
					if e.Position.IsAdjacent(g.PC.Position) {
						suc, dam := e.Attack(g.PC)
						if suc {
							g.BackgroundStuff <- fmt.Sprintf("%s attacked you for %d damage", e.Name, dam)
							continue inner
						}
						g.BackgroundStuff <- fmt.Sprintf("%s attempted to attack you, but failed.", e.Name)
						continue inner
					}
				}
			}
		}
	}()
}

func (g *Game) Start(username, password string) {
	g.PC = NewPlayer(username, password)
}
