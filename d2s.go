package d2s

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Parse will read the data from a d2s character file and return a normalized struct.
func Parse(file io.Reader) (*Character, error) {
	// Implements buffered reading, wraps io.Reader.
	bfr := bufio.NewReader(file)

	return parse(bfr)
}

// ParseFromContent will read the character from a byte slice.
func ParseFromContent(data []byte) (*Character, error) {
	// Create a reader from the byte slice.
	r := bytes.NewReader(data)

	// Implements buffered reading, wraps io.Reader.
	bfr := bufio.NewReader(r)

	return parse(bfr)
}

// parse will do the actual reading from an io.Reader.
func parse(bfr *bufio.Reader) (*Character, error) {
	// Create character pointer, we'll pass it around.
	char := &Character{}

	if err := parseHeader(bfr, char); err != nil {
		return nil, fmt.Errorf("failed to parse header: %s, %s", char.Header.Name, err)
	}

	if err := parseAttributes(bfr, char); err != nil {
		return nil, fmt.Errorf("failed to parse attributes: %s, %s", char.Header.Name, err)
	}

	if err := parseSkills(bfr, char); err != nil {
		return nil, fmt.Errorf("failed to parse skills: %s, %s", char.Header.Name, err)
	}

	if err := parseItems(bfr, char); err != nil {
		return nil, fmt.Errorf("failed to parse items: %s, %s", char.Header.Name, err)
	}

	if err := parseCorpse(bfr, char); err != nil {
		return nil, fmt.Errorf("failed to parse corpse: %s, %s", char.Header.Name, err)
	}

	// Normalize the character status, that is being stored as a byte.
	status := char.Header.Status.Readable()

	if status.Expansion {
		if err := parseMercItems(bfr, char); err != nil {
			return nil, fmt.Errorf("failed to parse merc items: %s, %s", char.Header.Name, err)
		}
	}

	if char.Header.Class == Necromancer && status.Expansion {
		if err := parseIronGolem(bfr, char); err != nil {
			return nil, fmt.Errorf("failed to parse iron golem: %s, %s", char.Header.Name, err)
		}
	}

	return char, nil
}

// Parses the 767 byte static header, all characters have this data.
func parseHeader(bfr io.Reader, char *Character) error {
	// Make a buffer that can hold 767 bytes, which can hold the entire header.
	// The header is 765 bytes, but the first 2 bytes of the attributes section
	// is included here, because it's a 2 bytes header that we're not really
	// interested in.
	buf := make([]byte, 767)

	_, err := io.ReadFull(bfr, buf)
	if err != nil {
		return err
	}

	err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &char.Header)
	if err != nil {
		return err
	}

	return nil
}

// Parses the total attributes the character has, before items are calculated.
func parseAttributes(bfr io.ByteReader, char *Character) error {
	// Create a bit reader from the buffered reader to read the stats.
	br := bitReader{r: bfr}

	for {
		id, err := br.ReadBits(9)
		if err != nil {
			return err
		}

		// If all 9 bits are set, we've hit the end of the attributes section
		//  at 0x1ff and exit the loop.
		if id == 0x1ff {
			break
		}

		// The attribute value bit length, so we'll know how many bits to read next.
		length, ok := attributeBitMap[id]
		if !ok {
			return fmt.Errorf("unknown attribute id: %d", id)
		}

		// The attribute value.
		attr, err := br.ReadBits(length)
		if err != nil {
			return err
		}

		switch id {
		case strength:
			char.Attributes.Strength = attr
		case energy:
			char.Attributes.Energy = attr
		case dexterity:
			char.Attributes.Dexterity = attr
		case vitality:
			char.Attributes.Vitality = attr
		case unusedStats:
			char.Attributes.UnusedStats = attr
		case unusedSkills:
			char.Attributes.UnusedSkillPoints = attr
		case currentHP:
			char.Attributes.CurrentHP = attr / 256
		case maxHP:
			char.Attributes.MaxHP = attr / 256
		case currentMana:
			char.Attributes.CurrentMana = attr / 256
		case maxMana:
			char.Attributes.MaxMana = attr / 256
		case currentStamina:
			char.Attributes.CurrentStamina = attr / 256
		case maxStamina:
			char.Attributes.MaxStamina = attr / 256
		case level:
			char.Attributes.Level = attr
		case experience:
			char.Attributes.Experience = attr
		case gold:
			char.Attributes.Gold = attr
		case stashedGold:
			char.Attributes.StashedGold = attr
		}
	}

	return nil
}

// Parses the skill set, this contains of a 2 byte header and a 30 byte list of
// 30 indexes with the allocated points for that skill. The class determines
// what offset we start reading from in the skill list.
func parseSkills(bfr io.Reader, char *Character) error {
	// Make a buffer that can hold 32 bytes, which can hold the entire skillset.
	buf := make([]byte, 32)

	_, err := io.ReadFull(bfr, buf)
	if err != nil {
		return err
	}

	skillHeaderData := skillData{}
	if err := binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &skillHeaderData); err != nil {
		return err
	}

	if string(skillHeaderData.Header[:]) != "if" {
		return fmt.Errorf("failed to find skill header")
	}

	skillOffset, ok := skillOffsetMap[uint(char.Header.Class)]
	if !ok {
		return fmt.Errorf("unknown skill offset for class %d", char.Header.Class)
	}

	for i, allocatedPoints := range skillHeaderData.List {
		id := (i + skillOffset)

		skillName, ok := skillMap[id]
		if !ok {
			return fmt.Errorf("unknown skill id %d", id)
		}
		s := Skill{
			ID:     id,
			Points: int(allocatedPoints),
			Name:   skillName,
		}
		char.Skills = append(char.Skills, s)
	}

	return nil
}

// Parses all items on the character, this includes items in stash, cube,
// inventory and equipped items.
func parseItems(bfr io.ByteReader, char *Character) error {
	// Make a buffer that can hold 4 bytes, which can hold the items header.
	buf := make([]byte, 4)

	_, err := io.ReadFull(bfr.(io.Reader), buf)
	if err != nil {
		return err
	}

	itemHeaderData := itemData{}
	if err := binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &itemHeaderData); err != nil {
		return err
	}

	if string(itemHeaderData.Header[:]) != "JM" {
		return errors.New("failed to find the items header")
	}

	char.Items, err = ParseItemList(bfr, int(itemHeaderData.Count))
	if err != nil {
		return err
	}

	return nil
}

// Parses corpse data, if it exists, otherwise just reads the item header.
// If the item count of the header is 1, this means the character is dead and
// will have a corpse at its feet when loading into the game. The 12 bytes
// following is the corpse data, which we don't really care about.
func parseCorpse(bfr io.ByteReader, char *Character) error {
	// Make a buffer that can hold 16 bytes, which can hold the corpse data.
	buf := make([]byte, 16)

	// Read the first 4 bytes that contain the header string and if the char
	// is dead or not.
	_, err := io.ReadFull(bfr.(io.Reader), buf[:4])
	if err != nil {
		return err
	}

	corpseHeaderData := corpseData{}
	err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &corpseHeaderData)

	if err != nil {
		return err
	}

	if string(corpseHeaderData.Header[:]) != "JM" {
		return errors.New("failed to find the corpse items header")
	}

	// The character is currently dead and will have the corpse item list here.
	if corpseHeaderData.Count == 1 {
		// Character is dead, so we'll save the state.
		char.IsDead = 1

		// 12 Unknown bytes.
		_, err := io.ReadFull(bfr.(io.Reader), buf[:12])
		if err != nil {
			return err
		}

		_, err = io.ReadFull(bfr.(io.Reader), buf[:4])
		if err != nil {
			return err
		}

		itemHeaderData := itemData{}
		err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &itemHeaderData)

		if err != nil {
			return err
		}

		if string(itemHeaderData.Header[:]) != "JM" {
			return errors.New("failed to find the merc items header")
		}

		corpseItems, err := ParseItemList(bfr, int(itemHeaderData.Count))
		if err != nil {
			return err
		}

		char.CorpseItems = corpseItems
	}

	return nil
}

// Parses all items on the merc, if it exists, otherwise just reads the header.
func parseMercItems(bfr io.ByteReader, char *Character) error {
	br := bitReader{r: bfr}

	// offset: 0 "j"
	j, err := br.ReadByte()
	if err != nil {
		return err
	}

	// offset: 8, "f"
	f, err := br.ReadByte()
	if err != nil {
		return err
	}

	if string(j) != "j" || string(f) != "f" {
		return errors.New("failed to find merc header jf")
	}

	// If you have a merc, we'll read the item list of the merc here.
	if char.Header.MercID != 0 {
		// Make a buffer that can hold 4 bytes, which can hold the items header.
		buf := make([]byte, 4)

		_, err := io.ReadFull(bfr.(io.Reader), buf[:4])
		if err != nil {
			return err
		}

		itemHeaderData := itemData{}
		err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &itemHeaderData)

		if err != nil {
			return err
		}

		if string(itemHeaderData.Header[:]) != "JM" {
			return errors.New("failed to find the merc items header")
		}

		items, err := ParseItemList(bfr, int(itemHeaderData.Count))
		if err != nil {
			return err
		}

		char.MercItems = items
	}

	return nil
}

// If the character has a golem, we'll read the item list for the golem,
// since an iron golem is made from an item, the properties of that item is stored.
func parseIronGolem(bfr io.ByteReader, char *Character) error {
	// Make a buffer that can hold 3 bytes, which can hold the items header.
	buf := make([]byte, 3)

	_, err := io.ReadFull(bfr.(io.Reader), buf[:3])
	if err != nil {
		return err
	}

	golemHeaderData := golemData{}
	err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &golemHeaderData)

	if err != nil {
		return err
	}

	if string(golemHeaderData.Header[:]) != "kf" {
		return errors.New("failed to find the golem header")
	}

	if golemHeaderData.HasGolem == 1 {
		item, err := ParseItemList(bfr, 1)
		if err != nil {
			return err
		}

		char.GolemItem = &item[0]
	}

	return nil
}

func ParseItemList(byteReader io.ByteReader, itemCount int) ([]Item, error) {
	var itemList []Item
	br := bitReader{r: byteReader}

	// We'll start this number at items count, but the thing is, if an item
	// has items socketed on it, they will follow the item they're socketed in,
	// so every time we find an item with socketed items, we'll increment this
	// list in order to read them as well.
	numberOfItemsToRead := itemCount

	for i := 0; i < numberOfItemsToRead; i++ {
		// Reset the bit reader when reading a new item.
		br.Reset()

		item := Item{}
		// Read the 111 bit basic item structure, all items have this structure.
		err := parseSimpleBits(&br, &item)
		if err != nil {
			return itemList, err
		}

		if item.SimpleItem == 0 {
			// offset 111, item id is 8 chars, each 4 bit
			if item.ID, err = br.ReadBits(32); err != nil {
				return itemList, err
			}

			// offset 143
			if item.Level, err = br.ReadBits(7); err != nil {
				return itemList, err
			}

			// offset 150
			if item.Quality, err = br.ReadBits(4); err != nil {
				return itemList, err
			}

			// If true, it means the item has more than one picture associated with it.
			if item.MultiplePictures, err = br.ReadBits(1); err != nil {
				return itemList, err
			}

			if item.MultiplePictures == 1 {
				// The next 3 bits contain the picture ID.
				if item.PictureID, err = br.ReadBits(3); err != nil {
					return itemList, err
				}
			}

			// If this is true, it means the item is class specific.
			if item.ClassSpecific, err = br.ReadBits(1); err != nil {
				return itemList, err
			}

			// If the item is class specific, the next 11 bits will
			// contain the class specific data.
			if item.ClassSpecific == 1 {
				// TODO: Parse this into something useful
				br.ReadBits(11)
			}

			switch item.Quality {
			case lowQuality:
				if item.LowQualityID, err = br.ReadBits(3); err != nil {
					return itemList, err
				}

			case normal:
			// No extra data present
			case highQuality:
				// TODO: Figure out what these 3 bits are on a high quality item
				br.ReadBits(3)
			case magicallyEnhanced:
				if item.MagicPrefix, err = br.ReadBits(11); err != nil {
					return itemList, err
				}

				if prefixName, ok := magicalPrefixes[item.MagicPrefix]; ok {
					item.MagicPrefixName = prefixName
				}

				if item.MagicSuffix, err = br.ReadBits(11); err != nil {
					return itemList, err
				}

				if suffixName, ok := magicalSuffixes[item.MagicSuffix]; ok {
					item.MagicSuffixName = suffixName
				}
			case partOfSet:
				if item.SetID, err = br.ReadBits(12); err != nil {
					return itemList, err
				}

				if setName, ok := setNames[item.SetID]; ok {
					item.SetName = setName
				}
			case rare, crafted:
				if err := parseRareOrCraftedBits(&br, &item); err != nil {
					return itemList, err
				}

			case unique:
				if item.UniqueID, err = br.ReadBits(12); err != nil {
					return itemList, err
				}

				if uniqueName, ok := uniqueNames[item.UniqueID]; ok {
					item.UniqueName = uniqueName
				}
			}

			// MARK: Runeword data
			if item.GivenRuneword == 1 {
				if err := parseRunewordBits(&br, &item); err != nil {
					return itemList, err
				}
			}

			if item.Personalized == 1 {
				var name string
				for {
					c, err := br.ReadBits(7)
					if err != nil {
						return itemList, err
					}

					if c == 0 {
						break
					}

					name += fmt.Sprintf("%c", c)
				}

				item.PersonalizedName = name
			}

			// If the item is a tome, it will contain 5 extra bits, we're not
			// interested in these bits, the value is usually 1, but not sure
			// what it is.
			if tomeMap[item.Type] {
				br.ReadBits(5)
			}

			// All items have this field between the personalization (if it exists)
			// and the item specific data
			if item.Timestamp, err = br.ReadBits(1); err != nil {
				return itemList, err
			}

			if item.TypeID == armor || item.TypeID == shield {
				// If the item is an armor, it will have this field of defense data.
				defRating, err := br.ReadBits(11)
				if err != nil {
					return itemList, err
				}

				// We need to subtract 10 defense rating from all armors for
				// some reason, I'm not sure why.
				item.DefenseRating = int64((defRating - 10))
			}

			if item.TypeID == armor || item.TypeID == weapon || item.TypeID == shield {
				if item.MaxDurability, err = br.ReadBits(8); err != nil {
					return itemList, err
				}

				// Some weapons like phase blades don't have durability, so we'll
				// check if the item has max durability, then we can safely assume
				// it has current durability too.
				if item.MaxDurability > 0 {
					if item.CurrentDurability, err = br.ReadBits(8); err != nil {
						return itemList, err
					}

					// Seems to be a random bit here.
					br.ReadBits(1)
				}
			}

			if quantityMap[item.Type] {
				// If the item is a stacked item, e.g. a javelin or something, these 9
				// bits will contain the quantity.
				if item.Quantity, err = br.ReadBits(9); err != nil {
					return itemList, err
				}
			}

			// If the item is socketed, it will contain 4 bits of data which are the nr
			// of total sockets the item have, regardless of how many are occupied by
			// an item.
			if item.Socketed == 1 {
				if item.TotalNrOfSockets, err = br.ReadBits(4); err != nil {
					return itemList, err
				}
			}

			// If the item is part of a set, these bit will tell us how many lists
			// of magical properties follow the one regular magical property list.
			var setListValue uint64 = 0
			if item.Quality == partOfSet {
				setListValue, err = br.ReadBits(5)
				if err != nil {
					return itemList, err
				}

				listCount, ok := setListMap[setListValue]
				if !ok {
					return itemList, fmt.Errorf("unknown set list number %d", setListValue)
				}

				item.SetListCount = listCount
			}

			// MARK: Time to parse 9 bit magical property ids followed by their n bit
			// length values, but only if the item is magical or above.
			magicAttrList, err := parseMagicalList(&br)
			if err != nil {
				return itemList, err
			}

			item.MagicAttributes = magicAttrList

			// Item has more magical property lists due to being a set item.
			if item.SetListCount > 0 {
				for i := 0; i < int(item.SetListCount); i++ {
					setAttrList, err := parseMagicalList(&br)
					if err != nil {
						return itemList, err
					}

					item.SetAttributes = append(item.SetAttributes, setAttrList)
				}
				reqSetIDs, ok := setReqIDsMap[item.SetID]
				if ok {
					item.SetAttributesIDsReq = reqSetIDs
				} else {
					// The bits set in setListValue correspond to the number
					// of items that need to be worn for each list of magical properties
					// to be active
					for i := 0; i < 5; i++ {
						if setListValue&(1<<uint(i)) == 0 {
							continue
						}
						// bit position 0 means it requires >= 2 items worn, etc
						item.SetAttributesNumReq = append(item.SetAttributesNumReq, uint(i+2))
					}
				}
			}

			if item.GivenRuneword == 1 {
				runewordAttrList, err := parseMagicalList(&br)
				if err != nil {
					return itemList, err
				}

				item.RunewordAttributes = runewordAttrList
			}
		}

		if item.LocationID == socketed {
			last := len(itemList) - 1

			// The socketed item is a weapon, so we'll read the socketed properties
			// from the weapons map.
			if itemList[last].TypeID == weapon {
				if attrList, ok := socketablesWeapons[item.Type]; ok {
					item.MagicAttributes = attrList
				}
			}

			// The socketed item is an armor piece, so we'll read the socketed properties
			// from the armor map.
			if itemList[last].TypeID == armor {
				if attrList, ok := socketablesArmor[item.Type]; ok {
					item.MagicAttributes = attrList
				}
			}

			// The socketed item is a shield, so we'll read the socketed properties
			// from the shield map.
			if itemList[last].TypeID == shield {
				if attrList, ok := socketablesShields[item.Type]; ok {
					item.MagicAttributes = attrList
				}
			}

			// Add item to the socket list
			itemList[last].SocketedItems = append(itemList[last].SocketedItems, item)

		} else {
			// Ok, so this item it self is not in a socket, but it might have socketed
			// items in it, if it does, we'll need to increment the number of total
			// items we read in this loop, since the items glued into this item sockets,
			// will follow directly after this item.
			//
			// We'll also make sure the item isn't a simple item, because apparently
			// some quest items like Mephisto's Soul Stone has 2 sockets, but
			// no items in it.
			if item.NrOfItemsInSockets > 0 && item.SimpleItem == 0 {
				numberOfItemsToRead += int(item.NrOfItemsInSockets)
			}

			itemList = append(itemList, item)
		}

		// If the item is not byte aligned, we'll have to byte align it before
		// reading the next item.
		br.Align()
	}

	return itemList, nil
}

// Parses the 108 bits of data all items have, both simple items and extended items.
func parseSimpleBits(br *bitReader, item *Item) error {
	// offset: 0 "J"
	j, err := br.ReadByte()
	if err != nil {
		return err
	}

	// offset: 8, "M"
	m, err := br.ReadByte()
	if err != nil {
		return err
	}

	if string(j) != "J" || string(m) != "M" {
		return errors.New("failed to find item header JM")
	}

	// offset: 16, unknown
	br.ReadBits(4)

	// offset: 20
	if item.Identified, err = br.ReadBits(1); err != nil {
		return err
	}

	// offset: 21, unknown
	br.ReadBits(6)

	// offset: 27
	if item.Socketed, err = br.ReadBits(1); err != nil {
		return err
	}

	// offset 28, unknown
	br.ReadBits(1)

	// offset 29
	if item.New, err = br.ReadBits(1); err != nil {
		return err
	}

	// offset 30, unknown
	br.ReadBits(2)

	// offset 32
	if item.IsEar, err = br.ReadBits(1); err != nil {
		return err
	}

	// offset 33
	if item.StarterItem, err = br.ReadBits(1); err != nil {
		return err
	}

	// offset 34, unknown
	br.ReadBits(3)

	// offset 37, if it's a simple item, it only contains 111 bits data
	if item.SimpleItem, err = br.ReadBits(1); err != nil {
		return err
	}

	// offset 38
	if item.Ethereal, err = br.ReadBits(1); err != nil {
		return err
	}

	// offset 39, unknown
	br.ReadBits(1)

	// offset 40
	if item.Personalized, err = br.ReadBits(1); err != nil {
		return err
	}

	// offset 41, unknown
	br.ReadBits(1)

	// offset 42
	if item.GivenRuneword, err = br.ReadBits(1); err != nil {
		return err
	}

	// offset 43, unknown
	br.ReadBits(5)

	// offset 48
	if item.Version, err = br.ReadBits(8); err != nil {
		return err
	}

	// offset 56, unknown
	br.ReadBits(2)

	// offset 58
	if item.LocationID, err = br.ReadBits(3); err != nil {
		return err
	}

	// offset 61
	if item.EquippedID, err = br.ReadBits(4); err != nil {
		return err
	}

	// offset 65
	if item.PositionX, err = br.ReadBits(4); err != nil {
		return err
	}

	// offset 69
	if item.PositionY, err = br.ReadBits(3); err != nil {
		return err
	}

	// offset 72
	br.ReadBits(1)

	// offset 73, if item is neither equipped or in the belt, this tells us where it is.
	if item.AltPositionID, err = br.ReadBits(3); err != nil {
		return err
	}

	if item.IsEar == 0 {
		// offset 76, item type, 4 chars, each 8 bit (not byte aligned)
		var itemType string
		for j := 0; j < 4; j++ {
			t, err := br.ReadBits(8)
			if err != nil {
				return err
			}
			itemType += fmt.Sprintf("%c", t)
		}

		item.Type = strings.Trim(itemType, " ")
		item.TypeID = item.getTypeID()

		switch item.TypeID {
		case armor:
			typeName, ok := armorCodes[item.Type]
			if ok {
				item.TypeName = typeName
			}
		case shield:
			typeName, ok := shieldCodes[item.Type]
			if ok {
				item.TypeName = typeName
			}
		case weapon:
			typeName, ok := weaponCodes[item.Type]
			if ok {
				item.TypeName = typeName
			}

			// Weapons have base damage, so we'll check our
			// map for the base damage of this weapon type.
			baseDamage, ok := weaponDamageMap[item.Type]
			if ok {

				// If the item is ethereal we need to add 50% enhanced
				// damage to the base damage.
				if item.Ethereal == 1 {
					baseDamage.Min = int((float64(baseDamage.Min) * 1.5))
					baseDamage.Max = int((float64(baseDamage.Max) * 1.5))
					baseDamage.TwoMin = int((float64(baseDamage.TwoMin) * 1.5))
					baseDamage.TwoMax = int((float64(baseDamage.TwoMax) * 1.5))
				}

				item.BaseDamage = &baseDamage
			}
		case other:
			typeName, ok := miscCodes[item.Type]
			if ok {
				item.TypeName = typeName
			}
		}

		// offset 108
		// If sockets exist, read the items, they'll be 108 bit basic items * nrOfSockets
		if item.NrOfItemsInSockets, err = br.ReadBits(3); err != nil {
			return err
		}
	} else {
		// offset 76, the item is an ear, we need to read the ear data.
		earClass, err := br.ReadBits(3)
		if err != nil {
			return err
		}

		earLevel, err := br.ReadBits(7)
		if err != nil {
			return err
		}

		var name string
		for {
			c, err := br.ReadBits(7)
			if err != nil {
				return err
			}

			if c == 0 {
				break
			}

			name += fmt.Sprintf("%c", c)
		}

		item.EarAttributes = earAttributes{
			Class: earClass,
			Level: earLevel,
			Name:  name,
		}

		// If the ear is not byte aligned, we'll have to byte align it before
		// reading the next property.
		br.Align()
	}

	return nil
}

// Parses the rare or crafted bits that only exists on rare and crafted items,
// the bit length may vary depending on how many properties the item have.
func parseRareOrCraftedBits(br *bitReader, item *Item) error {
	nameID1, err := br.ReadBits(8)
	if err != nil {
		return err
	}

	name1, ok := rareNames[nameID1]
	if !ok {
		return fmt.Errorf("unknown rare name id: %d", nameID1)
	}

	item.RareName = name1

	nameID2, err := br.ReadBits(8)
	if err != nil {
		return err
	}

	name2, ok := rareNames[nameID2]
	if !ok {
		return fmt.Errorf("unknown rare name id: %d", nameID2)
	}

	item.RareName2 = name2

	// Array of 6 possible prefixes and suffixes. First read 1 bit. If 1, read 11 more
	// for the prefix/suffix id defined in MagicPrefix.txt and MagicSuffix.txt.
	// Even indices are prefixes, odd suffixes.
	for i := 0; i < 6; i++ {
		prefix, err := br.ReadBits(1)
		if err != nil {
			return err
		}

		if prefix == 1 {
			magicalID, err := br.ReadBits(11)
			if err != nil {
				return err
			}
			item.MagicalNameIDs = append(item.MagicalNameIDs, magicalID)
		} else {
			item.MagicalNameIDs = append(item.MagicalNameIDs, 0)
		}
	}

	return nil
}

func parseRunewordBits(br *bitReader, item *Item) error {
	runewordID, err := br.ReadBits(12)
	if err != nil {
		return err
	}

	item.RunewordID = runewordID

	if runewordName, ok := runewordNames[runewordID]; ok {
		item.RunewordName = runewordName
	}

	// Unknown 4 bits, seems to be 5 all the time.
	br.ReadBits(4)

	return nil
}

// Parses the magical property list in the byte queue that belongs to an item
// and returns the list of properties.
func parseMagicalList(br *bitReader) ([]MagicAttribute, error) {
	var magicAttributes []MagicAttribute
	for {
		id, err := br.ReadBits(9)
		if err != nil {
			return magicAttributes, err
		}

		// If all 9 bits are set, we've hit the end of the stats section
		//  at 0x1ff and exit the loop.
		if id == 0x1ff {
			break
		}

		prop, ok := magicalProperties[id]
		if !ok {
			return magicAttributes, fmt.Errorf("unknown magical property: %d", id)
		}

		var values []int64
		for _, bitLength := range prop.Bits {
			val, err := br.ReadBits(bitLength)
			if err != nil {
				return magicAttributes, err
			}

			if prop.Bias != 0 {
				val -= prop.Bias
			}

			values = append(values, int64(val))
		}

		attr := MagicAttribute{
			ID:     id,
			Name:   prop.Name,
			Values: values,
		}

		magicAttributes = append(magicAttributes, attr)
	}

	return magicAttributes, nil
}
