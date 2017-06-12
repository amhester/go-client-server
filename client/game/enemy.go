package game

import (
	"math/rand"

	t "github.com/amhester/go-client-server/types"
)

type Enemy struct {
	Name      string
	Health    int32
	Power     int32
	Speed     int32
	Dexterity int32
	Weapons   map[string]Weapon
	Position  t.Vector
}

func NewEnemy(name string, startPos t.Vector) *Enemy {
	return &Enemy{
		Name:      "slime" + name,
		Health:    4,
		Power:     1,
		Speed:     1,
		Dexterity: 1,
		Weapons:   map[string]Weapon{},
		Position:  startPos,
	}
}

func (e *Enemy) Attack(p *Player) (bool, int32) {
	probs := rand.Int31n(p.Dexterity * 2)
	if e.Speed >= probs {
		p.Health -= e.Power
		return true, e.Power
	}
	return false, 0
}

func (e *Enemy) String() string {
	return e.Name
}

func (e *Enemy) Move(v t.Vector) {
	e.Position = e.Position.Add(v)
}
