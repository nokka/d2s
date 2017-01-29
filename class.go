package d2s

import "fmt"

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

var classMap = map[class]string{
	Amazon:      "Amazon",
	Sorceress:   "Sorceress",
	Necromancer: "Necromancer",
	Paladin:     "Paladin",
	Barbarian:   "Barbarian",
	Druid:       "Druid",
	Assassin:    "Assassin",
}

// class struct will represent the characters in game class.
type class byte

// String will return a class specific name.
func (c class) String() string {
	name, ok := classMap[c]
	if !ok {
		return fmt.Sprintf("Unknown class %#02x", byte(c))
	}
	return name
}
