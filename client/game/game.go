package game

type PlayerStats struct {
	Attacks     int32
	Successfull int32
	TotalDamage int32
}

type Game struct {
	PC      *Player
	Enemies []*Enemy
	Stats   *PlayerStats
}

func NewGame() *Game {
	return &Game{
		Stats: &PlayerStats{},
		Enemies: []*Enemy{
			NewEnemy("1"),
			NewEnemy("2"),
			NewEnemy("3"),
			NewEnemy("4"),
			NewEnemy("5"),
			NewEnemy("6"),
			NewEnemy("7"),
			NewEnemy("8"),
		},
	}
}

func (g *Game) Start(username, password string) {
	g.PC = NewPlayer(username, password)
}
