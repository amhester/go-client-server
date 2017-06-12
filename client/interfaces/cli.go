package interfaces

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLI struct {
	in     *bufio.Reader
	Events chan string
}

func NewCli() *CLI {
	return &CLI{
		in:     bufio.NewReader(os.Stdin),
		Events: make(chan string),
	}
}

func (c *CLI) Start() {
	go func(events chan string, in *bufio.Reader) {
		for {
			event := <-events
			temp := make([]byte, in.Buffered())
			in.Read(temp)
			in.Discard(in.Buffered())
			fmt.Print(event)
			w := bufio.NewWriter(os.Stdin)
			w.WriteString(string(temp))
			w.Flush()
		}
	}(c.Events, c.in)
}

func (c *CLI) Println(event string) {
	c.Events <- event + "\n"
}

func (c *CLI) Print(event string) {
	c.Events <- event
}

func (c *CLI) Readln() string {
	text, _ := c.in.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
