package d2s

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

// Char is the public type of the character.
type Char struct{}

// Character represents all the d2s character data.
type character struct {
	header
	attributes
	skills []skill
	items  []Item
}

// Header determines the header data of a d2s file.
type header struct {
	Identifier        uint32     // : 0 4 bytes
	Version           uint32     // : 4 4 bytes
	FileSize          uint32     // : 8 4 bytes
	CheckSum          uint32     // : 12 4 bytes
	ActiveArms        uint32     // : 16 4 bytes
	Name              name       // : 20 16 bytes
	Status            byte       // : 36 4 bytes
	Progression       byte       // : 37 1 bytes
	_                 [2]byte    // : 38 2 bytes
	Class             class      // : 40 1 bytes
	_                 [2]byte    // : 41 2 bytes
	Level             byte       // : 43 1 bytes
	_                 [4]byte    // : 44 4 bytes
	LastPlayed        uint32     // : 48 4 bytes
	_                 [4]byte    // : 52 4 bytes
	AssignedSkills    [16]uint32 // : 56 64 bytes
	LeftSkill         uint32     // : 120 4 bytes
	RightSkill        uint32     // : 124 4 bytes
	LeftSwapSkill     uint32     // : 128 4 bytes
	RightSwapSkill    uint32     // : 132 4 bytes
	_                 [32]byte   // : 136 32 bytes
	CurrentDifficulty difficulty // : 168 3 bytes
	MapID             uint32     // : 171 4 bytes
	_                 [2]byte    // : 175 2 bytes
	DeadMerc          uint16     // : 177 2 bytes
	MercID            uint32     // : 179 4 bytes
	MercNameID        uint16     // : 183 2 bytes
	MercType          uint16     // : 185 2 bytes
	MercExp           uint32     // : 187 4 bytes
	_                 [144]byte  // : 191 144 bytes
	QuestHeader       [4]byte    // : 335 4 bytes
	_                 [6]byte    // : 339 6 bytes
	QuestsNormal      [96]byte   // : 345 96 bytes
	QuestsNm          [96]byte   // : 441 96 bytes
	QuestsHell        [96]byte   // : 537 96 bytes
	WaypointHeader    [2]byte    // : 633 2 bytes
	_                 [6]byte    // : 635 6 bytes
	WaypointsNormal   [24]byte   // : 641 24 bytes
	WaypointsNm       [24]byte   // : 665 24 bytes
	WaypointsHell     [24]byte   // : 689 24 bytes
	WaypointTrailer   byte       // : 713 1 byte
	NPCHeader         [2]byte    // : 714 2 byte
	_                 byte       // : 716 1 byte
	NPCIntroNormal    [5]byte    // : 717 5 byte
	_                 [3]byte    // : 722 3 byte
	NPCIntroNm        [5]byte    // : 725 5 byte
	_                 [3]byte    // : 730 3 byte
	NPCIntroHell      [5]byte    // : 733 5 byte
	_                 [3]byte    // : 738 3 byte
	NPCReturnNorm     [4]byte    // : 741 4 byte
	_                 [4]byte    // : 745 4 byte
	NPCReturnNm       [4]byte    // : 749 4 byte
	_                 [4]byte    // : 753 4 byte
	NPCReturnHell     [4]byte    // : 757 4 byte
	_                 [4]byte    // : 761 4 byte
	StatHeader        [2]byte    // : 765 2 byte
}

type skillData struct {
	Header [2]byte
	List   [30]byte
}

type itemData struct {
	Header [2]byte
	Count  uint16
}

// Parse will read the data from a d2s character file and return a normalized struct.
func Parse(file io.Reader) (Char, error) {

	// Implements buffered reading, wraps io.Reader.
	bfr := bufio.NewReader(file)

	// Create character, we'll pass it around.
	char := character{}

	err := parseHeader(bfr, &char)

	if err != nil {
		return Char{}, err
	}

	err = parseAttributes(bfr, &char)

	if err != nil {
		return Char{}, err
	}

	err = parseSkills(bfr, &char)

	if err != nil {
		return Char{}, err
	}

	err = parseItems(bfr, &char)

	if err != nil {
		return Char{}, err
	}

	/*for _, item := range char.items {
		fmt.Printf("\n%+v\n\n\n", item)
	}*/

	return Char{}, nil
}

// Parses the 767 byte static header, all characters have this data.
func parseHeader(bfr io.Reader, char *character) error {

	// Make a buffer that can hold 767 bytes, which can hold the entire header.
	buf := make([]byte, 767)

	_, err := io.ReadFull(bfr, buf)
	if err != nil {
		return err
	}

	err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &char.header)
	if err != nil {
		return err
	}

	return nil
}

// Parses the total attributes the character has, before items are calculated.
func parseAttributes(bfr io.ByteReader, char *character) error {

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
			return fmt.Errorf("Unknown attribute id: %d", id)
		}

		// The attribute value.
		attr := reverseBits(br.ReadBits64(length, true), length)
		if br.Err() != nil {
			return br.Err()
		}

		switch id {
		case strength:
			char.attributes.Strength = attr
		case energy:
			char.attributes.Energy = attr
		case dexterity:
			char.attributes.Dexterity = attr
		case vitality:
			char.attributes.Vitality = attr
		case unusedStats:
			char.attributes.UnusedStats = attr
		case unusedSkills:
			char.attributes.UnusedSkillPoints = attr
		case currentHP:
			char.attributes.CurrentHP = attr / 256
		case maxHP:
			char.attributes.MaxHP = attr / 256
		case currentMana:
			char.attributes.CurrentMana = attr / 256
		case maxMana:
			char.attributes.MaxMana = attr / 256
		case currentStamina:
			char.attributes.CurrentStamina = attr / 256
		case maxStamina:
			char.attributes.MaxStamina = attr / 256
		case level:
			char.attributes.Level = attr
		case experience:
			char.attributes.Experience = attr
		case gold:
			char.attributes.Gold = attr
		case stashedGold:
			char.attributes.StashedGold = attr
		}
	}

	return nil
}

// Parses the skill set, this contains of a 2 byte header and a 30 byte list of
// 30 indexes with the allocated points for that skill. The class determines
// what offset we start reading from in the skill list.
func parseSkills(bfr io.Reader, char *character) error {

	// Make a buffer that can hold 32 bytes, which can hold the entire skillset.
	buf := make([]byte, 32)

	_, err := io.ReadFull(bfr, buf[:32])
	if err != nil {
		return err
	}

	rawSkills := skillData{}
	err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &rawSkills)

	if err != nil {
		return err
	}

	if string(rawSkills.Header[:]) != "if" {
		return errors.New("Failed to find skill header")
	}

	skillOffset, ok := skillOffsetMap[uint(char.header.Class)]
	if !ok {
		return fmt.Errorf("Unknown skill offset for class %d", char.header.Class)
	}

	for i, allocatedPoints := range rawSkills.List {
		id := (i + skillOffset)
		s := skill{
			id:     id,
			points: int(allocatedPoints),
			name:   skillMap[id],
		}
		char.skills = append(char.skills, s)
	}

	return nil
}

// Parses all items on the character, this includes items in stash, cube,
// inventory and equipped items.
func parseItems(bfr io.ByteReader, char *character) error {

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
		return errors.New("Failed to find the items header")
	}

	ibr := bitReader{r: bfr}

	// We'll start this number at items count, but the thing is, if an item
	// has items socketed on it, they will follow the item they're socketed in,
	// so every time we find an item with socketed items, we'll increment this
	// list in order to read them as well.
	numberOfItemsToRead := int(itemHeaderData.Count)

	for i := 0; i < numberOfItemsToRead; i++ {
		var readBits int
		item := Item{}

		// Read the 111 bit basic item structure, all items have this structure.
		err = parseSimpleBits(&ibr, &item)
		readBits += 111

		if err != nil {
			return err
		}

		if item.SimpleItem == 0 {
			// offset 111, item id is 8 chars, each 4 bit
			item.ID = reverseBits(ibr.ReadBits64(32, true), 32)
			readBits += 32

			// offset 143
			item.Level = reverseBits(ibr.ReadBits64(7, true), 7)
			readBits += 7

			// offset 150
			item.Quality = reverseBits(ibr.ReadBits64(4, true), 4)
			readBits += 4

			// If this is true, it means the item has more than one picture associated
			// with it.
			item.MultiplePictures = reverseBits(ibr.ReadBits64(1, true), 1)
			readBits++

			if item.MultiplePictures == 1 {
				// The next 3 bits contain the picture ID.
				item.PictureID = reverseBits(ibr.ReadBits64(3, true), 3)
				readBits += 3
			}

			// If this is true, it means the item is class specific.
			item.ClassSpecific = reverseBits(ibr.ReadBits64(1, true), 1)
			readBits++

			// If the item is class specific, the next 11 bits will
			// contain the class specific data.
			if item.ClassSpecific == 1 {
				// TODO: Parse this into something useful
				reverseBits(ibr.ReadBits64(11, true), 11)
				readBits += 11
			}

			// MARK: Quality based data.

			switch item.Quality {

			case lowQuality:
				item.LowQualityID = reverseBits(ibr.ReadBits64(3, true), 3)
				readBits += 3

			case normal:
				// No extra data present

			case highQuality:
				// TODO: Figure out what these 3 bits are on a high quality item
				reverseBits(ibr.ReadBits64(3, true), 3)
				readBits += 3

			case magicallyEnhanced:
				item.MagicPrefix = reverseBits(ibr.ReadBits64(11, true), 11)
				item.MagicSuffix = reverseBits(ibr.ReadBits64(11, true), 11)
				readBits += 22

			case partOfSet:
				item.SetID = reverseBits(ibr.ReadBits64(12, true), 12)
				readBits += 12

			case rare, crafted:
				rBits, _ := parseRareOrCraftedBits(&ibr, &item)
				readBits += rBits

			case unique:
				item.UniqueID = reverseBits(ibr.ReadBits64(12, true), 12)
				readBits += 12
			}

			// MARK: Runeword data
			// Parse 16 bits here if the item has been given a runeword.
			if item.GivenRuneword == 1 {
				parseRunewordBits(&ibr, &item)
				readBits += 16
			}

			// MARK: Personalization data
			// TODO: Parse Personalization data here if the item is personalized.

			// MARK: Structure - all items have this part.

			// If the item is a tomb, it will contain 5 extra bits, we're not
			// interested in these bits, they value is usually 1, but not sure
			// what it is.
			if tombMap[item.Type] {
				reverseBits(ibr.ReadBits64(5, true), 5)
				readBits += 5
			}

			// All items have this field between the personalization (if it exists)
			// and the item specific data
			fmt.Printf("Bits read before timestamp: %d \n", readBits)
			item.Timestamp = reverseBits(ibr.ReadBits64(1, true), 1)
			readBits++

			typeID := item.getTypeID()

			if typeID == armor {
				fmt.Printf("Bits read before armor: %d \n", readBits)
				// TODO: Figure out which items have 11 bits of defense rating data.
				// If the item is an armor, it will have this field of defense data.
				defRating := reverseBits(ibr.ReadBits64(11, true), 11)

				// We need to substract 10 defense rating from all armors for
				// some reason, I'm not sure why.
				item.DefenseRating = int64((defRating - 10))

				fmt.Printf("Defense rating: %d \n", item.DefenseRating)
				readBits += 11
			}

			if typeID == armor || typeID == weapon {
				item.MaxDurability = reverseBits(ibr.ReadBits64(8, true), 8)
				readBits += 8
				item.CurrentDurability = reverseBits(ibr.ReadBits64(8, true), 8)
				readBits += 8

				// Seems to be a random bit here.
				reverseBits(ibr.ReadBits64(1, true), 1)
				readBits++
			}

			if quantityMap[item.Type] {
				// If the item is a stacked item, e.g. a javelin or something, these 9
				// bits will contain the quantity.
				item.Quantity = reverseBits(ibr.ReadBits64(9, true), 9)
				readBits += 9
			}

			// If the item is socketed, it will contain 4 bits of data which are the nr
			// of total sockets the item have, regardless of how many are occupied by
			// an item.
			if item.Socketed == 1 {
				item.TotalNrOfSockets = reverseBits(ibr.ReadBits64(4, true), 4)
				readBits += 4
			}

			// If the item is part of a set, these bit will tell us how many lists
			// of magical properties follow the one regular magical property list.
			if item.Quality == partOfSet {
				setListValue := reverseBits(ibr.ReadBits64(5, true), 5)
				readBits += 5

				listCount, ok := setListMap[setListValue]
				if !ok {
					return fmt.Errorf("Unknown set list number %d", setListValue)
				}

				item.SetListCount = listCount
			}

			fmt.Printf("Read bits until magic list: %d \n", readBits)

			// MARK: Time to parse 9 bit magical property ids followed by their n bit
			// length values, but only if the item is magical or above.
			magicAttrList, rb := parseMagicalList(&ibr)
			readBits += rb

			item.MagicAttributes = magicAttrList

			fmt.Printf("Read bits after magic list: %d \n", readBits)

			// Item has more magical property lists due to being a set item.
			if item.SetListCount > 0 {
				for i := 0; i < int(item.SetListCount); i++ {
					setAttrList, rb := parseMagicalList(&ibr)
					readBits += rb

					item.SetAttributes = append(item.SetAttributes, setAttrList)
					fmt.Printf("Read bits after set list: %d \n", readBits)
				}
			}

			if item.GivenRuneword == 1 {
				fmt.Printf("Read bits before runeword list: %d \n", readBits)
				runewordAttrList, rb := parseMagicalList(&ibr)
				readBits += rb
				item.RunewordAttributes = runewordAttrList
			}
		}

		if item.LocationID == socketed {
			last := len(char.items) - 1
			char.items[last].SocketedItems = append(char.items[last].SocketedItems, item)
			fmt.Printf("Found socketed item: %+v\n\n", item)
		} else {
			// Ok, so this item it self is not in a socket, but it might have socketed
			// items in it, if it does, we'll need to increment the number of total
			// items we read in this loop, since the items glued into this item sockets,
			// will follow directly after this item.
			if item.NrOfItemsInSockets > 0 {
				numberOfItemsToRead += int(item.NrOfItemsInSockets)
				fmt.Println(numberOfItemsToRead)
			}
			char.items = append(char.items, item)
		}

		fmt.Printf("%+v\n\n", item)

		// If the item is not byte aligned, we'll have to byte align it before
		// reading the next item, so we'll simply queue the reader at the next
		// byte boundry by calculating the remainder.
		remainder := readBits % 8
		if remainder > 0 {
			bitsToAlign := uint(8 - remainder)
			reverseBits(ibr.ReadBits64(bitsToAlign, true), bitsToAlign)
		}
	}

	return nil
}

// Parses the 108 bits of data all items have, both simple items and extended items.
func parseSimpleBits(ibr *bitReader, item *Item) error {

	// offset: 0 "J"
	j := ibr.ReadBits64(8, false)

	// offset: 8, "M"
	m := ibr.ReadBits64(8, false)

	if string(j) != "J" || string(m) != "M" {
		return errors.New("Failed to find item header JM")
	}

	ibr.ReadBits64(4, true)

	// offset: 20
	item.Identified = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset: 21, unknown
	ibr.ReadBits64(6, true)

	// offset: 27
	item.Socketed = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 28, unknown
	ibr.ReadBits64(1, true)

	// offset 29
	item.New = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 30, unknown
	reverseBits(ibr.ReadBits64(2, true), 2)

	// offset 32
	item.IsEar = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 33
	item.StarterItem = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 34, unknown
	reverseBits(ibr.ReadBits64(3, true), 3)

	// offset 37, if it's a simple item, it only contains 111 bits data
	item.SimpleItem = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 38
	item.Ethereal = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 39, unknown
	reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 40
	item.Personalized = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 41, unknown
	reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 42
	item.GivenRuneword = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 43, unknown
	reverseBits(ibr.ReadBits64(15, true), 15)

	// offset 58
	item.LocationID = reverseBits(ibr.ReadBits64(3, true), 3)

	// offset 61
	item.EquippedID = reverseBits(ibr.ReadBits64(4, true), 4)

	// offset 65
	item.PositionY = reverseBits(ibr.ReadBits64(4, true), 4)

	// offset 69
	item.PositionX = reverseBits(ibr.ReadBits64(3, true), 3)

	// offset 72
	reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 73, if item is neither equipped or in the belt, this tells us where it is.
	item.AltPositionID = reverseBits(ibr.ReadBits64(3, true), 3)

	// offset 76, item type, 4 chars, each 8 bit (not byte aligned)
	var itemType string
	for j := 0; j < 4; j++ {
		itemType += string(reverseBits(ibr.ReadBits64(8, true), 8))
	}

	item.Type = strings.Trim(itemType, " ")

	// offset 108
	// If sockets exist, read the items, they'll be 108 bit basic items * nrOfSockets
	item.NrOfItemsInSockets = reverseBits(ibr.ReadBits64(3, true), 3)

	return nil
}

// Parses the rare or crafted bits that only exists on rare and crafted items,
// the bit length may vary depending on how many properties the item have.
func parseRareOrCraftedBits(ibr *bitReader, item *Item) (int, error) {

	var readBits int

	nameID1 := reverseBits(ibr.ReadBits64(8, true), 8)
	name1, ok := rareNames[nameID1]
	if !ok {
		log.Fatalf("Unknown rare name id: %d", nameID1)
	}
	readBits += 8

	item.RareName = name1

	nameID2 := reverseBits(ibr.ReadBits64(8, true), 8)
	name2, ok := rareNames[nameID2]
	if !ok {
		log.Fatalf("Unknown rare name id: %d", nameID2)
	}
	readBits += 8

	item.RareName2 = name2

	// Following the name IDs, we got 6 possible magical name IDs, the pattern
	// is 1 bit id, 11 bit value... But the value will only exist if the prefix
	// is 1. So we'll read the id first and check it against 1.
	for i := 0; i < 6; i++ {
		prefix := reverseBits(ibr.ReadBits64(1, true), 1)
		readBits++

		if prefix == 1 {
			item.MagicalNameIDs[i] = reverseBits(ibr.ReadBits64(11, true), 11)
			readBits += 11
		}
	}

	return readBits, nil
}

func parseRunewordBits(ibr *bitReader, item *Item) error {
	runewordID := reverseBits(ibr.ReadBits64(12, true), 12)
	item.RunewordID = runewordID

	// Unknown 4 bits, seems to be 5 all the time.
	reverseBits(ibr.ReadBits64(4, true), 4)

	return nil
}

// Parses the upcoming magical property list in the byte queue,
// and return a list of properties.
func parseMagicalList(ibr *bitReader) ([]magicAttribute, int) {

	var magicAttributes []magicAttribute
	var readBits int

	for {
		id := reverseBits(ibr.ReadBits64(9, true), 9)
		readBits += 9

		if ibr.Err() != nil {
			fmt.Println(ibr.Err())
		}

		// If all 9 bits are set, we've hit the end of the stats section
		//  at 0x1ff and exit the loop.
		if id == 0x1ff {
			fmt.Println("breaking out at 511")
			break
		}

		prop, ok := magicalProperties[id]
		if !ok {
			log.Fatalf("Unknown magical property: %d", id)
		}

		var values []int64
		for _, bitLength := range prop.Bits {

			val := reverseBits(ibr.ReadBits64(bitLength, true), bitLength)
			readBits += int(bitLength)
			fmt.Printf("found id: %d, reading bit size field: %d:, value is: %d\n", id, bitLength, val)

			if prop.Bias != 0 {
				val = val - prop.Bias
			}

			values = append(values, int64(val))
		}

		attr := magicAttribute{
			ID:     id,
			Name:   prop.Name,
			Values: values,
		}

		magicAttributes = append(magicAttributes, attr)
		fmt.Printf("bits read after property id %d: %d \n", id, readBits)
	}

	return magicAttributes, readBits
}
