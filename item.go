package d2s

type magicalProperty struct {
	Bits     int
	Bias     int
	Property string
}

var magicalProperties = map[uint64]magicalProperty{
	0: {Bits: 7, Bias: 32, Property: "Strength"},
	1: {Bits: 7, Bias: 32, Property: "Energy"},
	2: {Bits: 7, Bias: 32, Property: "Dexterity"},
}

// Rarity ID.
const (
	lowQuality        = 0x01
	normal            = 0x02
	highQuality       = 0x03
	magicallyEnhanced = 0x04
	partOfSet         = 0x05
	rare              = 0x06
	unique            = 0x07
	crafted           = 0x08
)

// Item represents a base 111 bit item.
// Item represents an actual item
type Item struct {
	Identified         uint64
	Socketed           uint64
	New                uint64
	IsEar              uint64
	StarterItem        uint64
	SimpleItem         uint64
	Ethereal           uint64
	Personalized       uint64
	GivenRuneword      uint64
	LocationID         uint64
	EquippedID         uint64
	PositionY          uint64
	PositionX          uint64
	AltPositionID      uint64
	Type               string
	NrOfItemsInSockets uint64
	ID                 uint64
	Level              uint64
	Quality            uint64
	MultiplePictures   uint64
	PictureID          uint64
	ClassSpecific      uint64
	LowQualityID       uint64
	StructureHeader    uint64
	DefenseRating      uint64
	MaxDurability      uint64
	CurrentDurability  uint64
	TotalNrOfSockets   uint64
	Quantity           uint64

	// Magical Item properties
	MagicPrefix uint64
	MagicSuffix uint64

	// Set item properties
	SetID        uint64
	SetItemLists uint64

	// All item types >= magicallyEnhanced
	MagicalProperties []magicalProperty
}
