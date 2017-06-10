package commander

import (
	"fmt"

	"github.com/amhester/go-client-server/client/game"
)

var handlerMap = map[string]func(args ...string) (string, error){
	"hello":  world,
	"attack": attack,
}

func world(args ...string) (string, error) {
	return "world!", nil
}

func attack(args ...string) (string, error) {
	enemy := args[0]
	var weapon string
	if len(args) == 2 {
		weapon = args[1]
	}
	for _, e := range Game.Enemies {
		if e.Name == enemy {
			suc, dam := Game.PC.Attack(weapon, e)
			if suc {
				if e.Health <= 0 {
					newEnemies := []*game.Enemy{}
					for _, e2 := range Game.Enemies {
						if e2.Name != enemy {
							newEnemies = append(newEnemies, e2)
						}
					}
					if lu := Game.PC.AddExp(5); lu {
						return fmt.Sprintf("You defeated %s, you also leveled up, good for you!%s", enemy, Game.PC), nil
					}
					return fmt.Sprintf("You defated %s.", enemy), nil
				}
				return fmt.Sprintf("You successfully attacked %s for %d damage.", enemy, dam), nil
			}
			return "Your attack missed.", nil
		}
	}
	fmt.Println("ENEMIES:", Game.Enemies)
	return fmt.Sprintf("No enemy named %s.", enemy), nil
}
