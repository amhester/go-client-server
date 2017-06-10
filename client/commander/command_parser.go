package commander

import (
	"strings"
)

func parseCommand(command string) (string, []string) {
	parts := strings.Split(command, ">")
	return parts[0], parts[1:]
}
