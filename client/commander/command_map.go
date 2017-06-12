package commander

import (
	"errors"
	"fmt"

	"strconv"

	"github.com/amhester/go-client-server/client/game"
	t "github.com/amhester/go-client-server/types"
)

var handlerMap = map[string]func(args ...string) (string, error){
	"hello":  world,
	"attack": attack,
	"move":   move,
	"look":   look,
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
			if !Game.PC.Position.IsAdjacent(e.Position) {
				return fmt.Sprintf("%s is not in range of your attack.", e.Name), nil
			}
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

func move(args ...string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("Invalid number of arguments passed to move.")
	}
	moveSpaces := Game.PC.Speed
	if len(args) == 2 {
		override, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			return "", err
		}
		if int32(override) < moveSpaces && override > 0 {
			moveSpaces = int32(override)
		}
	}
	switch args[0] {
	case "up":
		Game.PC.Move(t.Vector{0, moveSpaces, 0})
	case "down":
		Game.PC.Move(t.Vector{0, -moveSpaces, 0})
	case "left":
		Game.PC.Move(t.Vector{-moveSpaces, 0, 0})
	case "right":
		Game.PC.Move(t.Vector{moveSpaces, 0, 0})
	default:
		return "Are you confused, that's not a direction!", nil
	}
	return fmt.Sprintf("Moved %d spaces %s", moveSpaces, args[0]), nil
}

func look(args ...string) (string, error) {
	ret := ""
	for _, e := range Game.Enemies {
		ret = fmt.Sprintf("%s\n%s: %v away.", ret, e.Name, Game.PC.Position.Distance(e.Position))
	}
	return ret, nil
}
