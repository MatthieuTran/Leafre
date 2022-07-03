package login

type World struct {
	// ID of the world
	Id byte `json:"id"`
	// Name of the world
	Name string `json:"name"`
	// State of the world (0: Nothing, 1: Event, 2: New, 3: Hot)
	State byte `json:"state"`
	// Flag to stop characters from being created
	BlockCharCreation byte `json:"block_char_creation"`
}

type Worlds []World

type WorldRepository interface {
	// Get a list of all active worlds
	FetchAll() (Worlds, error)
	// Change the specified world's state
	SetWorldState(id int, state byte)
}
