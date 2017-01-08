package d2s

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"strings"
)

// Character is the exported character
type Character struct {
	Name string
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

// Skills holds a list reference on allocated skills.
type Skills struct {
	Header [2]byte
	List   [30]byte
}

// ItemHeader determines the header data of a d2s file.
type ItemHeader struct {
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

// Parse does stuff
func Parse(character io.Reader) Character {

	// Implements buffered reading, wraps io.Reader.
	bfr := bufio.NewReader(character)

	// MARK: Header.

	// Make a buffer that can hold 767 bytes, which can hold the entire header.
	// We'll reuse this buffer through out to avoid another alloc.
	buf := make([]byte, 767)

	_, err := io.ReadFull(bfr, buf)
	if err != nil {
		log.Fatal("Error reading from file:", err)
	}

	header := Header{}
	err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &header)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	fmt.Printf("Parsed data:\n%+v\n", header)

	// MARK: Stats

	// Create a bit reader from the buffered reader to read the stats
	// by 9 bit stat ids and n bit stat values, also bit reversed... twice.
	br := bitReader{r: bfr}

	characterStats := CharacterStats{}

	for {

		// 9 bit stat id, bit reversed twice.
		id := reverseBits(br.ReadBits64(9, true), 9)

		if br.Err() != nil {
			log.Fatal(br.Err())
		}

		// If all 9 bits are set, we've hit the end of the stats section
		//  at 0x1ff and exit the loop.
		if id == 0x1ff {
			break
		}

		// The stat value bit length, so we'll know how many bits to read next.
		bitLength, ok := statsBitMap[id]
		if !ok {
			log.Fatalf("Unknown stat id: %d", id)
		}

		// The stat value, bit reversed, twice.
		statVal := reverseBits(br.ReadBits64(bitLength, true), bitLength)
		if br.Err() != nil {
			log.Fatal(br.Err())
		}

		switch id {
		case 0:
			characterStats.Strength = statVal
		case 1:
			characterStats.Energy = statVal
		case 2:
			characterStats.Dexterity = statVal
		case 3:
			characterStats.Vitality = statVal
		case 4:
			characterStats.UnusedStats = statVal
		case 5:
			characterStats.UnusedSkillPoints = statVal
		case 6:
			characterStats.CurrentHP = statVal / 256
		case 7:
			characterStats.MaxHP = statVal / 256
		case 8:
			characterStats.CurrentMana = statVal / 256
		case 9:
			characterStats.MaxMana = statVal / 256
		case 10:
			characterStats.CurrentStamina = statVal / 256
		case 11:
			characterStats.MaxStamina = statVal / 256
		case 12:
			characterStats.Level = statVal
		case 13:
			characterStats.Experience = statVal
		case 14:
			characterStats.Gold = statVal
		case 15:
			characterStats.StashedGold = statVal
		}
	}

	fmt.Printf("Parsed data:\n%+v\n", characterStats)

	// MARK: Skills.

	// Right now we've read n amount of bits, which means we're probably
	// not byte aligned, offset % 8 = remainder, and if remainder is not 0,
	// we need to read (8 - remainder) bits to reach the next byte boundry.
	// BitReader reads in 1 byte chunks, which means bfr is queued at
	// the next byte boundry already. We'll reuse the buf from before.
	_, err = io.ReadFull(bfr, buf[:32])
	if err != nil {
		log.Fatal("Error reading from file:", err)
	}

	skills := Skills{}
	err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &skills)

	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	fmt.Printf("Parsed data:\n%+v\n", skills)

	skillOffset, ok := skillOffsetMap[uint(header.Class)]
	if !ok {
		log.Fatalf("Unknown skill offset for class %d", header.Class)
	}

	for i, allocatedPoints := range skills.List {
		fmt.Printf("%s: %d \n", skillMap[i+skillOffset], allocatedPoints)
	}

	// MARK: Items.

	_, err = io.ReadFull(bfr, buf[:4])
	if err != nil {
		log.Fatal("Error reading from file:", err)
	}

	items := ItemHeader{}
	err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &items)

	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	if string(items.Header[:]) != "JM" {
		log.Fatal("Failed to find the items header")
	}

	fmt.Printf("Item section header: %s\n", string(items.Header[:]))
	fmt.Printf("Items count: %d\n", items.Count)

	// TODO: Refactor into for loop where i < items.Count

	ibr := bitReader{r: bfr}

	// offset: 0 "J"
	ibr.ReadBits64(8, false)

	// offset: 8, "M"
	ibr.ReadBits64(8, false)

	item := Item{}

	// offset: 16, unknown
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
	ibr.ReadBits64(2, true)

	// offset 32
	item.IsEar = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 33
	item.StarterItem = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 34, unknown
	ibr.ReadBits64(3, true)

	// offset 37, if it's a simple item, it only contains 111 bits data
	item.SimpleItem = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 38
	item.Ethereal = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 39, unknown
	ibr.ReadBits64(1, true)

	// offset 40
	item.Personalized = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 41, unknown
	ibr.ReadBits64(1, true)

	// offset 42
	item.GivenRuneword = reverseBits(ibr.ReadBits64(1, true), 1)

	// offset 43, unknown
	ibr.ReadBits64(15, true)

	// offset 58
	item.LocationID = reverseBits(ibr.ReadBits64(3, true), 3)

	// offset 61
	item.EquippedID = reverseBits(ibr.ReadBits64(4, true), 4)

	// offset 65
	item.PositionY = reverseBits(ibr.ReadBits64(4, true), 4)

	// offset 69
	item.PositionX = reverseBits(ibr.ReadBits64(3, true), 3)

	// offset 72
	ibr.ReadBits64(1, true)

	// offset 73, if item is neither equipped or in the belt
	// this tells us where it is.
	item.AltPositionID = reverseBits(ibr.ReadBits64(3, true), 3)

	// offset 76, item type, 4 chars, each 8 bit (not byte aligned)
	var itemType string
	for i := 0; i < 4; i++ {
		itemType += string(reverseBits(ibr.ReadBits64(8, true), 8))
	}

	item.Type = strings.Trim(itemType, " ")

	// offset 108
	// TODO: If sockets exist, read the items, they'll be 108 bit basic items * nrOfSockets
	item.NrOfItemsInSockets = reverseBits(ibr.ReadBits64(3, true), 3)

	// offset 111, item id is 8 chars, each 4 bit
	// TODO: Convert to hex, 4 bit each, should be 59BA3CAB
	item.ID = reverseBits(ibr.ReadBits64(32, true), 32)

	// offset 143
	item.Level = reverseBits(ibr.ReadBits64(7, true), 7)

	// offset 150
	item.Quality = reverseBits(ibr.ReadBits64(4, true), 4)

	// If this is true, it means the item has more than one picture associated
	// with it.
	item.MultiplePictures = reverseBits(ibr.ReadBits64(1, true), 1)

	if item.MultiplePictures == 1 {
		// The next 3 bits contain the picture ID.
		item.PictureID = reverseBits(ibr.ReadBits64(3, true), 3)
	}

	// If this is true, it means the item is class specific.
	item.ClassSpecific = reverseBits(ibr.ReadBits64(1, true), 1)

	// If the item is class specific, the next 11 bits will
	// contain the class specific data.
	if item.ClassSpecific == 1 {
		// TODO: Parse this into something useful
		reverseBits(ibr.ReadBits64(11, true), 11)
	}

	// MARK: Quality based data.

	switch item.Quality {
	// Low quality
	case 1:
		item.LowQualityID = reverseBits(ibr.ReadBits64(3, true), 3)

	// Normal item
	case 2:
		// No extra data present

	// Superior white item
	case 3:
		// TODO: Figure out what these 3 bits are on a superior item
		reverseBits(ibr.ReadBits64(3, true), 3)

	// Magical item
	case 4:
		item := MagicalItem{Item: item}
		item.MagicPrefix = reverseBits(ibr.ReadBits64(11, true), 11)
		item.MagicSuffix = reverseBits(ibr.ReadBits64(11, true), 11)
	}

	// MARK: Structure - all items have this part.

	// All items have this field between the personalization (if it exists)
	// and the item specific data
	item.StructureHeader = reverseBits(ibr.ReadBits64(1, true), 1)

	// TODO: Make an item type mapper and determine type from here on out.
	// If the item is an armor, it will have this field of defense data.
	//item.DefenseRating = reverseBits(ibr.ReadBits64(10, true), 10)

	// TODO: Make an item type mapper and determine type from here on out.
	// If item is an armor or weapon it will have 8x2 bits of durability data.
	item.MaxDurability = reverseBits(ibr.ReadBits64(8, true), 8)
	item.CurrentDurability = reverseBits(ibr.ReadBits64(8, true), 8)

	// If the item is socketed, it will contain 4 bits of data which are the nr
	// of total sockets the item have, regardless of how many are occupied by
	// an item.
	if item.Socketed == 1 {
		item.TotalNrOfSockets = reverseBits(ibr.ReadBits64(4, true), 4)
	}

	// TODO: Find out if item is a tome, then read 5 bits of unknown data here
	// reverseBits(ibr.ReadBits64(5, true), 5)

	// If the item is a stacked item, e.g. a javelin or something, these 9
	// bits will contain the quantity.
	item.Quantity = reverseBits(ibr.ReadBits64(9, true), 9)

	fmt.Printf("Item data:\n%+v\n", item)

	// MARK: Compose to the exposed interface.
	c := Character{
		Name: header.Name.String(),
	}

	return c
}
