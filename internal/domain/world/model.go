package world

type World struct {
	ID                byte
	Name              string
	State             byte
	EventDesc         string
	EventEXP          uint16
	EventDrop         uint16
	BlockCharCreation byte
	Balloon           uint16
}

type Worlds []World
