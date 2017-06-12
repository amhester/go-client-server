package commander

import (
	"bufio"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func parseCommand(command string) (string, []string) {
	parts := strings.Split(command, " ")
	return parts[0], parts[1:]
}
