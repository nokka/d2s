package d2s

var socketablesWeapons = map[string][]magicAttribute{

	// Runes

	"r01": []magicAttribute{
		magicAttribute{ID: 19, Name: "+{0} to Attack rating", Values: []int64{50}},
		magicAttribute{ID: 89, Name: "+{0} to Light Radius", Values: []int64{1}},
	},

	"r02": []magicAttribute{
		magicAttribute{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{75}},
		magicAttribute{ID: 124, Name: "+{0} to Attack Rating against Undead", Values: []int64{50}},
	},
	"r03": []magicAttribute{magicAttribute{ID: 138, Name: "+{0} to Mana After Each Kill", Values: []int64{2}}},
	"r04": []magicAttribute{magicAttribute{ID: 81, Name: "Knockback", Values: []int64{1}}},
	"r05": []magicAttribute{magicAttribute{ID: 116, Name: "{0}% Target Defense", Values: []int64{-25}}},

	"r06": []magicAttribute{
		magicAttribute{ID: 22, Name: "+{0} to Maximum 1-handed damage", Values: []int64{9}},
		magicAttribute{ID: 24, Name: "+{0} to Maximum 2-handed damage", Values: []int64{9}},
	},
	"r07": []magicAttribute{magicAttribute{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{154, 154, 125}}},
	"r08": []magicAttribute{magicAttribute{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{5, 30}}},
	"r09": []magicAttribute{magicAttribute{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 50}}},
	"r10": []magicAttribute{magicAttribute{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{3, 14, 75}}},
	"r11": []magicAttribute{magicAttribute{ID: 60, Name: "{0}% Life stolen per hit", Values: []int64{7}}},
	"r12": []magicAttribute{
		magicAttribute{ID: 21, Name: "+{0} to Minimum 1-handed damage", Values: []int64{9}},
		magicAttribute{ID: 23, Name: "+{0} to Minimum 2-handed damage", Values: []int64{9}},
	},
	"r13": []magicAttribute{magicAttribute{ID: 93, Name: "{0}% Increased Attack Speed", Values: []int64{20}}},
	"r14": []magicAttribute{magicAttribute{ID: 112, Name: "Hit Causes Monsters to Flee {0}%", Values: []int64{25}}},
	"r15": []magicAttribute{magicAttribute{ID: 91, Name: "Requirements {0}%", Values: []int64{-20}}},
	"r16": []magicAttribute{magicAttribute{ID: 3, Name: "+{0} to Vitality", Values: []int64{10}}},
	"r17": []magicAttribute{magicAttribute{ID: 1, Name: "+{0} to Energy", Values: []int64{10}}},
	"r18": []magicAttribute{magicAttribute{ID: 2, Name: "+{0} to Dexterity", Values: []int64{10}}},
	"r19": []magicAttribute{magicAttribute{ID: 0, Name: "+{0} to Strength", Values: []int64{10}}},
	"r20": []magicAttribute{magicAttribute{ID: 79, Name: "{0}% Extra Gold from Monsters", Values: []int64{75}}},
	"r21": []magicAttribute{
		magicAttribute{ID: 121, Name: "+{0}% Damage to Demons", Values: []int64{75}},
		magicAttribute{ID: 123, Name: "+{0} to Attack Rating against Demons", Values: []int64{100}},
	},
	"r22": []magicAttribute{magicAttribute{ID: 135, Name: "{0}% Chance of Open Wounds", Values: []int64{25}}},
	"r23": []magicAttribute{magicAttribute{ID: 117, Name: "Prevent Monster Heal", Values: []int64{1}}},
	"r24": []magicAttribute{magicAttribute{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{30}}},
	"r25": []magicAttribute{magicAttribute{ID: 119, Name: "{0}% Bonus to Attack Rating", Values: []int64{20}}},
	"r26": []magicAttribute{magicAttribute{ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{7}}},
	"r27": []magicAttribute{magicAttribute{ID: 17, Name: "+{0}% Enhanced Damage", Values: []int64{50, 50}}},
	"r28": []magicAttribute{magicAttribute{ID: 141, Name: "{0}% Deadly Strike", Values: []int64{20}}},
	"r29": []magicAttribute{magicAttribute{ID: 113, Name: "Hit Blinds Target +{0}", Values: []int64{20}}},
	"r30": []magicAttribute{magicAttribute{ID: 136, Name: "{0}% Chance of Crushing Blow", Values: []int64{20}}},
	"r31": []magicAttribute{magicAttribute{ID: 115, Name: "Ignore Target Defense", Values: []int64{1}}},
	"r32": []magicAttribute{magicAttribute{ID: 134, Name: "Freezes Target +{0}", Values: []int64{3}}},
	"r33": []magicAttribute{magicAttribute{ID: 152, Name: "Indestructible", Values: []int64{1}}},

	// Amethyst
	"gcv": []magicAttribute{magicAttribute{ID: 19, Name: "+{0} to Attack rating", Values: []int64{40}}},
	"gfv": []magicAttribute{magicAttribute{ID: 19, Name: "+{0} to Attack rating", Values: []int64{60}}},
	"gsv": []magicAttribute{magicAttribute{ID: 19, Name: "+{0} to Attack rating", Values: []int64{80}}},
	"gzv": []magicAttribute{magicAttribute{ID: 19, Name: "+{0} to Attack rating", Values: []int64{100}}},
	"gpv": []magicAttribute{magicAttribute{ID: 19, Name: "+{0} to Attack rating", Values: []int64{150}}},

	// Diamond
	"gcw": []magicAttribute{magicAttribute{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{28}}},
	"gfw": []magicAttribute{magicAttribute{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{34}}},
	"gsw": []magicAttribute{magicAttribute{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{44}}},
	"glw": []magicAttribute{magicAttribute{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{54}}},
	"gpw": []magicAttribute{magicAttribute{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{68}}},

	// Emerald
	"gcg": []magicAttribute{magicAttribute{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{35, 35, 75}}},
	"gfg": []magicAttribute{magicAttribute{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{52, 52, 100}}},
	"gsg": []magicAttribute{magicAttribute{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{82, 82, 125}}},
	"glg": []magicAttribute{magicAttribute{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{103, 103, 150}}},
	"gpg": []magicAttribute{magicAttribute{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{147, 147, 175}}},

	// Ruby
	"gcr": []magicAttribute{magicAttribute{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{3, 4}}},
	"gfr": []magicAttribute{magicAttribute{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{5, 8}}},
	"gsr": []magicAttribute{magicAttribute{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{8, 12}}},
	"glr": []magicAttribute{magicAttribute{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{10, 16}}},
	"gpr": []magicAttribute{magicAttribute{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{15, 20}}},

	// Sapphire
	"gcb": []magicAttribute{magicAttribute{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{1, 3}}},
	"gfb": []magicAttribute{magicAttribute{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{3, 5}}},
	"gsb": []magicAttribute{magicAttribute{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{4, 7}}},
	"glb": []magicAttribute{magicAttribute{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{6, 10}}},
	"gpb": []magicAttribute{magicAttribute{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{10, 14}}},

	// Skull
	"skc": []magicAttribute{
		magicAttribute{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{2}},
		magicAttribute{ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{1}},
	},

	"skf": []magicAttribute{
		magicAttribute{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{2}},
		magicAttribute{ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{2}},
	},

	"sku": []magicAttribute{
		magicAttribute{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{3}},
		magicAttribute{ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{2}},
	},

	"skl": []magicAttribute{
		magicAttribute{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{3}},
		magicAttribute{ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{3}},
	},

	"skz": []magicAttribute{
		magicAttribute{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{4}},
		magicAttribute{ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{3}},
	},

	// Topaz
	"gcy": []magicAttribute{magicAttribute{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 8}}},
	"gfy": []magicAttribute{magicAttribute{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 14}}},
	"gsy": []magicAttribute{magicAttribute{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 22}}},
	"gly": []magicAttribute{magicAttribute{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 30}}},
	"gpy": []magicAttribute{magicAttribute{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 40}}},
}

var socketablesArmor = map[string][]magicAttribute{
	"r01": []magicAttribute{
		magicAttribute{ID: 31, Name: "+{0} Defense", Values: []int64{15}},
		magicAttribute{ID: 89, Name: "+{0} to Light Radius", Values: []int64{1}},
	},
	"r02": []magicAttribute{magicAttribute{ID: 154, Name: "{0}% Slower Stamina Drain", Values: []int64{15}}},
	"r03": []magicAttribute{magicAttribute{ID: 138, Name: "+{0} to Mana After Each Kill", Values: []int64{2}}},
	"r04": []magicAttribute{magicAttribute{ID: 32, Name: "+{0} vs. Missile", Values: []int64{30}}},
	"r05": []magicAttribute{magicAttribute{ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{15}}},
	"r06": []magicAttribute{magicAttribute{ID: 114, Name: "{0}% Damage Taken Goes to Mana", Values: []int64{15}}},
	"r07": []magicAttribute{magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{35}}},
	"r08": []magicAttribute{magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{35}}},
	"r09": []magicAttribute{magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{35}}},
	"r10": []magicAttribute{magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{35}}},
	"r11": []magicAttribute{magicAttribute{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{14}}},
	"r12": []magicAttribute{magicAttribute{ID: 34, Name: "Damage Reduced by {0}", Values: []int64{7}}},
	"r13": []magicAttribute{magicAttribute{ID: 99, Name: "{0}% Faster Hit Recovery", Values: []int64{20}}},
	"r14": []magicAttribute{magicAttribute{ID: 74, Name: "Replenish Life +{0}", Values: []int64{7}}},
	"r15": []magicAttribute{magicAttribute{ID: 91, Name: "Requirements {0}%", Values: []int64{-15}}},
	"r16": []magicAttribute{magicAttribute{ID: 3, Name: "+{0} to Vitality", Values: []int64{10}}},
	"r17": []magicAttribute{magicAttribute{ID: 1, Name: "+{0} to Energy", Values: []int64{10}}},
	"r18": []magicAttribute{magicAttribute{ID: 2, Name: "+{0} to Dexterity", Values: []int64{10}}},
	"r19": []magicAttribute{magicAttribute{ID: 0, Name: "+{0} to Strength", Values: []int64{10}}},
	"r20": []magicAttribute{magicAttribute{ID: 79, Name: "{0}% Extra Gold from Monsters", Values: []int64{50}}},
	"r21": []magicAttribute{magicAttribute{ID: 16, Name: "+{0}% Enhanced Defense", Values: []int64{30}}},
	"r22": []magicAttribute{
		magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{15}},
		magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{15}},
		magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{15}},
		magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{15}},
	},
	"r23": []magicAttribute{magicAttribute{ID: 35, Name: "Magic Damage Reduced by {0}", Values: []int64{7}}},
	"r24": []magicAttribute{magicAttribute{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{25}}},
	"r25": []magicAttribute{magicAttribute{ID: 46, Name: "+{0}% to Maximum Poison Resist", Values: []int64{5}}},
	"r26": []magicAttribute{magicAttribute{ID: 40, Name: "+{0}% to Maximum Fire Resist", Values: []int64{5}}},
	"r27": []magicAttribute{magicAttribute{ID: 44, Name: "+{0}% to Maximum Cold Resist", Values: []int64{5}}},
	"r28": []magicAttribute{magicAttribute{ID: 42, Name: "+{0}% to Maximum Lightning Resist", Values: []int64{5}}},
	"r29": []magicAttribute{magicAttribute{ID: 77, Name: "Increase Maximum Mana {0}%", Values: []int64{5}}},
	"r30": []magicAttribute{magicAttribute{ID: 36, Name: "Damage Reduced by {0}%", Values: []int64{8}}},
	"r31": []magicAttribute{magicAttribute{ID: 76, Name: "Increase Maximum Life {0}%", Values: []int64{5}}},
	"r32": []magicAttribute{magicAttribute{ID: 153, Name: "Cannot Be Frozen", Values: []int64{1}}},
	"r33": []magicAttribute{magicAttribute{ID: 152, Name: "Indestructible", Values: []int64{1}}},

	// Amethyst
	"gcv": []magicAttribute{magicAttribute{ID: 0, Name: "+{0} to Strength", Values: []int64{3}}},
	"gfv": []magicAttribute{magicAttribute{ID: 0, Name: "+{0} to Strength", Values: []int64{4}}},
	"gsv": []magicAttribute{magicAttribute{ID: 0, Name: "+{0} to Strength", Values: []int64{6}}},
	"gzv": []magicAttribute{magicAttribute{ID: 0, Name: "+{0} to Strength", Values: []int64{8}}},
	"gpv": []magicAttribute{magicAttribute{ID: 0, Name: "+{0} to Strength", Values: []int64{10}}},

	// Diamond
	"gcw": []magicAttribute{magicAttribute{ID: 19, Name: "+{0} to Attack rating", Values: []int64{20}}},
	"gfw": []magicAttribute{magicAttribute{ID: 19, Name: "+{0} to Attack rating", Values: []int64{40}}},
	"gsw": []magicAttribute{magicAttribute{ID: 19, Name: "+{0} to Attack rating", Values: []int64{60}}},
	"glw": []magicAttribute{magicAttribute{ID: 19, Name: "+{0} to Attack rating", Values: []int64{80}}},
	"gpw": []magicAttribute{magicAttribute{ID: 19, Name: "+{0} to Attack rating", Values: []int64{100}}},

	// Emerald
	"gcg": []magicAttribute{magicAttribute{ID: 2, Name: "+{0} to Dexterity", Values: []int64{3}}},
	"gfg": []magicAttribute{magicAttribute{ID: 2, Name: "+{0} to Dexterity", Values: []int64{4}}},
	"gsg": []magicAttribute{magicAttribute{ID: 2, Name: "+{0} to Dexterity", Values: []int64{6}}},
	"glg": []magicAttribute{magicAttribute{ID: 2, Name: "+{0} to Dexterity", Values: []int64{8}}},
	"gpg": []magicAttribute{magicAttribute{ID: 2, Name: "+{0} to Dexterity", Values: []int64{10}}},

	// Ruby
	"gcr": []magicAttribute{magicAttribute{ID: 7, Name: "+{0} to Life", Values: []int64{10}}},
	"gfr": []magicAttribute{magicAttribute{ID: 7, Name: "+{0} to Life", Values: []int64{17}}},
	"gsr": []magicAttribute{magicAttribute{ID: 7, Name: "+{0} to Life", Values: []int64{24}}},
	"glr": []magicAttribute{magicAttribute{ID: 7, Name: "+{0} to Life", Values: []int64{31}}},
	"gpr": []magicAttribute{magicAttribute{ID: 7, Name: "+{0} to Life", Values: []int64{38}}},

	// Sapphire
	"gcb": []magicAttribute{magicAttribute{ID: 9, Name: "+{0} to Mana", Values: []int64{10}}},
	"gfb": []magicAttribute{magicAttribute{ID: 9, Name: "+{0} to Mana", Values: []int64{17}}},
	"gsb": []magicAttribute{magicAttribute{ID: 9, Name: "+{0} to Mana", Values: []int64{24}}},
	"glb": []magicAttribute{magicAttribute{ID: 9, Name: "+{0} to Mana", Values: []int64{31}}},
	"gpb": []magicAttribute{magicAttribute{ID: 9, Name: "+{0} to Mana", Values: []int64{38}}},

	// Skull
	"skc": []magicAttribute{
		magicAttribute{ID: 74, Name: "Replenish Life +{0}", Values: []int64{2}},
		magicAttribute{ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{8}},
	},

	"skf": []magicAttribute{
		magicAttribute{ID: 74, Name: "Replenish Life +{0}", Values: []int64{3}},
		magicAttribute{ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{8}},
	},

	"sku": []magicAttribute{
		magicAttribute{ID: 74, Name: "Replenish Life +{0}", Values: []int64{3}},
		magicAttribute{ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{12}},
	},

	"skl": []magicAttribute{
		magicAttribute{ID: 74, Name: "Replenish Life +{0}", Values: []int64{4}},
		magicAttribute{ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{12}},
	},

	"skz": []magicAttribute{
		magicAttribute{ID: 74, Name: "Replenish Life +{0}", Values: []int64{5}},
		magicAttribute{ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{19}},
	},

	// Topaz
	"gcy": []magicAttribute{magicAttribute{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{9}}},
	"gfy": []magicAttribute{magicAttribute{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{13}}},
	"gsy": []magicAttribute{magicAttribute{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{16}}},
	"gly": []magicAttribute{magicAttribute{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{20}}},
	"gpy": []magicAttribute{magicAttribute{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{24}}},
}

var socketablesShields = map[string][]magicAttribute{
	"r01": []magicAttribute{
		magicAttribute{ID: 31, Name: "+{0} Defense", Values: []int64{15}},
		magicAttribute{ID: 89, Name: "+{0} to Light Radius", Values: []int64{1}},
	},
	"r02": []magicAttribute{magicAttribute{ID: 20, Name: "+{0}% Increased chance of blocking", Values: []int64{7}}},
	"r03": []magicAttribute{magicAttribute{ID: 138, Name: "+{0} to Mana After Each Kill", Values: []int64{2}}},
	"r04": []magicAttribute{magicAttribute{ID: 32, Name: "+{0} vs. Missile", Values: []int64{30}}},
	"r05": []magicAttribute{magicAttribute{ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{15}}},
	"r06": []magicAttribute{magicAttribute{ID: 114, Name: "{0}% Damage Taken Goes to Mana", Values: []int64{15}}},
	"r07": []magicAttribute{magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{35}}},
	"r08": []magicAttribute{magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{35}}},
	"r09": []magicAttribute{magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{35}}},
	"r10": []magicAttribute{magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{35}}},
	"r11": []magicAttribute{magicAttribute{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{14}}},
	"r12": []magicAttribute{magicAttribute{ID: 34, Name: "Damage Reduced by {0}", Values: []int64{7}}},
	"r13": []magicAttribute{magicAttribute{ID: 102, Name: "{0}% Faster Block Rate", Values: []int64{20}}},
	"r14": []magicAttribute{magicAttribute{ID: 74, Name: "Replenish Life +{0}", Values: []int64{7}}},
	"r15": []magicAttribute{magicAttribute{ID: 91, Name: "Requirements {0}%", Values: []int64{-15}}},
	"r16": []magicAttribute{magicAttribute{ID: 3, Name: "+{0} to Vitality", Values: []int64{10}}},
	"r17": []magicAttribute{magicAttribute{ID: 1, Name: "+{0} to Energy", Values: []int64{10}}},
	"r18": []magicAttribute{magicAttribute{ID: 2, Name: "+{0} to Dexterity", Values: []int64{10}}},
	"r19": []magicAttribute{magicAttribute{ID: 0, Name: "+{0} to Strength", Values: []int64{10}}},
	"r20": []magicAttribute{magicAttribute{ID: 79, Name: "{0}% Extra Gold from Monsters", Values: []int64{50}}},
	"r21": []magicAttribute{magicAttribute{ID: 16, Name: "+{0}% Enhanced Defense", Values: []int64{30}}},
	"r22": []magicAttribute{
		magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{22}},
		magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{22}},
		magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{22}},
		magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{22}},
	},
	"r23": []magicAttribute{magicAttribute{ID: 35, Name: "Magic Damage Reduced by {0}", Values: []int64{7}}},
	"r24": []magicAttribute{magicAttribute{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{25}}},
	"r25": []magicAttribute{magicAttribute{ID: 46, Name: "+{0}% to Maximum Poison Resist", Values: []int64{5}}},
	"r26": []magicAttribute{magicAttribute{ID: 40, Name: "+{0}% to Maximum Fire Resist", Values: []int64{5}}},
	"r27": []magicAttribute{magicAttribute{ID: 44, Name: "+{0}% to Maximum Cold Resist", Values: []int64{5}}},
	"r28": []magicAttribute{magicAttribute{ID: 42, Name: "+{0}% to Maximum Lightning Resist", Values: []int64{5}}},
	"r29": []magicAttribute{magicAttribute{ID: 9, Name: "+{0} to Mana", Values: []int64{50}}},
	"r30": []magicAttribute{magicAttribute{ID: 36, Name: "Damage Reduced by {0}%", Values: []int64{8}}},
	"r31": []magicAttribute{magicAttribute{ID: 7, Name: "+{0} to Life", Values: []int64{50}}},
	"r32": []magicAttribute{magicAttribute{ID: 153, Name: "Cannot Be Frozen", Values: []int64{1}}},
	"r33": []magicAttribute{magicAttribute{ID: 152, Name: "Indestructible", Values: []int64{1}}},

	// Amethyst
	"gcv": []magicAttribute{magicAttribute{ID: 31, Name: "+{0} Defense", Values: []int64{8}}},
	"gfv": []magicAttribute{magicAttribute{ID: 31, Name: "+{0} Defense", Values: []int64{12}}},
	"gsv": []magicAttribute{magicAttribute{ID: 31, Name: "+{0} Defense", Values: []int64{18}}},
	"gzv": []magicAttribute{magicAttribute{ID: 31, Name: "+{0} Defense", Values: []int64{24}}},
	"gpv": []magicAttribute{magicAttribute{ID: 31, Name: "+{0} Defense", Values: []int64{30}}},

	// Diamond

	"gcw": []magicAttribute{
		magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{6}},
		magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{6}},
		magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{6}},
		magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{6}},
	},

	"gfw": []magicAttribute{
		magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{8}},
		magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{8}},
		magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{8}},
		magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{8}},
	},

	"gsw": []magicAttribute{
		magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{11}},
		magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{11}},
		magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{11}},
		magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{11}},
	},

	"glw": []magicAttribute{
		magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{14}},
		magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{14}},
		magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{14}},
		magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{14}},
	},

	"gpw": []magicAttribute{
		magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{19}},
		magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{19}},
		magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{19}},
		magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{19}},
	},

	// Emerald
	"gcg": []magicAttribute{magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{12}}},
	"gfg": []magicAttribute{magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{16}}},
	"gsg": []magicAttribute{magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{22}}},
	"glg": []magicAttribute{magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{28}}},
	"gpg": []magicAttribute{magicAttribute{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{40}}},

	// Ruby
	"gcr": []magicAttribute{magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{12}}},
	"gfr": []magicAttribute{magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{16}}},
	"gsr": []magicAttribute{magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{22}}},
	"glr": []magicAttribute{magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{28}}},
	"gpr": []magicAttribute{magicAttribute{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{40}}},

	// Sapphire
	"gcb": []magicAttribute{magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{12}}},
	"gfb": []magicAttribute{magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{16}}},
	"gsb": []magicAttribute{magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{22}}},
	"glb": []magicAttribute{magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{28}}},
	"gpb": []magicAttribute{magicAttribute{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{40}}},

	// Skull
	"skc": []magicAttribute{magicAttribute{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{4}}},
	"skf": []magicAttribute{magicAttribute{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{8}}},
	"sku": []magicAttribute{magicAttribute{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{12}}},
	"skl": []magicAttribute{magicAttribute{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{16}}},
	"skz": []magicAttribute{magicAttribute{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{20}}},

	// Topaz
	"gcy": []magicAttribute{magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{12}}},
	"gfy": []magicAttribute{magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{16}}},
	"gsy": []magicAttribute{magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{22}}},
	"gly": []magicAttribute{magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{28}}},
	"gpy": []magicAttribute{magicAttribute{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{40}}},
}
