package d2s

// quests struct will represent the characters in game progress of quests for a given difficulty.
type quests struct {
	// Set to 1 if you have been introduced to act I by Warriv.
	_ [2]byte

	// All 6 quests for act I.
	ActI questsActI `json:"act_i"`

	// These bytes are set to 1 if you have traveled to act II.
	_ [2]byte

	// Set to 1 if you have been introduced to act II by Jerhyn.
	_ [2]byte

	// All 6 quests for act II.
	ActII questsActII `json:"act_ii"`

	// These bytes are set to 1 if you have traveled to act III.
	_ [2]byte

	// Set to 1 if you have been introduced to act III by Hratli.
	_ [2]byte

	// All 6 qusts for act III.
	ActIII questsActIII `json:"act_iii"`

	// These bytes are set to 1 if you have traveled to act IV.
	_ [2]byte

	// Set to 1 if you have been introduced to act IV. (which you have, if you travelled to act IV)
	_ [2]byte

	// All 6 quests for act IV. Since act IV only have 3 quests, we have 3 quests (6 bytes) of empty data here.
	ActIV questsActIV `json:"act_iv"`

	// These bytes are set to 1 if you have traveled to act V.
	_ [2]byte

	// Seems to be set to 1 after completing Terror's End and talking to Cain in act IV.
	_ [2]byte

	// Unknown, seems to be some kind of padding.
	_ [4]byte

	// All 6 quests for act V.
	ActV questsActV `json:"act_v"`

	// Unknown, some kind of padding after all the quest data.
	_ [14]byte
}

type questsActI struct {
	DenOfEvil             quest `json:"den_of_evil"`
	SistersBurialGrounds  quest `json:"sisters_burial_grounds"`
	ToolsOfTheTrade       quest `json:"tools_of_the_trade"`
	TheSearchForCain      quest `json:"the_search_for_cain"`
	TheForgottenTower     quest `json:"the_forgotten_tower"`
	SistersToTheSlaughter quest `json:"sisters_to_the_slaughter"`
}

type questsActII struct {
	RadamentsLair    quest `json:"radaments_lair"`
	TheHoradricStaff quest `json:"the_horadric_staff"`
	TaintedSun       quest `json:"tainted_sun"`
	ArcaneSanctuary  quest `json:"arcane_sanctuary"`
	TheSummoner      quest `json:"the_summoner"`
	TheSevenTombs    quest `json:"the_seven_tombs"`
}

type questsActIII struct {
	LamEsensTome          quest `json:"lam_esens_tome"`
	KhalimsWill           quest `json:"khalims_will"`
	BladeOfTheOldReligion quest `json:"blade_of_the_old_religion"`
	TheGoldenBird         quest `json:"the_golden_bird"`
	TheBlackenedTemple    quest `json:"the_blackened_temple"`
	TheGuardian           quest `json:"the_guardian"`
}

type questsActIV struct {
	TheFallenAngel quest `json:"the_fallen_angel"`
	TerrorsEnd     quest `json:"terrors_end"`
	HellForge      quest `json:"hellforge"`
	_              [2]byte
	_              [2]byte
	_              [2]byte
}

type questsActV struct {
	SiegeOnHarrogath    quest       `json:"siege_on_harrogath"`
	RescueOnMountArreat quest       `json:"rescue_on_mount_arreat"`
	PrisonOfIce         prisonOfIce `json:"prison_of_ice"`
	BetrayalOfHarrogath quest       `json:"betrayal_of_harrogath"`
	RiteOfPassage       quest       `json:"rite_of_passage"`
	EveOfDestruction    quest       `json:"eve_of_destruction"`
}
