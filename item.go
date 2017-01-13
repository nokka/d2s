package d2s

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
	MagicAttributes []magicAttribute
}

// Rarity IDs.
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

// Item type IDs.
const (
	Ring         = "rin"
	CompositeBow = "cbw"
)

type magicAttribute struct {
	ID     uint64
	Name   string
	Values []uint64
}

type magicalProperty struct {
	Bits []uint
	Bias uint64
	Name string
}

var magicalProperties = map[uint64]magicalProperty{
	0:  {Bits: []uint{7}, Bias: 32, Name: "+X to Strength"},
	1:  {Bits: []uint{7}, Bias: 32, Name: "+X to Energy"},
	2:  {Bits: []uint{7}, Bias: 32, Name: "+X to Dexterity"},
	3:  {Bits: []uint{7}, Bias: 32, Name: "+X to Vitality"},
	7:  {Bits: []uint{8}, Bias: 32, Name: "+X to Life"},
	9:  {Bits: []uint{8}, Bias: 32, Name: "+X to Mana"},
	11: {Bits: []uint{8}, Bias: 32, Name: "+X to Maximum Stamina"},
	16: {Bits: []uint{9}, Name: "+X% Enhanced Defense"},
	17: {Bits: []uint{9, 9}, Name: "+X% Enhanced Damage"},
	19: {Bits: []uint{10}, Name: "+X to Attack rating"},
	20: {Bits: []uint{6}, Name: "+X% Increased chance of blocking"},
	21: {Bits: []uint{6}, Name: "+X to Minimum damage"},
	22: {Bits: []uint{7}, Name: "+X to Maximum damage"},
	// TODO: Possible duplicate of id: 21, usually seen together
	23: {Bits: []uint{6}, Name: "+X to Minimum damage"},
	// TODO: Possible duplicate of id: 22, usually seen together
	24: {Bits: []uint{7}, Name: "+X to Maximum damage"},
	27: {Bits: []uint{8}, Name: "Regenerate Mana X%"},
	28: {Bits: []uint{8}, Name: "Heal Stamina X%"},
	31: {Bits: []uint{10}, Bias: 10, Name: "+X Defense"},
	32: {Bits: []uint{10}, Bias: 10, Name: "+X vs. Missile"},
	33: {Bits: []uint{10}, Bias: 10, Name: "+X vs. Melee"},
	34: {Bits: []uint{6}, Name: "Damage Reduced by X"},
}

// All item types that contain the defense rating bits will exist in here,
// we'll use this when reading items to make sure we only read defense rating bits
// when they exist, or we'll ruin the rest of the bit offsets for the item.
var defenseRatingMap = map[string]bool{
	Ring:         false,
	CompositeBow: false,
}

// All item types that contain the durability bits will exist in here,
// we'll use this when reading items to make sure we only read durability bits
// when they exist, or we'll ruin the rest of the bit offsets for the item.
var durabilityMap = map[string]bool{
	Ring:         false,
	CompositeBow: true,
}

// All item types that contain the quantity bits will exist in here,
// we'll use this when reading items to make sure we only read quantity bits
// when they exist, or we'll ruin the rest of the bit offsets for the item.
var quantityMap = map[string]bool{
	Ring:         false,
	CompositeBow: false,
}
