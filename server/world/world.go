package world

type IWorld interface {
	Init(seed []byte, state []byte)
	GetBlock()
}

type World struct {
}
