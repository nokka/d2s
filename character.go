package d2s

import (
	"encoding/json"
	"fmt"
)

// Character represents all the d2s character data.
type character struct {
	header `json:"header"`
	attributes
	skills      []skill
	items       []Item
	corpseItems []Item
	mercItems   []Item
	golemItem   Item
}

func (h *header) MarshalJSON() ([]byte, error) {
	type Alias header
	return json.Marshal(&struct {
		Identifier string `json:"identifier"`
		CheckSum   string `json:"checksum"`
		Name       string `json:"name"`
		Class      string `json:"class"`
		*Alias
	}{
		Identifier: fmt.Sprintf("%x", h.Identifier),
		CheckSum:   fmt.Sprintf("%x", h.CheckSum),
		Name:       h.Name.String(),
		Class:      h.Class.String(),
		Alias:      (*Alias)(h),
	})
}

// Header determines the header data of a d2s file.
type header struct {
	Identifier        uint32 `json:"identifier"`
	Version           uint32 `json:"version"`
	FileSize          uint32 `json:"filesize"`
	CheckSum          uint32 `json:"checksum"`
	ActiveArms        uint32 `json:"active_arms"`
	Name              name   `json:"name"`
	Status            byte
	Progression       byte
	_                 [2]byte
	Class             class `json:"class"`
	_                 [2]byte
	Level             byte
	_                 [4]byte
	LastPlayed        uint32
	_                 [4]byte
	AssignedSkills    [16]uint32
	LeftSkill         uint32
	RightSkill        uint32
	LeftSwapSkill     uint32
	RightSwapSkill    uint32
	_                 [32]byte
	CurrentDifficulty difficulty
	MapID             uint32
	_                 [2]byte
	DeadMerc          uint16
	MercID            uint32
	MercNameID        uint16
	MercType          uint16
	MercExp           uint32
	_                 [144]byte
	QuestHeader       [4]byte
	_                 [6]byte
	QuestsNormal      [96]byte
	QuestsNm          [96]byte
	QuestsHell        [96]byte
	WaypointHeader    [2]byte
	_                 [6]byte
	WaypointsNormal   [24]byte
	WaypointsNm       [24]byte
	WaypointsHell     [24]byte
	WaypointTrailer   byte
	NPCHeader         [2]byte
	_                 byte
	NPCIntroNormal    [5]byte
	_                 [3]byte
	NPCIntroNm        [5]byte
	_                 [3]byte
	NPCIntroHell      [5]byte
	_                 [3]byte
	NPCReturnNorm     [4]byte
	_                 [4]byte
	NPCReturnNm       [4]byte
	_                 [4]byte
	NPCReturnHell     [4]byte
	_                 [4]byte
	StatHeader        [2]byte
}

type skillData struct {
	Header [2]byte
	List   [30]byte
}

type itemData struct {
	Header [2]byte
	Count  uint16
}

type corpseData struct {
	Header  [2]byte
	Count   uint16
	Unknown [12]byte
}

type golemData struct {
	Header   [2]byte
	HasGolem byte
}
