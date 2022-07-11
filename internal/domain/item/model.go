package item

type Item struct {
	ID          uint32
	CharacterID uint32

	InventoryType InventoryType
	SlotID        int
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

type BodyPart uint16

const (
	Hair BodyPart = iota
	Cap
	FaceAcc
	EyeAcc
	EarAcc
	Clothes
	Pants
	Shoes
	Gloves
	Cape
	Shield
	Weapon
	Ring1
	Ring2
	PetWear
	Ring3
	Ring4
	Pendant
	TamingMob
	Saddle
	MobEquip
	PetRingLabel
	PetAbilItem
	PetAbilMeso
	PetAbilHpConsume
	PetAbilMechanicConsume
	PetAbilSweepForDrop
	PetAbilLongRange
	PetAbilPickupOthers
	PetRingQuote
	PetWear2
	PetRingLabel2
	PetRingQuote2
	PetAbilItem2
	PetAbilMeso2
	PetAbilSweepForDrop2
	PetAbilLongRange2
	PetAbilPickupOthers2
	PetWear3
	PetRingLabel3
	PetRingQuote3
	PetAbilItem3
	PetAbilMeso3
	PetAbilSweepForDrop3
	PetAbilLongRange3
	PetAbilPickupOthers3
	PetAbilIgnoreItems1
	PetAbilIgnoreItems2
	PetAbilIgnoreItems3
	Medal
	Belt
	Shoulder

	Nothing3 BodyPart = iota + 2
	Nothing2
	Nothing1
	Nothing0

	Ext0 BodyPart = iota + 2
	ExtPendant1
	Ext1
	Ext2
	Ext3
	Ext4
	Ext5
	Ext6

	Sticker BodyPart = 100

	DragonCap     BodyPart = 1000
	DragonPendant BodyPart = 1001
	DragonWing    BodyPart = 1002
	DragonShoes   BodyPart = 1003

	MechanicEngine     BodyPart = 1100
	MechanicArm        BodyPart = 1101
	MechanicLeg        BodyPart = 1102
	MechanicFrame      BodyPart = 1103
	MechanicTransistor BodyPart = 1104
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
