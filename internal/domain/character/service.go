package character

import (
	"context"
	"errors"
)

type CharacterService interface {
	CheckName(ctx context.Context, name string) (byte, error) // CheckName returns 0 if the name provided has not been taken, 1 if it has been taken
}

var ErrIncorrectPassword = errors.New("incorrect password")

func NewCharacterService(charRepo CharacterRepository) CharacterService {
	return defaultCharacterService{charRepo: charRepo}
}

type defaultCharacterService struct {
	charRepo CharacterRepository
}

const (
	available   byte = 0
	unavailable byte = 1
)

func (s defaultCharacterService) CheckName(ctx context.Context, name string) (byte, error) {
	_, err := s.charRepo.GetByName(ctx, name)
	if err != nil {
		if errors.Is(err, ErrCharDoesNotExist) {
			return available, nil
		}

		return unavailable, err
	}

	return unavailable, nil
}
