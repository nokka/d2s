package d2s

var socketablesWeapons = map[string][]magicAttribute{
	// Runes
	"r01": {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{50}}, {ID: 89, Name: "+{0} to Light Radius", Values: []int64{1}}},
	"r02": {{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{75}}, {ID: 124, Name: "+{0} to Attack Rating against Undead", Values: []int64{50}}},
	"r03": {{ID: 138, Name: "+{0} to Mana After Each Kill", Values: []int64{2}}},
	"r04": {{ID: 81, Name: "Knockback", Values: []int64{1}}},
	"r05": {{ID: 116, Name: "{0}% Target Defense", Values: []int64{-25}}},
	"r06": {{ID: 22, Name: "+{0} to Maximum 1-handed damage", Values: []int64{9}}, {ID: 24, Name: "+{0} to Maximum 2-handed damage", Values: []int64{9}}},
	"r07": {{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{154, 154, 125}}},
	"r08": {{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{5, 30}}},
	"r09": {{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 50}}},
	"r10": {{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{3, 14, 75}}},
	"r11": {{ID: 60, Name: "{0}% Life stolen per hit", Values: []int64{7}}},
	"r12": {{ID: 21, Name: "+{0} to Minimum 1-handed damage", Values: []int64{9}}, {ID: 23, Name: "+{0} to Minimum 2-handed damage", Values: []int64{9}}},
	"r13": {{ID: 93, Name: "{0}% Increased Attack Speed", Values: []int64{20}}},
	"r14": {{ID: 112, Name: "Hit Causes Monsters to Flee {0}%", Values: []int64{25}}},
	"r15": {{ID: 91, Name: "Requirements {0}%", Values: []int64{-20}}},
	"r16": {{ID: 3, Name: "+{0} to Vitality", Values: []int64{10}}},
	"r17": {{ID: 1, Name: "+{0} to Energy", Values: []int64{10}}},
	"r18": {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{10}}},
	"r19": {{ID: 0, Name: "+{0} to Strength", Values: []int64{10}}},
	"r20": {{ID: 79, Name: "{0}% Extra Gold from Monsters", Values: []int64{75}}},
	"r21": {{ID: 121, Name: "+{0}% Damage to Demons", Values: []int64{75}}, {ID: 123, Name: "+{0} to Attack Rating against Demons", Values: []int64{100}}},
	"r22": {{ID: 135, Name: "{0}% Chance of Open Wounds", Values: []int64{25}}},
	"r23": {{ID: 117, Name: "Prevent Monster Heal", Values: []int64{1}}},
	"r24": {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{30}}},
	"r25": {{ID: 119, Name: "{0}% Bonus to Attack Rating", Values: []int64{20}}},
	"r26": {{ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{7}}},
	"r27": {{ID: 17, Name: "+{0}% Enhanced Damage", Values: []int64{50, 50}}},
	"r28": {{ID: 141, Name: "{0}% Deadly Strike", Values: []int64{20}}},
	"r29": {{ID: 113, Name: "Hit Blinds Target +{0}", Values: []int64{20}}},
	"r30": {{ID: 136, Name: "{0}% Chance of Crushing Blow", Values: []int64{20}}},
	"r31": {{ID: 115, Name: "Ignore Target Defense", Values: []int64{1}}},
	"r32": {{ID: 134, Name: "Freezes Target +{0}", Values: []int64{3}}},
	"r33": {{ID: 152, Name: "Indestructible", Values: []int64{1}}},

	// Amethyst
	"gcv": {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{40}}},
	"gfv": {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{60}}},
	"gsv": {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{80}}},
	"gzv": {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{100}}},
	"gpv": {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{150}}},

	// Diamond
	"gcw": {{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{28}}},
	"gfw": {{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{34}}},
	"gsw": {{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{44}}},
	"glw": {{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{54}}},
	"gpw": {{ID: 122, Name: "+{0}% Damage to Undead", Values: []int64{68}}},

	// Emerald
	"gcg": {{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{35, 35, 75}}},
	"gfg": {{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{52, 52, 100}}},
	"gsg": {{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{82, 82, 125}}},
	"glg": {{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{103, 103, 150}}},
	"gpg": {{ID: 57, Name: "Adds {0}-{1} Poison Damage over {2} Seconds", Values: []int64{147, 147, 175}}},

	// Ruby
	"gcr": {{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{3, 4}}},
	"gfr": {{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{5, 8}}},
	"gsr": {{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{8, 12}}},
	"glr": {{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{10, 16}}},
	"gpr": {{ID: 48, Name: "Adds {0}-{1} Fire Damage", Values: []int64{15, 20}}},

	// Sapphire
	"gcb": {{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{1, 3}}},
	"gfb": {{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{3, 5}}},
	"gsb": {{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{4, 7}}},
	"glb": {{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{6, 10}}},
	"gpb": {{ID: 54, Name: "Adds {0}-{1} Cold Damage", Values: []int64{10, 14}}},

	// Skull
	"skc": {{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{2}}, {ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{1}}},
	"skf": {{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{2}}, {ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{2}}},
	"sku": {{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{3}}, {ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{2}}},
	"skl": {{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{3}}, {ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{3}}},
	"skz": {{ID: 60, Name: "{0}% Life Stolen Per Hit", Values: []int64{4}}, {ID: 62, Name: "{0}% Mana Stolen Per Hit", Values: []int64{3}}},

	// Topaz
	"gcy": {{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 8}}},
	"gfy": {{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 14}}},
	"gsy": {{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 22}}},
	"gly": {{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 30}}},
	"gpy": {{ID: 50, Name: "Adds {0}-{1} Lightning Damage", Values: []int64{1, 40}}},
}

var socketablesArmor = map[string][]magicAttribute{
	"r01": {{ID: 31, Name: "+{0} Defense", Values: []int64{15}}, {ID: 89, Name: "+{0} to Light Radius", Values: []int64{1}}},
	"r02": {{ID: 154, Name: "{0}% Slower Stamina Drain", Values: []int64{15}}},
	"r03": {{ID: 138, Name: "+{0} to Mana After Each Kill", Values: []int64{2}}},
	"r04": {{ID: 32, Name: "+{0} vs. Missile", Values: []int64{30}}},
	"r05": {{ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{15}}},
	"r06": {{ID: 114, Name: "{0}% Damage Taken Goes to Mana", Values: []int64{15}}},
	"r07": {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{35}}},
	"r08": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{35}}},
	"r09": {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{35}}},
	"r10": {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{35}}},
	"r11": {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{14}}},
	"r12": {{ID: 34, Name: "Damage Reduced by {0}", Values: []int64{7}}},
	"r13": {{ID: 99, Name: "{0}% Faster Hit Recovery", Values: []int64{20}}},
	"r14": {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{7}}},
	"r15": {{ID: 91, Name: "Requirements {0}%", Values: []int64{-15}}},
	"r16": {{ID: 3, Name: "+{0} to Vitality", Values: []int64{10}}},
	"r17": {{ID: 1, Name: "+{0} to Energy", Values: []int64{10}}},
	"r18": {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{10}}},
	"r19": {{ID: 0, Name: "+{0} to Strength", Values: []int64{10}}},
	"r20": {{ID: 79, Name: "{0}% Extra Gold from Monsters", Values: []int64{50}}},
	"r21": {{ID: 16, Name: "+{0}% Enhanced Defense", Values: []int64{30}}},
	"r22": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{15}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{15}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{15}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{15}}},
	"r23": {{ID: 35, Name: "Magic Damage Reduced by {0}", Values: []int64{7}}},
	"r24": {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{25}}},
	"r25": {{ID: 46, Name: "+{0}% to Maximum Poison Resist", Values: []int64{5}}},
	"r26": {{ID: 40, Name: "+{0}% to Maximum Fire Resist", Values: []int64{5}}},
	"r27": {{ID: 44, Name: "+{0}% to Maximum Cold Resist", Values: []int64{5}}},
	"r28": {{ID: 42, Name: "+{0}% to Maximum Lightning Resist", Values: []int64{5}}},
	"r29": {{ID: 77, Name: "Increase Maximum Mana {0}%", Values: []int64{5}}},
	"r30": {{ID: 36, Name: "Damage Reduced by {0}%", Values: []int64{8}}},
	"r31": {{ID: 76, Name: "Increase Maximum Life {0}%", Values: []int64{5}}},
	"r32": {{ID: 153, Name: "Cannot Be Frozen", Values: []int64{1}}},
	"r33": {{ID: 152, Name: "Indestructible", Values: []int64{1}}},

	// Amethyst
	"gcv": {{ID: 0, Name: "+{0} to Strength", Values: []int64{3}}},
	"gfv": {{ID: 0, Name: "+{0} to Strength", Values: []int64{4}}},
	"gsv": {{ID: 0, Name: "+{0} to Strength", Values: []int64{6}}},
	"gzv": {{ID: 0, Name: "+{0} to Strength", Values: []int64{8}}},
	"gpv": {{ID: 0, Name: "+{0} to Strength", Values: []int64{10}}},

	// Diamond
	"gcw": {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{20}}},
	"gfw": {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{40}}},
	"gsw": {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{60}}},
	"glw": {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{80}}},
	"gpw": {{ID: 19, Name: "+{0} to Attack rating", Values: []int64{100}}},

	// Emerald
	"gcg": {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{3}}},
	"gfg": {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{4}}},
	"gsg": {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{6}}},
	"glg": {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{8}}},
	"gpg": {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{10}}},

	// Ruby
	"gcr": {{ID: 7, Name: "+{0} to Life", Values: []int64{10}}},
	"gfr": {{ID: 7, Name: "+{0} to Life", Values: []int64{17}}},
	"gsr": {{ID: 7, Name: "+{0} to Life", Values: []int64{24}}},
	"glr": {{ID: 7, Name: "+{0} to Life", Values: []int64{31}}},
	"gpr": {{ID: 7, Name: "+{0} to Life", Values: []int64{38}}},

	// Sapphire
	"gcb": {{ID: 9, Name: "+{0} to Mana", Values: []int64{10}}},
	"gfb": {{ID: 9, Name: "+{0} to Mana", Values: []int64{17}}},
	"gsb": {{ID: 9, Name: "+{0} to Mana", Values: []int64{24}}},
	"glb": {{ID: 9, Name: "+{0} to Mana", Values: []int64{31}}},
	"gpb": {{ID: 9, Name: "+{0} to Mana", Values: []int64{38}}},

	// Skull
	"skc": {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{2}}, {ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{8}}},
	"skf": {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{3}}, {ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{8}}},
	"sku": {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{3}}, {ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{12}}},
	"skl": {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{4}}, {ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{12}}},
	"skz": {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{5}}, {ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{19}}},

	// Topaz
	"gcy": {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{9}}},
	"gfy": {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{13}}},
	"gsy": {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{16}}},
	"gly": {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{20}}},
	"gpy": {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{24}}},
}

var socketablesShields = map[string][]magicAttribute{
	"r01": {{ID: 31, Name: "+{0} Defense", Values: []int64{15}}, {ID: 89, Name: "+{0} to Light Radius", Values: []int64{1}}},
	"r02": {{ID: 20, Name: "+{0}% Increased chance of blocking", Values: []int64{7}}},
	"r03": {{ID: 138, Name: "+{0} to Mana After Each Kill", Values: []int64{2}}},
	"r04": {{ID: 32, Name: "+{0} vs. Missile", Values: []int64{30}}},
	"r05": {{ID: 27, Name: "Regenerate Mana {0}%", Values: []int64{15}}},
	"r06": {{ID: 114, Name: "{0}% Damage Taken Goes to Mana", Values: []int64{15}}},
	"r07": {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{35}}},
	"r08": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{35}}},
	"r09": {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{35}}},
	"r10": {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{35}}},
	"r11": {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{14}}},
	"r12": {{ID: 34, Name: "Damage Reduced by {0}", Values: []int64{7}}},
	"r13": {{ID: 102, Name: "{0}% Faster Block Rate", Values: []int64{20}}},
	"r14": {{ID: 74, Name: "Replenish Life +{0}", Values: []int64{7}}},
	"r15": {{ID: 91, Name: "Requirements {0}%", Values: []int64{-15}}},
	"r16": {{ID: 3, Name: "+{0} to Vitality", Values: []int64{10}}},
	"r17": {{ID: 1, Name: "+{0} to Energy", Values: []int64{10}}},
	"r18": {{ID: 2, Name: "+{0} to Dexterity", Values: []int64{10}}},
	"r19": {{ID: 0, Name: "+{0} to Strength", Values: []int64{10}}},
	"r20": {{ID: 79, Name: "{0}% Extra Gold from Monsters", Values: []int64{50}}},
	"r21": {{ID: 16, Name: "+{0}% Enhanced Defense", Values: []int64{30}}},
	"r22": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{22}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{22}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{22}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{22}}},
	"r23": {{ID: 35, Name: "Magic Damage Reduced by {0}", Values: []int64{7}}},
	"r24": {{ID: 80, Name: "{0}% Better Chance of Getting Magic Items", Values: []int64{25}}},
	"r25": {{ID: 46, Name: "+{0}% to Maximum Poison Resist", Values: []int64{5}}},
	"r26": {{ID: 40, Name: "+{0}% to Maximum Fire Resist", Values: []int64{5}}},
	"r27": {{ID: 44, Name: "+{0}% to Maximum Cold Resist", Values: []int64{5}}},
	"r28": {{ID: 42, Name: "+{0}% to Maximum Lightning Resist", Values: []int64{5}}},
	"r29": {{ID: 9, Name: "+{0} to Mana", Values: []int64{50}}},
	"r30": {{ID: 36, Name: "Damage Reduced by {0}%", Values: []int64{8}}},
	"r31": {{ID: 7, Name: "+{0} to Life", Values: []int64{50}}},
	"r32": {{ID: 153, Name: "Cannot Be Frozen", Values: []int64{1}}},
	"r33": {{ID: 152, Name: "Indestructible", Values: []int64{1}}},

	// Amethyst
	"gcv": {{ID: 31, Name: "+{0} Defense", Values: []int64{8}}},
	"gfv": {{ID: 31, Name: "+{0} Defense", Values: []int64{12}}},
	"gsv": {{ID: 31, Name: "+{0} Defense", Values: []int64{18}}},
	"gzv": {{ID: 31, Name: "+{0} Defense", Values: []int64{24}}},
	"gpv": {{ID: 31, Name: "+{0} Defense", Values: []int64{30}}},

	// Diamond

	"gcw": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{6}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{6}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{6}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{6}}},
	"gfw": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{8}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{8}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{8}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{8}}},
	"gsw": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{11}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{11}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{11}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{11}}},
	"glw": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{14}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{14}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{14}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{14}}},
	"gpw": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{19}}, {ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{19}}, {ID: 43, Name: "Cold Resist +{0}%", Values: []int64{19}}, {ID: 45, Name: "Poison Resist +{0}%", Values: []int64{19}}},

	// Emerald
	"gcg": {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{12}}},
	"gfg": {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{16}}},
	"gsg": {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{22}}},
	"glg": {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{28}}},
	"gpg": {{ID: 45, Name: "Poison Resist +{0}%", Values: []int64{40}}},

	// Ruby
	"gcr": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{12}}},
	"gfr": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{16}}},
	"gsr": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{22}}},
	"glr": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{28}}},
	"gpr": {{ID: 39, Name: "Fire Resist +{0}%", Values: []int64{40}}},

	// Sapphire
	"gcb": {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{12}}},
	"gfb": {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{16}}},
	"gsb": {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{22}}},
	"glb": {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{28}}},
	"gpb": {{ID: 43, Name: "Cold Resist +{0}%", Values: []int64{40}}},

	// Skull
	"skc": {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{4}}},
	"skf": {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{8}}},
	"sku": {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{12}}},
	"skl": {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{16}}},
	"skz": {{ID: 78, Name: "Attacker Takes Damage of {0}", Values: []int64{20}}},

	// Topaz
	"gcy": {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{12}}},
	"gfy": {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{16}}},
	"gsy": {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{22}}},
	"gly": {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{28}}},
	"gpy": {{ID: 41, Name: "Lightning Resist +{0}%", Values: []int64{40}}},
}
