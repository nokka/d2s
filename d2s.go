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

// Character stuff
type Character struct {
	Header
	CharacterStats
	Skills []skill
}

// CharacterStats represents the characters stats
type CharacterStats struct {
	Strength          uint64
	Dexterity         uint64
	Vitality          uint64
	Energy            uint64
	UnusedStats       uint64
	UnusedSkillPoints uint64
	CurrentHP         uint64
	MaxHP             uint64
	CurrentMana       uint64
	MaxMana           uint64
	CurrentStamina    uint64
	MaxStamina        uint64
	Level             uint64
	Experience        uint64
	Gold              uint64
	StashedGold       uint64
}

// Header determines the header data of a d2s file.
type Header struct {
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

type itemHeader struct {
	Header [2]byte
	Count  uint16
}

// statsBitMap holds all the references to bit sites of all attributes.
var statsBitMap = map[uint64]uint{
	0:  10,
	1:  10,
	2:  10,
	3:  10,
	4:  10,
	5:  8,
	6:  21,
	7:  21,
	8:  21,
	9:  21,
	10: 21,
	11: 21,
	12: 7,
	13: 32,
	14: 25,
	15: 25,
}

var skillOffsetMap = map[uint]int{
	0x00: 6,
	0x01: 36,
	0x02: 66,
	0x03: 96,
	0x04: 126,
	0x05: 221,
	0x06: 251,
}

func parseHeader(bfr io.Reader, char *Character) error {

	// Make a buffer that can hold 767 bytes, which can hold the entire header.
	// We'll reuse this buffer through out to avoid another alloc.
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

func parseStats(br bitReader, char *Character) error {

	for {
		// 9 bit stat id, bit reversed twice.
		id := reverseBits(br.ReadBits64(9, true), 9)

		if br.Err() != nil {
			return br.Err()
		}

		// If all 9 bits are set, we've hit the end of the stats section
		//  at 0x1ff and exit the loop.
		if id == 0x1ff {
			break
		}

		// The stat value bit length, so we'll know how many bits to read next.
		bitLength, ok := statsBitMap[id]
		if !ok {
			return fmt.Errorf("Unknown stat id: %d", id)
		}

		// The stat value, bit reversed, twice.
		statVal := reverseBits(br.ReadBits64(bitLength, true), bitLength)
		if br.Err() != nil {
			return br.Err()
		}

		switch id {
		case 0:
			char.CharacterStats.Strength = statVal
		case 1:
			char.CharacterStats.Energy = statVal
		case 2:
			char.CharacterStats.Dexterity = statVal
		case 3:
			char.CharacterStats.Vitality = statVal
		case 4:
			char.CharacterStats.UnusedStats = statVal
		case 5:
			char.CharacterStats.UnusedSkillPoints = statVal
		case 6:
			char.CharacterStats.CurrentHP = statVal / 256
		case 7:
			char.CharacterStats.MaxHP = statVal / 256
		case 8:
			char.CharacterStats.CurrentMana = statVal / 256
		case 9:
			char.CharacterStats.MaxMana = statVal / 256
		case 10:
			char.CharacterStats.CurrentStamina = statVal / 256
		case 11:
			char.CharacterStats.MaxStamina = statVal / 256
		case 12:
			char.CharacterStats.Level = statVal
		case 13:
			char.CharacterStats.Experience = statVal
		case 14:
			char.CharacterStats.Gold = statVal
		case 15:
			char.CharacterStats.StashedGold = statVal
		}
	}

	return nil
}

func parseSkills(bfr io.Reader, char *Character) error {

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
		return errors.New("Failed to find skill header, parsing cancelled")
	}

	skillOffset, ok := skillOffsetMap[uint(char.Header.Class)]
	if !ok {
		return fmt.Errorf("Unknown skill offset for class %d", char.Header.Class)
	}

	for i, allocatedPoints := range rawSkills.List {
		id := (i + skillOffset)
		s := skill{
			id:     id,
			points: int(allocatedPoints),
			name:   skillMap[id],
		}
		char.Skills = append(char.Skills, s)
	}

	return nil
}

func parseItems(bfr io.ByteReader, char *Character) error {

	// Make a buffer that can hold 4 bytes, which can hold the items header.
	buf := make([]byte, 4)

	_, err := io.ReadFull(bfr.(io.Reader), buf[:4])
	if err != nil {
		return err
	}

	itemHeaderData := itemHeader{}
	err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &itemHeaderData)

	if err != nil {
		return err
	}

	if string(itemHeaderData.Header[:]) != "JM" {
		return errors.New("Failed to find the items header")
	}

	fmt.Printf("Items count: %d\n", itemHeaderData.Count)

	// Read the list of items
	ibr := bitReader{r: bfr}

	var equippedItems []Item

	for i := 0; i < int(itemHeaderData.Count); i++ {
		var readBits int

		// offset: 0 "J"
		ibr.ReadBits64(8, false)
		readBits += 8

		// offset: 8, "M"
		ibr.ReadBits64(8, false)
		readBits += 8

		item := Item{}

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
		reverseBits(ibr.ReadBits64(2, true), 2)
		readBits += 2

		// offset 32
		item.IsEar = reverseBits(ibr.ReadBits64(1, true), 1)
		readBits++

		// offset 33
		item.StarterItem = reverseBits(ibr.ReadBits64(1, true), 1)
		readBits++

		// offset 34, unknown
		reverseBits(ibr.ReadBits64(3, true), 3)
		readBits += 3

		// offset 37, if it's a simple item, it only contains 111 bits data
		item.SimpleItem = reverseBits(ibr.ReadBits64(1, true), 1)
		readBits++

		// offset 38
		item.Ethereal = reverseBits(ibr.ReadBits64(1, true), 1)
		readBits++

		// offset 39, unknown
		reverseBits(ibr.ReadBits64(1, true), 1)
		readBits++

		// offset 40
		item.Personalized = reverseBits(ibr.ReadBits64(1, true), 1)
		readBits++

		// offset 41, unknown
		reverseBits(ibr.ReadBits64(1, true), 1)
		readBits++

		// offset 42
		item.GivenRuneword = reverseBits(ibr.ReadBits64(1, true), 1)
		readBits++

		// offset 43, unknown
		reverseBits(ibr.ReadBits64(15, true), 15)
		readBits += 15

		// offset 58
		item.LocationID = reverseBits(ibr.ReadBits64(3, true), 3)
		readBits += 3

		// offset 61
		item.EquippedID = reverseBits(ibr.ReadBits64(4, true), 4)
		readBits += 4

		// offset 65
		item.PositionY = reverseBits(ibr.ReadBits64(4, true), 4)
		readBits += 4

		// offset 69
		item.PositionX = reverseBits(ibr.ReadBits64(3, true), 3)
		readBits += 3

		// offset 72
		reverseBits(ibr.ReadBits64(1, true), 1)
		readBits++

		// offset 73, if item is neither equipped or in the belt
		// this tells us where it is.
		item.AltPositionID = reverseBits(ibr.ReadBits64(3, true), 3)
		readBits += 3

		// offset 76, item type, 4 chars, each 8 bit (not byte aligned)
		var itemType string
		for j := 0; j < 4; j++ {
			itemType += string(reverseBits(ibr.ReadBits64(8, true), 8))
		}

		item.Type = strings.Trim(itemType, " ")
		readBits += 32

		// offset 108
		// TODO: If sockets exist, read the items, they'll be 108 bit basic items * nrOfSockets
		item.NrOfItemsInSockets = reverseBits(ibr.ReadBits64(3, true), 3)
		readBits += 3

		// offset 111, item id is 8 chars, each 4 bit
		// TODO: Convert to hex, 4 bit each, should be 59BA3CAB
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

		case rare:
			// TODO: Parse rare bits.

		case unique:
			// TODO: Parse unique bits.

		case crafted:
			// TODO: Parse crafted bits.

		}

		// MARK: Runeword data
		// TODO: Parse 16 bits here if the item has a runeword.

		// MARK: Personalization data
		// TODO: Parse Personalization data here if the item is personalized.

		// MARK: Structure - all items have this part.

		// All items have this field between the personalization (if it exists)
		// and the item specific data
		// TODO: Should this be here?
		fmt.Printf("Bits read before timestamp: %d \n", readBits)
		item.Timestamp = reverseBits(ibr.ReadBits64(1, true), 1)
		readBits++

		typeID := item.getTypeID()
		fmt.Printf("Item type is %d\n", typeID)

		if typeID == armor {
			fmt.Printf("Bits read before armor: %d \n", readBits)
			// TODO: Figure out which items have 11 bits of defense rating data.
			// If the item is an armor, it will have this field of defense data.
			item.DefenseRating = reverseBits(ibr.ReadBits64(11, true), 11)
			fmt.Printf("Defense rating: %d \n", item.DefenseRating)
			readBits += 11
		}

		if typeID == armor || typeID == weapon {
			fmt.Printf("Bits read before durability: %d \n", readBits)
			item.MaxDurability = reverseBits(ibr.ReadBits64(8, true), 8)
			readBits += 8
			item.CurrentDurability = reverseBits(ibr.ReadBits64(8, true), 8)
			readBits += 8

			// If the item is an armor, it seems it has an extra bit here.
			if typeID == armor {
				reverseBits(ibr.ReadBits64(1, true), 1)
				readBits++
			}

			fmt.Printf("Max durability: %d \n", item.MaxDurability)
			fmt.Printf("Current durability: %d \n", item.CurrentDurability)

			// 1 extra bit here if it's a weapon
			if typeID == weapon {
				reverseBits(ibr.ReadBits64(1, true), 1)
				readBits++
			}
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

		// TODO: Find out if item is a tome, then read 5 bits of unknown data here
		// reverseBits(ibr.ReadBits64(5, true), 5)

		// If the item is part of a set, these bit will tell us how many lists
		// of magical properties follow the one regular magical property list.
		if item.Quality == partOfSet {
			item.SetItemLists = reverseBits(ibr.ReadBits64(5, true), 5)
			readBits += 5
		}

		fmt.Printf("Read bits until magic list: %d \n", readBits)

		// MARK: Time to parse 9 bit magical property ids followed by their n bit
		// length values.

		for {
			id := reverseBits(ibr.ReadBits64(9, true), 9)
			readBits += 9

			fmt.Printf("id before doing stuff: %d \n", id)

			if ibr.Err() != nil {
				return ibr.Err()
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

			var values []uint64
			for _, bitLength := range prop.Bits {

				val := reverseBits(ibr.ReadBits64(bitLength, true), bitLength)
				readBits += int(bitLength)
				fmt.Printf("found id: %d, reading bit size field: %d:, value is: %d\n", id, bitLength, val)

				if prop.Bias != 0 {
					val = val - prop.Bias
				}

				values = append(values, val)
			}

			attr := magicAttribute{
				ID:     id,
				Name:   prop.Name,
				Values: values,
			}

			item.MagicAttributes = append(item.MagicAttributes, attr)

			fmt.Printf("bits read after property id %d: %d \n", id, readBits)
		}

		if item.LocationID == equipped {
			equippedItems = append(equippedItems, item)
		}

		// If the item is not byte aligned, we'll have to byte align it before
		// reading the next item, so we'll simply queue the reader at the next
		// byte boundry.
		/*remainder := (8 - readBits%8)
		if remainder > 0 {
			reverseBits(ibr.ReadBits64(uint(remainder), true), uint(remainder))
		}*/

		fmt.Printf("Item index: %d\n%+v\n", i, item)
	}

	return nil
}

// Parse does stuff
func Parse(file io.Reader) {

	// Implements buffered reading, wraps io.Reader.
	bfr := bufio.NewReader(file)
	char := new(Character)

	_ = parseHeader(bfr, char)

	// Create a bit reader from the buffered reader to read the stats
	br := bitReader{r: bfr}

	_ = parseStats(br, char)

	_ = parseSkills(bfr, char)

	_ = parseItems(bfr, char)

	fmt.Printf("Character data:\n%+v\n", char)

	// Right now we've read n amount of bits, which means we're probably
	// not byte aligned, offset % 8 = remainder, and if remainder is not 0,
	// we need to read (8 - remainder) bits to reach the next byte boundry.
	// BitReader reads in 1 byte chunks, which means bfr is queued at
	// the next byte boundry already. We'll reuse the buf from before.

	/*
		// MARK: Items.



		//fmt.Printf("Item data:\n%+v\n", equippedItems)
	*/
}
