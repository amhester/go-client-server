package game

import (
	"fmt"
	"math/rand"

	t "github.com/amhester/go-client-server/types"
)

type Weapon struct {
	DamageMod float32
}

type Player struct {
	token     string
	Level     int32
	Exp       int64
	NextLevel int64
	Position  t.Vector
	Health    int32
	Power     int32
	Speed     int32
	Dexterity int32
	Weapons   map[string]Weapon
}

func NewPlayer(user, pass string) *Player {
	return &Player{
		token:     user + "_" + pass,
		Level:     1,
		Exp:       0,
		NextLevel: 10,
		Position:  t.Vector{0, 0, 0},
		Health:    10,
		Power:     3,
		Speed:     4,
		Dexterity: 4,
		Weapons: map[string]Weapon{
			"sword":          Weapon{1.5},
			"bat":            Weapon{1.2},
			"spear":          Weapon{1.6},
			"rubber_chicken": Weapon{2.5},
		},
	}
}

func (p *Player) Attack(weapon string, e *Enemy) (bool, int32) {
	probs := rand.Int31n(e.Dexterity * 2)
	if p.Speed >= probs {
		damage := p.Power
		if weapon, ok := p.Weapons[weapon]; ok {
			damage = int32(float32(damage) * weapon.DamageMod)
		}
		e.Health -= damage
		return true, damage
	}
	return false, 0
}

func (p *Player) Move(v t.Vector) {
	p.Position = p.Position.Add(v)
}

func (p *Player) AddExp(e int64) bool {
	p.Exp += e
	if p.Exp >= p.NextLevel {
		p.levelUp()
		return true
	}
	return false
}

func (p *Player) String() string {
	return fmt.Sprintf("\nLevel: %d\nHealth: %d\nPower: %d\nSpeed: %d\nDexterity: %d\n", p.Level, p.Health, p.Power, p.Speed, p.Dexterity)
}

func (p *Player) levelUp() {
	p.Level++
	p.NextLevel += int64(p.Level * p.Level * p.Level)
	p.Health += 2
	p.Power++
	p.Speed++
	p.Dexterity += p.Level % 2
	if p.Exp >= p.NextLevel {
		p.levelUp()
	}
}
