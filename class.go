package d2s

// Mapper of the hex values of classes.
const (
	Amazon      = 0x00
	Sorceress   = 0x01
	Necromancer = 0x02
	Paladin     = 0x03
	Barbarian   = 0x04
	Druid       = 0x05
	Assassin    = 0x06
)

// Class struct will represent the characters in game class.
type Class struct {
	ID byte
}

// Name will return a class specific name based on the ID:
func (c Class) Name() (name string) {
	switch c.ID {
	case Amazon:
		name = "Amazon"
	case Sorceress:
		name = "Sorceress"
	case Necromancer:
		name = "Necromancer"
	case Paladin:
		name = "Paladin"
	case Barbarian:
		name = "Barbarian"
	case Druid:
		name = "Druid"
	case Assassin:
		name = "Assassin"
	}
	return
}
