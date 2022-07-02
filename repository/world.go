package repository

import login "github.com/matthieutran/leafre-login"

type WorldRepository struct {
}

// Get a list of all active worlds
func FetchAll() (w login.Worlds, err error) {
	return []login.World{
		{
			Id:                0,
			Name:              "Scania",
			State:             2,
			BlockCharCreation: false,
		},
	}, nil
}

// Change the specified world's state
func SetWorldState(id int, state byte)
