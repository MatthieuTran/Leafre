package character

import (
	"context"
	"errors"

	"github.com/matthieutran/leafre-login/internal/domain/item"
)

type CharacterForm struct {
	Name      string
	Job       uint32
	SubJob    uint16
	Face      uint32
	Hair      uint32
	HairColor uint32
	Skin      byte
	Coat      uint32
	Pants     uint32
	Shoes     uint32
	Weapon    uint32
	Gender    byte
}

type CharacterService interface {
	CreateCharacter(ctx context.Context, char CharacterForm) (Character, error)
	CheckName(ctx context.Context, name string) (byte, error) // CheckName returns 0 if the name provided has not been taken, 1 if it has been taken
	GetCharacter(ctx context.Context, id uint32) (Character, error)
}

var ErrIncorrectPassword = errors.New("incorrect password")

func NewCharacterService(charRepo CharacterRepository, itemRepo item.ItemRepository) CharacterService {
	return defaultCharacterService{charRepo: charRepo, itemRepo: itemRepo}
}

type defaultCharacterService struct {
	charRepo CharacterRepository
	itemRepo item.ItemRepository
}

const (
	available   byte = 0
	unavailable byte = 1
)

func (s defaultCharacterService) CreateCharacter(ctx context.Context, charDetails CharacterForm) (c Character, err error) {
	// Build Character
	c.Name = charDetails.Name
	c.Job = Job(charDetails.Job)
	c.SubJob = charDetails.SubJob
	c.Face = charDetails.Face
	c.Hair = charDetails.Hair + charDetails.HairColor
	c.Skin = charDetails.Skin
	c.Gender = charDetails.Gender
	c.FieldID = c.Job.StartingField()

	// Add Character
	id, err := s.charRepo.Add(ctx, c)
	if err != nil {
		return
	}

	// Add items
	addEquip := func(templateID uint32, slotID item.BodyPart) {
		it := item.Item{
			CharacterID:   id,
			InventoryType: item.EQUIP,
			SlotID:        0 - int(slotID),
			TemplateID:    templateID,
		}

		s.itemRepo.Add(ctx, it)
	}

	addEquip(charDetails.Coat, item.Clothes)
	addEquip(charDetails.Pants, item.Pants)
	addEquip(charDetails.Shoes, item.Shoes)
	addEquip(charDetails.Weapon, item.Weapon)

	return s.GetCharacter(ctx, id)
}

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

func (s defaultCharacterService) GetCharacter(ctx context.Context, id uint32) (c Character, err error) {
	// Fetch Character
	c, err = s.charRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	// Fetch Character's Items
	charItems, err := s.itemRepo.GetAllByCharacterID(ctx, id)
	if err != nil {
		return
	}

	c.Inventory = make(map[item.InventoryType][]item.Item)

	// Add items to Character inventory map
	for _, i := range charItems {
		c.Inventory[i.InventoryType] = append(c.Inventory[i.InventoryType], i)
	}

	return
}
