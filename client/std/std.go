package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f := os.Stdout
	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	fmt.Println("FROMIN:", text)
	fs, _ := f.Stat()
	f.Truncate(fs.Size() - 6)
	os.Stdin.Truncate(4)
}
