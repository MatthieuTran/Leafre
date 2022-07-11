package character

import "github.com/matthieutran/leafre-login/internal/domain/item"

type Character struct {
	ID        uint32
	AccountID uint32
	Name      string

	Gender byte
	Skin   byte
	Face   uint32
	Hair   uint32

	Level        byte
	Job          Job
	Strength     uint16
	Dexterity    uint16
	Intelligence uint16
	Luck         uint16

	HP    uint32
	MaxHP uint32
	MP    uint32
	MaxMP uint32

	AP       uint16
	SP       uint16
	ExtendSP ExtendSP

	Experience uint32
	Popularity uint16

	TempExperience uint32

	FieldID     uint32
	FieldPortal byte

	PlayTime uint32

	SubJob uint16

	Inventory map[item.InventoryType][]item.Item // Inventory is a map of inventory type to a slice of items
}

type Characters []Character

// ExtendSP is defined as a map with the jobLevel as the key and skill point as the value
type ExtendSP map[byte]byte
