package d2s

// Item represents a base 111 bit item.
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
}

// MagicalItem specific fields.
type MagicalItem struct {
	Item
	MagicPrefix uint64
	MagicSuffix uint64
}
