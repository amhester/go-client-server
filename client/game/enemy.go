package game

import (
	"math/rand"
)

type Enemy struct {
	Name      string
	Health    int32
	Power     int32
	Speed     int32
	Dexterity int32
	Weapons   map[string]Weapon
}

func NewEnemy(name string) *Enemy {
	return &Enemy{
		Name:      "slime" + name,
		Health:    4,
		Power:     1,
		Speed:     1,
		Dexterity: 1,
		Weapons:   map[string]Weapon{},
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
