package main

var (
	forever = make(chan struct{})
)

func main() {

	<-forever
}
