package d2s

import (
	"encoding/json"
	"fmt"
)

// Character represents all the d2s character data.
type Character struct {
	Header      header     `json:"header"`
	Attributes  attributes `json:"attributes"`
	Skills      []skill    `json:"skills"`
	Items       []item     `json:"items"`
	CorpseItems []item     `json:"corpse_items"`
	MercItems   []item     `json:"merc_items"`
	GolemItem   *item      `json:"golem_item"`
	IsDead      uint16     `json:"is_dead"`
}

func (h *header) MarshalJSON() ([]byte, error) {
	type Alias header

	leftSkill, ok := skillMap[int(h.LeftSkill)]
	if !ok {
		leftSkill = ""
	}

	rightSkill, ok := skillMap[int(h.RightSkill)]
	if !ok {
		rightSkill = ""
	}

	leftSwapSkill, ok := skillMap[int(h.LeftSwapSkill)]
	if !ok {
		leftSwapSkill = ""
	}

	rightSwapSkill, ok := skillMap[int(h.RightSwapSkill)]
	if !ok {
		rightSwapSkill = ""
	}

	var assignedSkills []string
	for _, skillID := range h.AssignedSkills {
		skill, ok := skillMap[int(skillID)]
		if ok {
			assignedSkills = append(assignedSkills, skill)
		}
	}

	return json.Marshal(&struct {
		Identifier     string         `json:"identifier"`
		CheckSum       string         `json:"checksum"`
		Name           string         `json:"name"`
		Status         readableStatus `json:"status"`
		Class          string         `json:"class"`
		LastPlayed     uint32         `json:"last_played"`
		LeftSkill      string         `json:"left_skill"`
		RightSkill     string         `json:"right_skill"`
		LeftSwapSkill  string         `json:"left_swap_skill"`
		RightSwapSkill string         `json:"right_swap_skill"`
		MercID         string         `json:"merc_id"`
		AssignedSkills []string       `json:"assigned_skills"`
		QuestsNormal   quests         `json:"quests_normal"`
		QuestsNm       quests         `json:"quests_nm"`
		QuestsHell     quests         `json:"quests_hell"`
		*Alias
	}{
		Identifier:     fmt.Sprintf("%x", h.Identifier),
		CheckSum:       fmt.Sprintf("%x", h.CheckSum),
		Name:           h.Name.String(),
		Status:         h.Status.Readable(),
		Class:          h.Class.String(),
		LastPlayed:     h.LastPlayed,
		LeftSkill:      leftSkill,
		RightSkill:     rightSkill,
		LeftSwapSkill:  leftSwapSkill,
		RightSwapSkill: rightSwapSkill,
		MercID:         fmt.Sprintf("%x", h.MercID),
		AssignedSkills: assignedSkills,
		QuestsNormal:   h.QuestsNormal,
		QuestsNm:       h.QuestsNm,
		QuestsHell:     h.QuestsHell,
		Alias:          (*Alias)(h),
	})
}

// Header determines the header data of a d2s file.
type header struct {
	Identifier        uint32      `json:"identifier"`
	Version           uint32      `json:"version"`
	FileSize          uint32      `json:"filesize"`
	CheckSum          uint32      `json:"checksum"`
	ActiveArms        uint32      `json:"active_arms"`
	Name              name        `json:"name"`
	Status            status      `json:"status"`
	Progression       progression `json:"progression"`
	_                 [2]byte
	Class             class `json:"class"`
	_                 [2]byte
	Level             byte `json:"level"`
	_                 [4]byte
	LastPlayed        uint32 `json:"last_played"`
	_                 [4]byte
	AssignedSkills    [16]uint32 `json:"assigned_skills"`
	LeftSkill         uint32     `json:"left_skill"`
	RightSkill        uint32     `json:"right_skill"`
	LeftSwapSkill     uint32     `json:"left_swap_skill"`
	RightSwapSkill    uint32     `json:"right_swap_skill"`
	_                 [32]byte
	CurrentDifficulty difficulty `json:"difficulty"`
	MapID             uint32     `json:"map_id"`
	_                 [2]byte
	DeadMerc          uint16 `json:"dead_merc"`
	MercID            uint32 `json:"merc_id"`
	MercNameID        uint16 `json:"merc_name_id"`
	MercType          uint16 `json:"merc_type"`
	MercExp           uint32 `json:"merc_experience"`
	_                 [144]byte
	QuestHeader       [4]byte `json:"-"`
	_                 [6]byte
	QuestsNormal      quests  `json:"quests_normal"`
	QuestsNm          quests  `json:"quests_nm"`
	QuestsHell        quests  `json:"quests_hell"`
	WaypointHeader    [2]byte `json:"-"`
	_                 [6]byte
	WaypointsNormal   [24]byte `json:"-"`
	WaypointsNm       [24]byte `json:"-"`
	WaypointsHell     [24]byte `json:"-"`
	WaypointTrailer   byte     `json:"-"`
	NPCHeader         [2]byte  `json:"-"`
	_                 byte
	NPCIntroNormal    [5]byte `json:"-"`
	_                 [3]byte
	NPCIntroNm        [5]byte `json:"-"`
	_                 [3]byte
	NPCIntroHell      [5]byte `json:"-"`
	_                 [3]byte
	NPCReturnNorm     [4]byte `json:"-"`
	_                 [4]byte
	NPCReturnNm       [4]byte `json:"-"`
	_                 [4]byte
	NPCReturnHell     [4]byte `json:"-"`
	_                 [4]byte
	StatHeader        [2]byte `json:"-"`
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
