package inmem

import (
	"context"

	"github.com/matthieutran/leafre-login/internal/domain/character"
)

// CharacterRepository implements `character.CharacterRepository` with an in-memory map
type CharacterRepository struct {
	characters map[int]character.Character
}

func NewCharacterRepository() character.CharacterRepository {
	// Create map of character id -> character
	characters := make(map[int]character.Character)
	r := &CharacterRepository{characters: characters}

	return r
}

func (r CharacterRepository) nameExists(name string) bool {
	for _, c := range r.characters {
		if c.Name == name {
			return true
		}
	}

	return false
}

func (r CharacterRepository) Add(ctx context.Context, c character.Character) error {
	if r.nameExists(c.Name) {
		return character.ErrAlreadyExists
	}

	c.ID = len(r.characters)
	r.characters[c.ID] = c
	return nil
}

func (r CharacterRepository) GetByAccountID(ctx context.Context, accountID int) (chars character.Characters, err error) {
	for _, c := range r.characters {
		if c.AccountID == accountID {
			chars = append(chars, c)
		}
	}
	return
}

func (r CharacterRepository) GetByID(ctx context.Context, id int) (c character.Character, err error) {
	c, exists := r.characters[id]
	if !exists {
		err = character.ErrCharDoesNotExist
	}

	return
}

func (r CharacterRepository) GetByName(ctx context.Context, name string) (char character.Character, err error) {
	for _, c := range r.characters {
		if c.Name == name {
			return char, nil
		}
	}

	return char, character.ErrCharDoesNotExist
}

func (r CharacterRepository) Update(ctx context.Context, character character.Character) error {
	return nil
}

func (r CharacterRepository) Destroy(ctx context.Context, id int) (err error) {
	delete(r.characters, id)
	return nil
}
