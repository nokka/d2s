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

	// Perform the actual reading.
	return parse(bfr)
}

// ParseFromContent will read the character from a byte slice.
func ParseFromContent(data []byte) (*Character, error) {
	// Create a reader from the byte slice.
	r := bytes.NewReader(data)

	// Implements buffered reading, wraps io.Reader.
	bfr := bufio.NewReader(r)

	// Perform the actual reading.
	return parse(bfr)
}

// parse will do the actual reading from an io.Reader.
func parse(bfr *bufio.Reader) (*Character, error) {
	// Create character pointer, we'll pass it around.
	char := &Character{}

	if err := parseHeader(bfr, char); err != nil {
		return nil, encodeError(char, err.Error())
	}

	if err := parseAttributes(bfr, char); err != nil {
		return nil, encodeError(char, err.Error())
	}

	if err := parseSkills(bfr, char); err != nil {
		return nil, encodeError(char, err.Error())
	}

	if err := parseItems(bfr, char); err != nil {
		return nil, encodeError(char, err.Error())
	}

	if err := parseCorpse(bfr, char); err != nil {
		return nil, encodeError(char, err.Error())
	}

	// Normalize the character status, that is being stored as a byte.
	status := char.Header.Status.Readable()

	if status.Expansion {
		if err := parseMercItems(bfr, char); err != nil {
			return nil, encodeError(char, err.Error())
		}
	}

	if char.Header.Class == Necromancer && status.Expansion {
		if err := parseIronGolem(bfr, char); err != nil {
			return nil, encodeError(char, err.Error())
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
		id := reverseBits(br.ReadBits64(9, true), 9)

		if br.Err() != nil {
			return br.Err()
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
		attr := reverseBits(br.ReadBits64(length, true), length)
		if br.Err() != nil {
			return br.Err()
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

	_, err := io.ReadFull(bfr, buf[:32])
	if err != nil {
		return err
	}

	skillHeaderData := skillData{}
	err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &skillHeaderData)

	if err != nil {
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
		return errors.New("failed to find the items header")
	}

	items, err := ParseItemList(bfr, int(itemHeaderData.Count))
	if err != nil {
		return err
	}

	char.Items = items

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
	ibr := bitReader{r: bfr}

	// offset: 0 "j"
	j := ibr.ReadBits64(8, false)

	// offset: 8, "f"
	f := ibr.ReadBits64(8, false)

	if fmt.Sprintf("%c", j) != "j" || fmt.Sprintf("%c", f) != "f" {
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

func ParseItemList(bfr io.ByteReader, itemCount int) ([]Item, error) {
	var itemList []Item

	ibr := bitReader{r: bfr}

	// We'll start this number at items count, but the thing is, if an item
	// has items socketed on it, they will follow the item they're socketed in,
	// so every time we find an item with socketed items, we'll increment this
	// list in order to read them as well.
	numberOfItemsToRead := itemCount

	for i := 0; i < numberOfItemsToRead; i++ {
		var readBits int
		parsed := Item{}

		// Read the 111 bit basic item structure, all items have this structure.
		err := parseSimpleBits(&ibr, &parsed)
		readBits += 111

		if err != nil {
			return []Item{}, err
		}

		if parsed.SimpleItem == 0 {
			// offset 111, item id is 8 chars, each 4 bit
			parsed.ID = reverseBits(ibr.ReadBits64(32, true), 32)
			readBits += 32

			// offset 143
			parsed.Level = reverseBits(ibr.ReadBits64(7, true), 7)
			readBits += 7

			// offset 150
			parsed.Quality = reverseBits(ibr.ReadBits64(4, true), 4)
			readBits += 4

			// If this is true, it means the item has more than one picture associated
			// with it.
			parsed.MultiplePictures = reverseBits(ibr.ReadBits64(1, true), 1)
			readBits++

			if parsed.MultiplePictures == 1 {
				// The next 3 bits contain the picture ID.
				parsed.PictureID = reverseBits(ibr.ReadBits64(3, true), 3)
				readBits += 3
			}

			// If this is true, it means the item is class specific.
			parsed.ClassSpecific = reverseBits(ibr.ReadBits64(1, true), 1)
			readBits++

			// If the item is class specific, the next 11 bits will
			// contain the class specific data.
			if parsed.ClassSpecific == 1 {
				// TODO: Parse this into something useful
				_ = reverseBits(ibr.ReadBits64(11, true), 11)
				readBits += 11
			}

			switch parsed.Quality {
			case lowQuality:
				parsed.LowQualityID = reverseBits(ibr.ReadBits64(3, true), 3)
				readBits += 3

			case normal:
			// No extra data present

			case highQuality:
				// TODO: Figure out what these 3 bits are on a high quality item
				_ = reverseBits(ibr.ReadBits64(3, true), 3)
				readBits += 3

			case magicallyEnhanced:
				parsed.MagicPrefix = reverseBits(ibr.ReadBits64(11, true), 11)
				prefixName, ok := magicalPrefixes[parsed.MagicPrefix]
				if ok {
					parsed.MagicPrefixName = prefixName
				}

				parsed.MagicSuffix = reverseBits(ibr.ReadBits64(11, true), 11)
				suffixName, ok := magicalSuffixes[parsed.MagicSuffix]
				if ok {
					parsed.MagicSuffixName = suffixName
				}
				readBits += 22

			case partOfSet:
				parsed.SetID = reverseBits(ibr.ReadBits64(12, true), 12)

				setName, ok := setNames[parsed.SetID]
				if ok {
					parsed.SetName = setName
				}

				readBits += 12

			case rare, crafted:
				rBits, _ := parseRareOrCraftedBits(&ibr, &parsed)
				readBits += rBits

			case unique:
				parsed.UniqueID = reverseBits(ibr.ReadBits64(12, true), 12)

				uniqueName, ok := uniqueNames[parsed.UniqueID]
				if ok {
					parsed.UniqueName = uniqueName
				}

				readBits += 12
			}

			// MARK: Runeword data
			if parsed.GivenRuneword == 1 {
				parseRunewordBits(&ibr, &parsed)
				readBits += 16
			}

			if parsed.Personalized == 1 {
				var name string
				for {
					c := reverseBits(ibr.ReadBits64(7, true), 7)
					readBits += 7

					if c == 0 {
						break
					}

					name += fmt.Sprintf("%c", c)
				}

				parsed.PersonalizedName = name
			}

			// If the item is a tome, it will contain 5 extra bits, we're not
			// interested in these bits, the value is usually 1, but not sure
			// what it is.
			if tomeMap[parsed.Type] {
				_ = reverseBits(ibr.ReadBits64(5, true), 5)
				readBits += 5
			}

			// All items have this field between the personalization (if it exists)
			// and the item specific data
			parsed.Timestamp = reverseBits(ibr.ReadBits64(1, true), 1)
			readBits++

			if parsed.TypeID == armor || parsed.TypeID == shield {
				// If the item is an armor, it will have this field of defense data.
				defRating := reverseBits(ibr.ReadBits64(11, true), 11)
				readBits += 11

				// We need to subtract 10 defense rating from all armors for
				// some reason, I'm not sure why.
				parsed.DefenseRating = int64((defRating - 10))
			}

			if parsed.TypeID == armor || parsed.TypeID == weapon || parsed.TypeID == shield {
				parsed.MaxDurability = reverseBits(ibr.ReadBits64(8, true), 8)
				readBits += 8

				// Some weapons like phase blades don't have durability, so we'll
				// check if the item has max durability, then we can safely assume
				// it has current durability too.
				if parsed.MaxDurability > 0 {
					parsed.CurrentDurability = reverseBits(ibr.ReadBits64(8, true), 8)
					// Seems to be a random bit here.
					_ = reverseBits(ibr.ReadBits64(1, true), 1)

					readBits += 9
				}
			}

			if quantityMap[parsed.Type] {
				// If the item is a stacked item, e.g. a javelin or something, these 9
				// bits will contain the quantity.
				parsed.Quantity = reverseBits(ibr.ReadBits64(9, true), 9)
				readBits += 9
			}

			// If the item is socketed, it will contain 4 bits of data which are the nr
			// of total sockets the item have, regardless of how many are occupied by
			// an item.
			if parsed.Socketed == 1 {
				parsed.TotalNrOfSockets = reverseBits(ibr.ReadBits64(4, true), 4)
				readBits += 4
			}

			// If the item is part of a set, these bit will tell us how many lists
			// of magical properties follow the one regular magical property list.
			var setListValue uint64 = 0
			if parsed.Quality == partOfSet {
				setListValue = reverseBits(ibr.ReadBits64(5, true), 5)
				readBits += 5

				listCount, ok := setListMap[setListValue]
				if !ok {
					return []Item{}, fmt.Errorf("unknown set list number %d", setListValue)
				}

				parsed.SetListCount = listCount
			}

			// MARK: Time to parse 9 bit magical property ids followed by their n bit
			// length values, but only if the item is magical or above.
			magicAttrList, rb, err := parseMagicalList(&ibr)
			readBits += rb

			if err != nil {
				return itemList, err
			}

			parsed.MagicAttributes = magicAttrList

			// Item has more magical property lists due to being a set item.
			if parsed.SetListCount > 0 {
				for i := 0; i < int(parsed.SetListCount); i++ {
					setAttrList, rb, err := parseMagicalList(&ibr)
					readBits += rb

					if err != nil {
						return itemList, err
					}

					parsed.SetAttributes = append(parsed.SetAttributes, setAttrList)
				}
				reqSetIDs, ok := setReqIDsMap[parsed.SetID]
				if ok {
					parsed.SetAttributesIDsReq = reqSetIDs
				} else {
					// The bits set in setListValue correspond to the number
					// of items that need to be worn for each list of magical properties
					// to be active
					for i := 0; i < 5; i++ {
						if setListValue&(1<<uint(i)) == 0 {
							continue
						}
						// bit position 0 means it requires >= 2 items worn, etc
						parsed.SetAttributesNumReq = append(parsed.SetAttributesNumReq, uint(i+2))
					}
				}
			}

			if parsed.GivenRuneword == 1 {
				runewordAttrList, rb, err := parseMagicalList(&ibr)
				readBits += rb

				if err != nil {
					return itemList, err
				}

				parsed.RunewordAttributes = runewordAttrList
			}
		}

		if parsed.LocationID == socketed {
			last := len(itemList) - 1

			// The socketed item is a weapon, so we'll read the socketed properties
			// from the weapons map.
			if itemList[last].TypeID == weapon {
				attrList, ok := socketablesWeapons[parsed.Type]
				if ok {
					parsed.MagicAttributes = attrList
				}
			}

			// The socketed item is an armor piece, so we'll read the socketed properties
			// from the armor map.
			if itemList[last].TypeID == armor {
				attrList, ok := socketablesArmor[parsed.Type]
				if ok {
					parsed.MagicAttributes = attrList
				}
			}

			// The socketed item is a shield, so we'll read the socketed properties
			// from the shield map.
			if itemList[last].TypeID == shield {
				attrList, ok := socketablesShields[parsed.Type]
				if ok {
					parsed.MagicAttributes = attrList
				}
			}

			// Add item to the socket list
			itemList[last].SocketedItems = append(itemList[last].SocketedItems, parsed)

		} else {
			// Ok, so this item it self is not in a socket, but it might have socketed
			// items in it, if it does, we'll need to increment the number of total
			// items we read in this loop, since the items glued into this item sockets,
			// will follow directly after this item.
			//
			// We'll also make sure the item isn't a simple item, because apparently
			// some quest items like Mephisto's Soul Stone has 2 sockets, but
			// no items in it.
			if parsed.NrOfItemsInSockets > 0 && parsed.SimpleItem == 0 {
				numberOfItemsToRead += int(parsed.NrOfItemsInSockets)
			}

			itemList = append(itemList, parsed)
		}

		// If the item is not byte aligned, we'll have to byte align it before
		// reading the next item, so we'll simply queue the reader at the next
		// byte boundary by calculating the remainder.
		remainder := readBits % 8
		if remainder > 0 {
			bitsToAlign := uint(8 - remainder)
			_ = reverseBits(ibr.ReadBits64(bitsToAlign, true), bitsToAlign)
		}
	}

	return itemList, nil
}

// Parses the 108 bits of data all items have, both simple items and extended items.
func parseSimpleBits(ibr *bitReader, item *Item) error {
	var readBits int
	// offset: 0 "J"
	j := ibr.ReadBits64(8, false)
	readBits += 8

	// offset: 8, "M"
	m := ibr.ReadBits64(8, false)
	readBits += 8

	if fmt.Sprintf("%c", j) != "J" || fmt.Sprintf("%c", m) != "M" {
		return errors.New("failed to find item header JM")
	}

	// offset: 16, unknown
	ibr.ReadBits64(4, true)
	readBits += 4

	// offset: 20
	item.Identified = reverseBits(ibr.ReadBits64(1, true), 1)
	readBits++

	// offset: 21, unknown
	ibr.ReadBits64(6, true)
	readBits += 6

	// offset: 27
	item.Socketed = reverseBits(ibr.ReadBits64(1, true), 1)
	readBits++

	// offset 28, unknown
	ibr.ReadBits64(1, true)
	readBits++

	// offset 29
	item.New = reverseBits(ibr.ReadBits64(1, true), 1)
	readBits++

	// offset 30, unknown
	_ = reverseBits(ibr.ReadBits64(2, true), 2)
	readBits += 2

	// offset 32
	item.IsEar = reverseBits(ibr.ReadBits64(1, true), 1)
	readBits++

	// offset 33
	item.StarterItem = reverseBits(ibr.ReadBits64(1, true), 1)
	readBits++

	// offset 34, unknown
	_ = reverseBits(ibr.ReadBits64(3, true), 3)
	readBits += 3

	// offset 37, if it's a simple item, it only contains 111 bits data
	item.SimpleItem = reverseBits(ibr.ReadBits64(1, true), 1)
	readBits++

	// offset 38
	item.Ethereal = reverseBits(ibr.ReadBits64(1, true), 1)
	readBits++

	// offset 39, unknown
	_ = reverseBits(ibr.ReadBits64(1, true), 1)
	readBits++

	// offset 40
	item.Personalized = reverseBits(ibr.ReadBits64(1, true), 1)
	readBits++

	// offset 41, unknown
	_ = reverseBits(ibr.ReadBits64(1, true), 1)
	readBits++

	// offset 42
	item.GivenRuneword = reverseBits(ibr.ReadBits64(1, true), 1)
	readBits++

	// offset 43, unknown
	_ = reverseBits(ibr.ReadBits64(5, true), 5)
	readBits += 5

	// offset 48
	item.Version = reverseBits(ibr.ReadBits64(8, true), 8)
	readBits += 8

	// offset 56, unknown
	_ = reverseBits(ibr.ReadBits64(2, true), 2)
	readBits += 2

	// offset 58
	item.LocationID = reverseBits(ibr.ReadBits64(3, true), 3)
	readBits += 3

	// offset 61
	item.EquippedID = reverseBits(ibr.ReadBits64(4, true), 4)
	readBits += 4

	// offset 65
	item.PositionX = reverseBits(ibr.ReadBits64(4, true), 4)
	readBits += 4

	// offset 69
	item.PositionY = reverseBits(ibr.ReadBits64(3, true), 3)
	readBits += 3

	// offset 72
	_ = reverseBits(ibr.ReadBits64(1, true), 1)
	readBits++

	// offset 73, if item is neither equipped or in the belt, this tells us where it is.
	item.AltPositionID = reverseBits(ibr.ReadBits64(3, true), 3)
	readBits += 3

	if item.IsEar == 0 {
		// offset 76, item type, 4 chars, each 8 bit (not byte aligned)
		var itemType string
		for j := 0; j < 4; j++ {
			itemType += fmt.Sprintf("%c", reverseBits(ibr.ReadBits64(8, true), 8))
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
		item.NrOfItemsInSockets = reverseBits(ibr.ReadBits64(3, true), 3)
	} else {

		// offset 76, the item is an ear, we need to read the ear data.
		earClass := reverseBits(ibr.ReadBits64(3, true), 3)
		readBits += 3

		earLevel := reverseBits(ibr.ReadBits64(7, true), 7)
		readBits += 7

		var name string
		for {
			c := reverseBits(ibr.ReadBits64(7, true), 7)

			if c == 0 {
				break
			}
			readBits += 7
			name += fmt.Sprintf("%c", c)
		}

		item.EarAttributes = earAttributes{
			Class: earClass,
			Level: earLevel,
			Name:  name,
		}

		// If the ear is not byte aligned, we'll have to byte align it before
		// reading the next property, so we'll simply queue the reader at the next
		// byte boundary by calculating the remainder.
		remainder := readBits % 8
		if remainder > 0 {
			bitsToAlign := uint(8 - remainder)
			_ = reverseBits(ibr.ReadBits64(bitsToAlign, true), bitsToAlign)
		}
	}

	return nil
}

// Parses the rare or crafted bits that only exists on rare and crafted items,
// the bit length may vary depending on how many properties the item have.
func parseRareOrCraftedBits(ibr *bitReader, item *Item) (int, error) {
	var readBits int

	nameID1 := reverseBits(ibr.ReadBits64(8, true), 8)
	readBits += 8

	name1, ok := rareNames[nameID1]
	if !ok {
		return readBits, fmt.Errorf("unknown rare name id: %d", nameID1)
	}

	item.RareName = name1

	nameID2 := reverseBits(ibr.ReadBits64(8, true), 8)
	readBits += 8

	name2, ok := rareNames[nameID2]
	if !ok {
		return readBits, fmt.Errorf("unknown rare name id: %d", nameID2)
	}

	item.RareName2 = name2

	// Array of 6 possible prefixes and suffixes. First read 1 bit. If 1, read 11 more
	// for the prefix/suffix id defined in MagicPrefix.txt and MagicSuffix.txt.
	// Even indices are prefixes, odd suffixes.
	for i := 0; i < 6; i++ {
		prefix := reverseBits(ibr.ReadBits64(1, true), 1)
		readBits++

		if prefix == 1 {
			item.MagicalNameIDs = append(item.MagicalNameIDs, reverseBits(ibr.ReadBits64(11, true), 11))
			readBits += 11
		} else {
			item.MagicalNameIDs = append(item.MagicalNameIDs, 0)
		}
	}

	return readBits, nil
}

func parseRunewordBits(ibr *bitReader, item *Item) {
	runewordID := reverseBits(ibr.ReadBits64(12, true), 12)
	item.RunewordID = runewordID

	runewordName, ok := runewordNames[runewordID]
	if ok {
		item.RunewordName = runewordName
	}

	// Unknown 4 bits, seems to be 5 all the time.
	_ = reverseBits(ibr.ReadBits64(4, true), 4)
}

// Parses the magical property list in the byte queue that belongs to an item
// and returns the list of properties.
func parseMagicalList(ibr *bitReader) ([]MagicAttribute, int, error) {

	var magicAttributes []MagicAttribute
	var readBits int

	for {
		id := reverseBits(ibr.ReadBits64(9, true), 9)
		readBits += 9

		if ibr.Err() != nil {
			return magicAttributes, readBits, ibr.Err()
		}

		// If all 9 bits are set, we've hit the end of the stats section
		//  at 0x1ff and exit the loop.
		if id == 0x1ff {
			break
		}

		prop, ok := magicalProperties[id]
		if !ok {
			return magicAttributes, readBits, fmt.Errorf("unknown magical property: %d", id)
		}

		var values []int64
		for _, bitLength := range prop.Bits {
			val := reverseBits(ibr.ReadBits64(bitLength, true), bitLength)
			readBits += int(bitLength)

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

	return magicAttributes, readBits, nil
}

func encodeError(char *Character, msg string) error {
	return fmt.Errorf("char name: %s, %s", char.Header.Name, msg)
}
