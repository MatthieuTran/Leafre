package repository

import login "github.com/matthieutran/leafre-login"

type WorldRepository struct {
}

func NewWorldRepository() WorldRepository {
	return WorldRepository{}
}

// Get a list of all active worlds
func (r WorldRepository) FetchAll() (w login.Worlds, err error) {
	var blockCharCreationByte byte = 1

	blockCharCreation := false
	if !blockCharCreation {
		blockCharCreationByte = 0
	}

	return []login.World{
		{
			Id:                0,
			Name:              "Scania",
			State:             2,
			BlockCharCreation: blockCharCreationByte,
		},
	}, nil
}

// Change the specified world's state
func (r WorldRepository) SetWorldState(id int, state byte) {

}
