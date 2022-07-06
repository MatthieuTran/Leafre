package world

type WorldService interface {
	// Get a list of all active worlds
	FetchAll() (Worlds, error)
	// Change the specified world's state
	SetWorldState(id int, state byte)
}

// worldService provides a rough implementation of a WorldService using hard code-coded data
type worldService struct {
}

func NewWorldService() WorldService {
	return worldService{}
}

// Get a list of all active worlds
func (s worldService) FetchAll() (w Worlds, err error) {
	var blockCharCreationByte byte = 1

	blockCharCreation := false
	if !blockCharCreation {
		blockCharCreationByte = 0
	}

	return Worlds{
		{
			Id:                0,
			Name:              "Scania",
			State:             2,
			BlockCharCreation: blockCharCreationByte,
		},
	}, nil
}

// Change the specified world's state
func (s worldService) SetWorldState(id int, state byte) {

}
