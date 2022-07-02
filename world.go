package login

type World struct {
	Id                int    // ID of the world
	Name              string // Name of the world
	State             byte   // State of the world (0: Nothing, 1: Event, 2: New, 3: Hot)
	BlockCharCreation bool   // Flag to stop characters from being created
}

type Worlds []World

type WorldRepository interface {
	// Get a list of all active worlds
	FetchAll() (Worlds, error)
	// Change the specified world's state
	SetWorldState(id int, state byte)
}
