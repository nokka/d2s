package d2s

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

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
	Name              [16]byte   // : 20 16 bytes
	Status            byte       // : 36 4 bytes
	Progression       byte       // : 37 1 bytes
	_                 [2]byte    // : 38 2 bytes
	Class             byte       // : 40 1 bytes
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
	CurrentDifficulty [3]byte    // : 168 3 bytes
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
func Parse(character io.Reader) {

	// Implements buffered reading, wraps io.Reader.
	bfr := bufio.NewReader(character)

	// MARK: Header

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

	// MARK: Skills

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

	// MARK: Items

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

	// Unaligned bit reading

	ibr := bitReader{r: bfr}

	// offset: 0 "J"
	j := ibr.ReadBits64(8, false)

	// offset: 8, "M"
	m := ibr.ReadBits64(8, false)

	// offset: 16, unknown
	ibr.ReadBits64(4, false)

	// offset: 20
	isIdentified := ibr.ReadBits64(1, true)

	// offset: 21, unknown
	ibr.ReadBits64(6, false)

	// offset: 27
	isSocketed := ibr.ReadBits64(1, false)

	// offset 28, unknown
	ibr.ReadBits64(1, false)

	// offset 29
	newlyPickedUp := ibr.ReadBits64(1, false)

	// offset 30, unknown
	ibr.ReadBits64(2, false)

	// offset 32
	playerEar := ibr.ReadBits64(1, false)

	// offset 33
	starterItem := ibr.ReadBits64(1, false)

	// offset 34, unknown
	ibr.ReadBits64(3, false)

	// offset 37, if it's a simple item, it only contains 111 bits data
	simpleItem := ibr.ReadBits64(1, false)

	// offset 38
	ethereal := ibr.ReadBits64(1, false)

	// offset 39, unknown
	ibr.ReadBits64(1, false)

	// offset 40
	personalized := ibr.ReadBits64(1, false)

	// offset 41, unknown
	ibr.ReadBits64(1, false)

	// offset 42
	runeword := ibr.ReadBits64(1, false)

	// offset 43, unknown
	ibr.ReadBits64(15, false)

	// offset 58
	positionID := ibr.ReadBits64(3, false)

	// offset 61
	slotID := ibr.ReadBits64(4, false)

	// offset 65
	inventoryX := ibr.ReadBits64(4, false)

	// offset 69
	inventoryY := ibr.ReadBits64(3, false)

	// offset 72
	ibr.ReadBits64(1, false)

	// offset 73, if item is neither equipped or in the belt
	// this tells us where it is.
	altPositionID1 := ibr.ReadBits64(1, false)
	altPositionID2 := ibr.ReadBits64(1, false)
	altPositionID3 := ibr.ReadBits64(1, false)

	// offset 76, item type, 4 chars, each 8 bit (not byte aligned)
	itemType1 := ibr.ReadBits64(8, false)
	itemType2 := ibr.ReadBits64(8, false)
	itemType3 := ibr.ReadBits64(8, false)

	fmt.Printf("Header[0]: %s\n", string(j))
	fmt.Printf("Header[1]: %s\n", string(m))
	fmt.Printf("Is identified: %d\n", isIdentified)
	fmt.Printf("Is socketed: %d\n", isSocketed)
	fmt.Printf("Newly picked up: %d\n", newlyPickedUp)
	fmt.Printf("Is player ear??: %d\n", playerEar)
	fmt.Printf("Is starter item: %d\n", starterItem)
	fmt.Printf("Simple item: %d\n", simpleItem)
	fmt.Printf("Ethereal: %d\n", ethereal)
	fmt.Printf("Personalized: %d\n", personalized)
	fmt.Printf("Runeword: %d\n", runeword)
	fmt.Printf("Position ID: %d\n", positionID)
	fmt.Printf("Slot ID: %d\n", slotID)
	fmt.Printf("Inventory X: %d\n", inventoryX)
	fmt.Printf("Inventory Y: %d\n", inventoryY)

	fmt.Printf("Alt position ID1: %d\n", altPositionID1)
	fmt.Printf("Alt position ID2: %d\n", altPositionID2)
	fmt.Printf("Alt position ID3: %d\n", altPositionID3)

	fmt.Printf("Item type ID1: %s\n", string(itemType1))
	fmt.Printf("Item type ID2: %s\n", string(itemType2))
	fmt.Printf("Item type ID3: %s\n", string(itemType3))
}
