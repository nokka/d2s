package d2s

var socketablesWeapons = map[string][]MagicAttribute{
	// Runes
	ElRune:    {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{50}}, {ID: 89, Name: "+{0} to Light Radius", Values: []int64{1}}},
	EldRune:   {{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{75}}, {ID: 124, Name: "+{0} to Attack Rating against Undead", Values: []int64{50}}},
	TirRune:   {{ID: 138, Name: "+{0} to Mana After Each Kill", Values: []int64{2}}},
	NefRune:   {{ID: 81, Name: "Knockback", Values: []int64{1}}},
	EthRune:   {{ID: 116, Name: "{0}% Target Defense", Values: []int64{-25}}},
	IthRune:   {{ID: 22, Name: "+{0} to Maximum 1-handed damage", Values: []int64{9}}, {ID: 24, Name: "+{0} to Maximum 2-handed damage", Values: []int64{9}}},
	TalRune:   {{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{154, 154, 125}}},
	RalRune:   {{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{5, 30}}},
	OrtRune:   {{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 50}}},
	ThulRune:  {{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{3, 14, 75}}},
	AmnRune:   {{ID: 60, Name: "{0}% Life stolen per hit", Values: []int64{7}}},
	SolRune:   {{ID: 21, Name: "+{0} to Minimum 1-handed damage", Values: []int64{9}}, {ID: 23, Name: "+{0} to Minimum 2-handed damage", Values: []int64{9}}},
	ShaelRune: {{ID: 93, Name: "{0}% Increased Attack Speed", Values: []int64{20}}},
	DolRune:   {{ID: 112, Name: "Hit Causes Monsters to Flee {0}%", Values: []int64{25}}},
	HelRune:   {{ID: 91, Name: "Requirements {0}%", Values: []int64{-20}}},
	IoRune:    {{ID: 3, Name: "+{0} to Vitality", Values: []int64{10}}},
	LumRune:   {{ID: 1, Name: "+{0} to Energy", Values: []int64{10}}},
	KoRune:    {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{10}}},
	FalRune:   {{ID: 0, Name: "+{0} to Strength", Values: []int64{10}}},
	LemRune:   {{ID: 79, Name: "{0}% Extra Gold from Monsters", Values: []int64{75}}},
	PulRune:   {{ID: 121, Name: "+{0}% Damage to Demons", Values: []int64{75}}, {ID: 123, Name: "+{0} to Attack Rating against Demons", Values: []int64{100}}},
	UmRune:    {{ID: 135, Name: "{0}% Chance of Open Wounds", Values: []int64{25}}},
	MalRune:   {{ID: 117, Name: "Prevent Monster Heal", Values: []int64{1}}},
	IstRune:   {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{30}}},
	GulRune:   {{ID: 119, Name: "{0}% Bonus to Attack Rating", Values: []int64{20}}},
	VexRune:   {{ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{7}}},
	OhmRune:   {{ID: 17, Name: "+{0}% Enhanced Damage", Values: []int64{50, 50}}},
	LoRune:    {{ID: 141, Name: "{0}% Deadly Strike", Values: []int64{20}}},
	SurRune:   {{ID: 113, Name: "Hit Blinds Target +{0}", Values: []int64{20}}},
	BerRune:   {{ID: 136, Name: "{0}% Chance of Crushing Blow", Values: []int64{20}}},
	JahRune:   {{ID: 115, Name: "Ignore Target Defense", Values: []int64{1}}},
	ChamRune:  {{ID: 134, Name: "Freezes Target +{0}", Values: []int64{3}}},
	ZodRune:   {{ID: 152, Name: "Indestructible", Values: []int64{1}}},

	// Amethyst
	ChippedAmethyst:  {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{40}}},
	FlawedAmethyst:   {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{60}}},
	Amethyst:         {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{80}}},
	FlawlessAmethyst: {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{100}}},
	PerfectAmethyst:  {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{150}}},

	// Diamond
	ChippedDiamond:  {{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{28}}},
	FlawedDiamond:   {{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{34}}},
	Diamond:         {{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{44}}},
	FlawlessDiamond: {{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{54}}},
	PerfectDiamond:  {{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{68}}},

	// Emerald
	ChippedEmerald:  {{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{35, 35, 75}}},
	FlawedEmerald:   {{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{52, 52, 100}}},
	Emerald:         {{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{82, 82, 125}}},
	FlawlessEmerald: {{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{103, 103, 150}}},
	PerfectEmerald:  {{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{147, 147, 175}}},

	// Ruby
	ChippedRuby:  {{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{3, 4}}},
	FlawedRuby:   {{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{5, 8}}},
	Ruby:         {{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{8, 12}}},
	FlawlessRuby: {{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{10, 16}}},
	PerfectRuby:  {{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{15, 20}}},

	// Sapphire
	ChippedSapphire:  {{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{1, 3}}},
	FlawedSapphire:   {{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{3, 5}}},
	Sapphire:         {{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{4, 7}}},
	FlawlessSapphire: {{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{6, 10}}},
	PerfectSapphire:  {{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{10, 14}}},

	// Skull
	ChippedSkull:  {{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{2}}, {ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{1}}},
	FlawedSkull:   {{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{2}}, {ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{2}}},
	Skull:         {{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{3}}, {ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{2}}},
	FlawlessSkull: {{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{3}}, {ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{3}}},
	PerfectSkull:  {{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{4}}, {ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{3}}},

	// Topaz
	ChippedTopaz:  {{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 8}}},
	FlawedTopaz:   {{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 14}}},
	Topaz:         {{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 22}}},
	FlawlessTopaz: {{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 30}}},
	PerfectTopaz:  {{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 40}}},
}

var socketablesArmor = map[string][]MagicAttribute{
	ElRune:    {{ID: 31, Name: "+{0} Defense", Values: []int64{15}}, {ID: 89, Name: "+{0} to Light Radius", Values: []int64{1}}},
	EldRune:   {{ID: 154, Name: "{0}% Slower Stamina Drain", Values: []int64{15}}},
	TirRune:   {{ID: 138, Name: "+{0} to Mana After Each Kill", Values: []int64{2}}},
	NefRune:   {{ID: 32, Name: "+{0} vs. Missile", Values: []int64{30}}},
	EthRune:   {{ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{15}}},
	IthRune:   {{ID: 114, Name: "{0}% Damage Taken Goes to Mana", Values: []int64{15}}},
	TalRune:   {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{35}}},
	RalRune:   {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{35}}},
	OrtRune:   {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{35}}},
	ThulRune:  {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{35}}},
	AmnRune:   {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{14}}},
	SolRune:   {{ID: 34, Name: "Damage Reduced by {0}", Values: []int64{7}}},
	ShaelRune: {{ID: 99, Name: "{0}% Faster Hit Recovery", Values: []int64{20}}},
	DolRune:   {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{7}}},
	HelRune:   {{ID: 91, Name: "Requirements {0}%", Values: []int64{-15}}},
	IoRune:    {{ID: 3, Name: "+{0} to Vitality", Values: []int64{10}}},
	LumRune:   {{ID: 1, Name: "+{0} to Energy", Values: []int64{10}}},
	KoRune:    {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{10}}},
	FalRune:   {{ID: 0, Name: "+{0} to Strength", Values: []int64{10}}},
	LemRune:   {{ID: 79, Name: "{0}% Extra Gold from Monsters", Values: []int64{50}}},
	PulRune:   {{ID: 16, Name: "+{0}% Enhanced Defense", Values: []int64{30}}},
	UmRune:    {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{15}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{15}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{15}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{15}}},
	MalRune:   {{ID: 35, Name: "Magic Damage Reduced by {0}", Values: []int64{7}}},
	IstRune:   {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{25}}},
	GulRune:   {{ID: 46, Name: "+{0}% to Maximum Poison Resist", Values: []int64{5}}},
	VexRune:   {{ID: 40, Name: "+{0}% to Maximum Fire Resist", Values: []int64{5}}},
	OhmRune:   {{ID: 44, Name: "+{0}% to Maximum Cold Resist", Values: []int64{5}}},
	LoRune:    {{ID: 42, Name: "+{0}% to Maximum Lightning Resist", Values: []int64{5}}},
	SurRune:   {{ID: 77, Name: "Increase Maximum Mana {0}%", Values: []int64{5}}},
	BerRune:   {{ID: 36, Name: "Damage Reduced by {0}%", Values: []int64{8}}},
	JahRune:   {{ID: 76, Name: "Increase Maximum Life {0}%", Values: []int64{5}}},
	ChamRune:  {{ID: 153, Name: "Cannot Be Frozen", Values: []int64{1}}},
	ZodRune:   {{ID: 152, Name: "Indestructible", Values: []int64{1}}},

	// Amethyst
	ChippedAmethyst:  {{ID: 0, Name: "+{0} to Strength", Values: []int64{3}}},
	FlawedAmethyst:   {{ID: 0, Name: "+{0} to Strength", Values: []int64{4}}},
	Amethyst:         {{ID: 0, Name: "+{0} to Strength", Values: []int64{6}}},
	FlawlessAmethyst: {{ID: 0, Name: "+{0} to Strength", Values: []int64{8}}},
	PerfectAmethyst:  {{ID: 0, Name: "+{0} to Strength", Values: []int64{10}}},

	// Diamond
	ChippedDiamond:  {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{20}}},
	FlawedDiamond:   {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{40}}},
	Diamond:         {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{60}}},
	FlawlessDiamond: {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{80}}},
	PerfectDiamond:  {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{100}}},

	// Emerald
	ChippedEmerald:  {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{3}}},
	FlawedEmerald:   {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{4}}},
	Emerald:         {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{6}}},
	FlawlessEmerald: {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{8}}},
	PerfectEmerald:  {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{10}}},

	// Ruby
	ChippedRuby:  {{ID: 7, Name: "+{0} to Life", Values: []int64{10}}},
	FlawedRuby:   {{ID: 7, Name: "+{0} to Life", Values: []int64{17}}},
	Ruby:         {{ID: 7, Name: "+{0} to Life", Values: []int64{24}}},
	FlawlessRuby: {{ID: 7, Name: "+{0} to Life", Values: []int64{31}}},
	PerfectRuby:  {{ID: 7, Name: "+{0} to Life", Values: []int64{38}}},

	// Sapphire
	ChippedSapphire:  {{ID: 9, Name: "+{0} to Mana", Values: []int64{10}}},
	FlawedSapphire:   {{ID: 9, Name: "+{0} to Mana", Values: []int64{17}}},
	Sapphire:         {{ID: 9, Name: "+{0} to Mana", Values: []int64{24}}},
	FlawlessSapphire: {{ID: 9, Name: "+{0} to Mana", Values: []int64{31}}},
	PerfectSapphire:  {{ID: 9, Name: "+{0} to Mana", Values: []int64{38}}},

	// Skull
	ChippedSkull:  {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{2}}, {ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{8}}},
	FlawedSkull:   {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{3}}, {ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{8}}},
	Skull:         {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{3}}, {ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{12}}},
	FlawlessSkull: {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{4}}, {ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{12}}},
	PerfectSkull:  {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{5}}, {ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{19}}},

	// Topaz
	ChippedTopaz:  {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{9}}},
	FlawedTopaz:   {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{13}}},
	Topaz:         {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{16}}},
	FlawlessTopaz: {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{20}}},
	PerfectTopaz:  {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{24}}},
}

var socketablesShields = map[string][]MagicAttribute{
	ElRune:    {{ID: 31, Name: "+{0} Defense", Values: []int64{15}}, {ID: 89, Name: "+{0} to Light Radius", Values: []int64{1}}},
	EldRune:   {{ID: 20, Name: "+{0}% Increased chance of blocking", Values: []int64{7}}},
	TirRune:   {{ID: 138, Name: "+{0} to Mana After Each Kill", Values: []int64{2}}},
	NefRune:   {{ID: 32, Name: "+{0} vs. Missile", Values: []int64{30}}},
	EthRune:   {{ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{15}}},
	IthRune:   {{ID: 114, Name: "{0}% Damage Taken Goes to Mana", Values: []int64{15}}},
	TalRune:   {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{35}}},
	RalRune:   {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{35}}},
	OrtRune:   {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{35}}},
	ThulRune:  {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{35}}},
	AmnRune:   {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{14}}},
	SolRune:   {{ID: 34, Name: "Damage Reduced by {0}", Values: []int64{7}}},
	ShaelRune: {{ID: 102, Name: "{0}% Faster Block Rate", Values: []int64{20}}},
	DolRune:   {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{7}}},
	HelRune:   {{ID: 91, Name: "Requirements {0}%", Values: []int64{-15}}},
	IoRune:    {{ID: 3, Name: "+{0} to Vitality", Values: []int64{10}}},
	LumRune:   {{ID: 1, Name: "+{0} to Energy", Values: []int64{10}}},
	KoRune:    {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{10}}},
	FalRune:   {{ID: 0, Name: "+{0} to Strength", Values: []int64{10}}},
	LemRune:   {{ID: 79, Name: "{0}% Extra Gold from Monsters", Values: []int64{50}}},
	PulRune:   {{ID: 16, Name: "+{0}% Enhanced Defense", Values: []int64{30}}},
	UmRune:    {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{22}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{22}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{22}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{22}}},
	MalRune:   {{ID: 35, Name: "Magic Damage Reduced by {0}", Values: []int64{7}}},
	IstRune:   {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{25}}},
	GulRune:   {{ID: 46, Name: "+{0}% to Maximum Poison Resist", Values: []int64{5}}},
	VexRune:   {{ID: 40, Name: "+{0}% to Maximum Fire Resist", Values: []int64{5}}},
	OhmRune:   {{ID: 44, Name: "+{0}% to Maximum Cold Resist", Values: []int64{5}}},
	LoRune:    {{ID: 42, Name: "+{0}% to Maximum Lightning Resist", Values: []int64{5}}},
	SurRune:   {{ID: 9, Name: "+{0} to Mana", Values: []int64{50}}},
	BerRune:   {{ID: 36, Name: "Damage Reduced by {0}%", Values: []int64{8}}},
	JahRune:   {{ID: 7, Name: "+{0} to Life", Values: []int64{50}}},
	ChamRune:  {{ID: 153, Name: "Cannot Be Frozen", Values: []int64{1}}},
	ZodRune:   {{ID: 152, Name: "Indestructible", Values: []int64{1}}},

	// Amethyst
	ChippedAmethyst:  {{ID: 31, Name: "+{0} Defense", Values: []int64{8}}},
	FlawedAmethyst:   {{ID: 31, Name: "+{0} Defense", Values: []int64{12}}},
	Amethyst:         {{ID: 31, Name: "+{0} Defense", Values: []int64{18}}},
	FlawlessAmethyst: {{ID: 31, Name: "+{0} Defense", Values: []int64{24}}},
	PerfectAmethyst:  {{ID: 31, Name: "+{0} Defense", Values: []int64{30}}},

	// Diamond
	ChippedDiamond:  {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{6}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{6}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{6}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{6}}},
	FlawedDiamond:   {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{8}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{8}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{8}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{8}}},
	Diamond:         {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{11}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{11}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{11}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{11}}},
	FlawlessDiamond: {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{14}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{14}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{14}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{14}}},
	PerfectDiamond:  {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{19}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{19}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{19}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{19}}},

	// Emerald
	ChippedEmerald:  {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{12}}},
	FlawedEmerald:   {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{16}}},
	Emerald:         {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{22}}},
	FlawlessEmerald: {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{28}}},
	PerfectEmerald:  {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{40}}},

	// Ruby
	ChippedRuby:  {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{12}}},
	FlawedRuby:   {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{16}}},
	Ruby:         {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{22}}},
	FlawlessRuby: {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{28}}},
	PerfectRuby:  {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{40}}},

	// Sapphire
	ChippedSapphire:  {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{12}}},
	FlawedSapphire:   {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{16}}},
	Sapphire:         {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{22}}},
	FlawlessSapphire: {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{28}}},
	PerfectSapphire:  {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{40}}},

	// Skull
	ChippedSkull:  {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{4}}},
	FlawedSkull:   {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{8}}},
	Skull:         {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{12}}},
	FlawlessSkull: {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{16}}},
	PerfectSkull:  {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{20}}},

	// Topaz
	ChippedTopaz:  {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{12}}},
	FlawedTopaz:   {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{16}}},
	Topaz:         {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{22}}},
	FlawlessTopaz: {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{28}}},
	PerfectTopaz:  {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{40}}},
}
