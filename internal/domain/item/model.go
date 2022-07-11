package item

type Item struct {
	ID          uint32
	CharacterID uint32

	InventoryType InventoryType
	TemplateID    uint32
	CashItemSN    uint64
	DateExpire    MapleDateTime

	RUC byte
	CUC byte

	Strength     uint16
	Dexterity    uint16
	Intelligence uint16
	Luck         uint16

	MaxHP uint16
	MaxMP uint16

	PhysicalAttackDamage  uint16
	MagicalAttackDamage   uint16
	PhysicalDamageDefense uint16
	MagicalDamageDefense  uint16

	Accuracy uint16
	Evasion  uint16

	Craft uint16
	Speed uint16
	Jump  uint16

	Title     string
	Attribute uint16

	LevelUpType byte
	Level       byte
	Experience  uint32

	Durability uint32
	UIC        uint32

	Grade byte
	CHUC  byte

	Option1 uint16
	Option2 uint16
	Option3 uint16

	Socket1 uint16
	Socket2 uint16
}

type Items []Item

type InventoryType byte

const (
	EQUIP InventoryType = iota + 1
	CONSUME
	INSTALL
	ETC
	CASH
)

func (t InventoryType) String() string {
	switch t {
	case EQUIP:
		return "Equip"
	case CONSUME:
		return "Consume"
	case INSTALL:
		return "Install"
	case ETC:
		return "Etc"
	case CASH:
		return "Cash"
	default:
		return "Unknown"
	}
}
