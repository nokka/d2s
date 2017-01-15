package d2s

// Item represents a base 111 bit item.
// Item represents an actual item
type Item struct {
	Identified         uint64
	Socketed           uint64
	New                uint64
	IsEar              uint64
	StarterItem        uint64
	SimpleItem         uint64
	Ethereal           uint64
	Personalized       uint64
	GivenRuneword      uint64
	LocationID         uint64
	EquippedID         uint64
	PositionY          uint64
	PositionX          uint64
	AltPositionID      uint64
	Type               string
	TypeID             uint64
	NrOfItemsInSockets uint64
	ID                 uint64
	Level              uint64
	Quality            uint64
	MultiplePictures   uint64
	PictureID          uint64
	ClassSpecific      uint64
	LowQualityID       uint64
	Timestamp          uint64
	DefenseRating      uint64
	MaxDurability      uint64
	CurrentDurability  uint64
	TotalNrOfSockets   uint64
	Quantity           uint64

	// Magical Item properties
	MagicPrefix uint64
	MagicSuffix uint64

	// Set item properties
	SetID        uint64
	SetItemLists uint64

	// All item types >= magicallyEnhanced
	MagicAttributes []magicAttribute
}

func (i Item) getTypeID() uint64 {

	if _, ok := armorCodes[i.Type]; ok {
		return armor
	}

	if _, ok := weaponCodes[i.Type]; ok {
		return weapon
	}

	return other
}

// Item types.
const (
	armor  = 0x01
	weapon = 0x02
	other  = 0x03
)

// Item location IDs.
const (
	stored   = 0x00
	equipped = 0x01
	belt     = 0x02
	cursor   = 0x03
	socketed = 0x04
)

// Rarity IDs.
const (
	lowQuality        = 0x01
	normal            = 0x02
	highQuality       = 0x03
	magicallyEnhanced = 0x04
	partOfSet         = 0x05
	rare              = 0x06
	unique            = 0x07
	crafted           = 0x08
)

// Jewelry and misc codes.
const (
	Ring = "rin"
)

// Armor codes.
var armorCodes = map[string]string{
	"sml": "Small Shield",
	"qui": "Quilted Armor",
	"skp": "Skull Cap",
}

// Weapon Codes.
var weaponCodes = map[string]string{
	"cbw": "Composite Bow",
	"dgr": "Dagger",
}

// Weapon codes.
const (
	AncientAxe          = "9gi"
	AncientSword        = "9wd"
	Arbalest            = "8lx"
	ArchonStaff         = "6ws"
	AshwoodBow          = "am6"
	Ataghan             = "7sm"
	Axe                 = "axe"
	BalancedAxe         = "bal"
	BalancedKnife       = "bkf"
	Balista             = "8hx"
	BalrogBlade         = "7gs"
	BalrogSpear         = "7s7"
	BarbedClub          = "9sp"
	Bardiche            = "bar"
	BastardSword        = "bsw"
	BattleAxe           = "btx"
	BattleCestus        = "7cs"
	BattleDart          = "9tk"
	BattleHammer        = "9wh"
	BattleScythe        = "9s8"
	BattleStaff         = "bst"
	BattleSword         = "9bs"
	BeardedAxe          = "9ba"
	BecDeCorbin         = "9h9"
	BerserkerAxe        = "7wa"
	Bill                = "9vo"
	Blade               = "bld"
	BladeBow            = "6hb"
	BladeTalons         = "btl"
	BoneKnife           = "7dg"
	BoneWand            = "bwn"
	Brandistock         = "brn"
	BroadAxe            = "bax"
	BroadSword          = "bsd"
	BurntWand           = "9wn"
	Caduceus            = "7ws"
	CedarBow            = "8lb"
	CedarStaff          = "8cs"
	CeremonialBow       = "am7"
	CeremonialJavelin   = "ama"
	CeremonialPike      = "am9"
	CeremonialSpear     = "am8"
	Cestus              = "ces"
	ChampionAxe         = "7ga"
	ChampionSword       = "7b7"
	ChokingGasPotion    = "gpm"
	ChuKoNu             = "8rx"
	Cinquedeas          = "9kr"
	ClaspedOrb          = "ob4"
	Claws               = "clw"
	Claymore            = "clm"
	Cleaver             = "9ax"
	CloudySphere        = "ob8"
	Club                = "clb"
	ColossalSword       = "7fb"
	ColossalBlade       = "7gd"
	ColossusCrossbow    = "6hx"
	ColossusVoulge      = "7vo"
	CompositeBow        = "cbw"
	ConquestSword       = "7bs"
	Crossbow            = "mxb"
	Crowbill            = "9mp"
	CrusaderBow         = "6l7"
	CrypticAxe          = "7pa"
	CrypticSword        = "7ls"
	CrystalSword        = "crs"
	CrystallineGlobe    = "ob7"
	Cudgel              = "9cl"
	Cutlass             = "9sm"
	DacianFalx          = "9cm"
	Dagger              = "dgr"
	Decapitator         = "7bt"
	DecoyDagger         = "d33"
	DemonCrossbow       = "6rx"
	DemonHeart          = "obd"
	DevilStar           = "7mt"
	DiamondBow          = "6s7"
	DimensionalBlade    = "9cr"
	DimensionalShard    = "obf"
	Dirk                = "dir"
	DivineScepter       = "9ws"
	DoubleAxe           = "2ax"
	DoubleBow           = "8cb"
	DragonStone         = "ob5"
	EableOrb            = "ob1"
	EdgeBow             = "8sb"
	ElderStaff          = "6cs"
	EldritchOrb         = "obc"
	ElegantBlade        = "7sb"
	Espadon             = "92h"
	EttinAxe            = "72a"
	ExecutionerSword    = "9gd"
	ExplodingPotion     = "opm"
	Falcata             = "7ss"
	Falchion            = "flc"
	FangedKnife         = "7kr"
	Fascia              = "9xf"
	FeralAxe            = "7la"
	FeralClaws          = "7lw"
	Flail               = "fla"
	Flamberge           = "flb"
	FlangedMace         = "9ma"
	FlyingAxe           = "7ta"
	FlyingKnife         = "7tk"
	Francisca           = "9ta"
	FulmatingPotion     = "opl"
	Fuscina             = "9tr"
	GhostGlaive         = "7gl"
	GhostSpear          = "7st"
	GhostWand           = "7yw"
	GiantAxe            = "gix"
	GiantSword          = "gis"
	GiantThresher       = "7wc"
	Gidbinn             = "g33"
	Gladius             = "9ss"
	Glaive              = "glv"
	GloriousAxe         = "7gi"
	GlowingOrb          = "ob6"
	GnarledStaff        = "cst"
	GorgonCrossbow      = "6mx"
	GothicAxe           = "9ga"
	GothicBow           = "8lw"
	GothicStaff         = "8bs"
	GothicSword         = "9b9"
	GrandMatronBow      = "amc"
	GrandScepter        = "gsc"
	GraveWand           = "9gw"
	GreatAxe            = "gax"
	GreatBow            = "6cb"
	GreatMaul           = "gma"
	GreatPilum          = "9pi"
	GreatPoleaxe        = "7h7"
	GreatSword          = "gsd"
	GreaterClaws        = "9lw"
	GreaterTalons       = "9tw"
	GrimScythe          = "9wc"
	GrimWand            = "gwn"
	Halberd             = "hal"
	HandAxe             = "hax"
	HandScythe          = "9cs"
	Harpoon             = "9ts"
	Hatchet             = "9ha"
	HatchetHands        = "axf"
	HeavenlyStone       = "obb"
	HeavyCrossbow       = "hxb"
	HellforgeHammer     = "hfh"
	HighlandBlade       = "7cm"
	HolyWaterSprinkler  = "9qs"
	HoradricMalus       = "hdm"
	HoradricStaff       = "hst"
	HuntersBow          = "hbw"
	Hurlbat             = "9b8"
	HydraBow            = "6lw"
	HydraEdge           = "7fc"
	HyperionJavelin     = "7ja"
	HyperionSpear       = "7sr"
	JaggedStar          = "9mt"
	Javelin             = "jav"
	JoStaff             = "8ss"
	Katar               = "ktr"
	KhalimFlail         = "qf1"
	Knout               = "9fl"
	Kriss               = "kri"
	Lance               = "9p9"
	LargeAxe            = "lax"
	LegendSpike         = "7bl"
	LegendSword         = "72h"
	LegendaryMallet     = "7wh"
	LichWand            = "7bw"
	LightCrossbow       = "lxb"
	LochaberAxe         = "9b7"
	LongBattleBow       = "lbb"
	LongBow             = "lbw"
	LongSiegeBow        = "8l8"
	LongStaff           = "lst"
	LongSword           = "lsd"
	LongWarBow          = "lwb"
	Mace                = "mac"
	MaidenJavelin       = "am5"
	MaidenPike          = "am4"
	MaidenSpear         = "am3"
	Mancatcher          = "7br"
	MartelDeFar         = "9gm"
	MatriarchalBow      = "amb"
	MatriarchalPike     = "ame"
	MatriarchalSpear    = "amd"
	MatriarchalJavelin  = "amf"
	Maul                = "mau"
	MightyScepter       = "7sc"
	MilitaryAxe         = "9la"
	MilitaryPick        = "mpi"
	MithralPoint        = "7di"
	MorningStar         = "mst"
	MythicalSword       = "7wd"
	Naga                = "9wa"
	OgreAxe             = "7o7"
	OgreMaul            = "7m7"
	OilPotion           = "ops"
	Partizan            = "9pa"
	PelletBow           = "6lx"
	PetrifiedWand       = "9yw"
	PhaseBlade          = "7cr"
	Pike                = "pik"
	Pilum               = "pil"
	Poignard            = "9dg"
	Poleaxe             = "pax"
	PolishedWand        = "7wn"
	QuarterStaff        = "8ls"
	Quhab               = "9ar"
	RancidGasPotion     = "gps"
	RazorBow            = "8hb"
	ReflexBow           = "am2"
	ReinforcedMace      = "7ma"
	RepeatingCrossbow   = "rxb"
	Rondel              = "9di"
	RuneBow             = "8sw"
	RuneScepter         = "9sc"
	RuneStaff           = "8ws"
	RuneSword           = "9ls"
	RunicTalons         = "7tw"
	Sabre               = "sbr"
	SacredGlobe         = "ob2"
	Scepter             = "scp"
	Scimitar            = "scm"
	ScissorsKatar       = "skr"
	ScissorsQuhab       = "9qr"
	ScissorsSuwayyah    = "7qr"
	Scourge             = "7fl"
	Scythe              = "scy"
	SeraphRod           = "7qs"
	ShadowBow           = "6lb"
	Shamshir            = "9sb"
	Shillelah           = "6bs"
	ShortBattleBow      = "sbb"
	ShortBow            = "sbw"
	ShortSiegeBow       = "8s8"
	ShortSpear          = "ssp"
	ShortStaff          = "sst"
	ShortSword          = "ssd"
	ShortWarBow         = "swb"
	SiegeCrossbow       = "8mx"
	SilverEdgedAxe      = "7ba"
	Simbilan            = "9s9"
	SmallCrescent       = "7ax"
	SmokedSphere        = "ob3"
	SparklingBall       = "ob9"
	Spear               = "spr"
	Spetum              = "spt"
	Spiculum            = "9gl"
	SpiderBow           = "6sb"
	SpikedClub          = "spc"
	StaffOfTheKings     = "msf"
	StagBow             = "am1"
	Stalagmite          = "6ls"
	Stilleto            = "9bl"
	StranglingGasPotion = "gpl"
	StygianPike         = "7tr"
	StygianPilum        = "7pi"
	SuperKhalimFlail    = "qf2"
	Suwayyah            = "7ar"
	SwirlingCrystal     = "oba"
	Tabar               = "9bt"
	Thresher            = "7s8"
	ThrowingAxe         = "tax"
	ThrowingKnife       = "tkf"
	ThrowingSpear       = "tsp"
	ThunderMaul         = "7gm"
	Tomahawk            = "7ha"
	TombWand            = "9bw"
	Trident             = "tri"
	Truncheon           = "7cl"
	Tulwar              = "9fc"
	TuskSword           = "9gs"
	TwinAxe             = "92a"
	TwoHandedAxe        = "2hs"
	TyrantClub          = "7sp"
	UnearthedWand       = "7gw"
	VortexOrb           = "obe"
	Voulge              = "vou"
	WalkingStick        = "6ss"
	Wand                = "wnd"
	WarAxe              = "wax"
	WarClub             = "9m9"
	WarDart             = "9bk"
	WarFist             = "7xf"
	WarFork             = "9br"
	WarHammer           = "whm"
	WarJavelin          = "9ja"
	WarPike             = "7p7"
	WarScepter          = "wsp"
	WarScythe           = "wsc"
	WarSpear            = "9sr"
	WarSpike            = "7mp"
	WarStaff            = "wst"
	WarSword            = "wsd"
	WardBow             = "6sw"
	WingedAxe           = "7b8"
	WingedHarpoon       = "7ts"
	WingedKnife         = "7bk"
	WirtsLeg            = "leg"
	WristBlade          = "wrb"
	WristSpike          = "9wb"
	WristSword          = "7wb"
	Yari                = "9st"
	YewWand             = "ywn"
	Zweihander          = "9fb"
)

/*
// Weapon codes.
const (
	AncientAxe          = "9gi"
	AncientSword        = "9wd"
	Arbalest            = "8lx"
	ArchonStaff         = "6ws"
	AshwoodBow          = "am6"
	Ataghan             = "7sm"
	Axe                 = "axe"
	BalancedAxe         = "bal"
	BalancedKnife       = "bkf"
	Balista             = "8hx"
	BalrogBlade         = "7gs"
	BalrogSpear         = "7s7"
	BarbedClub          = "9sp"
	Bardiche            = "bar"
	BastardSword        = "bsw"
	BattleAxe           = "btx"
	BattleCestus        = "7cs"
	BattleDart          = "9tk"
	BattleHammer        = "9wh"
	BattleScythe        = "9s8"
	BattleStaff         = "bst"
	BattleSword         = "9bs"
	BeardedAxe          = "9ba"
	BecDeCorbin         = "9h9"
	BerserkerAxe        = "7wa"
	Bill                = "9vo"
	Blade               = "bld"
	BladeBow            = "6hb"
	BladeTalons         = "btl"
	BoneKnife           = "7dg"
	BoneWand            = "bwn"
	Brandistock         = "brn"
	BroadAxe            = "bax"
	BroadSword          = "bsd"
	BurntWand           = "9wn"
	Caduceus            = "7ws"
	CedarBow            = "8lb"
	CedarStaff          = "8cs"
	CeremonialBow       = "am7"
	CeremonialJavelin   = "ama"
	CeremonialPike      = "am9"
	CeremonialSpear     = "am8"
	Cestus              = "ces"
	ChampionAxe         = "7ga"
	ChampionSword       = "7b7"
	ChokingGasPotion    = "gpm"
	ChuKoNu             = "8rx"
	Cinquedeas          = "9kr"
	ClaspedOrb          = "ob4"
	Claws               = "clw"
	Claymore            = "clm"
	Cleaver             = "9ax"
	CloudySphere        = "ob8"
	Club                = "clb"
	ColossalSword       = "7fb"
	ColossalBlade       = "7gd"
	ColossusCrossbow    = "6hx"
	ColossusVoulge      = "7vo"
	CompositeBow        = "cbw"
	ConquestSword       = "7bs"
	Crossbow            = "mxb"
	Crowbill            = "9mp"
	CrusaderBow         = "6l7"
	CrypticAxe          = "7pa"
	CrypticSword        = "7ls"
	CrystalSword        = "crs"
	CrystallineGlobe    = "ob7"
	Cudgel              = "9cl"
	Cutlass             = "9sm"
	DacianFalx          = "9cm"
	Dagger              = "dgr"
	Decapitator         = "7bt"
	DecoyDagger         = "d33"
	DemonCrossbow       = "6rx"
	DemonHeart          = "obd"
	DevilStar           = "7mt"
	DiamondBow          = "6s7"
	DimensionalBlade    = "9cr"
	DimensionalShard    = "obf"
	Dirk                = "dir"
	DivineScepter       = "9ws"
	DoubleAxe           = "2ax"
	DoubleBow           = "8cb"
	DragonStone         = "ob5"
	EableOrb            = "ob1"
	EdgeBow             = "8sb"
	ElderStaff          = "6cs"
	EldritchOrb         = "obc"
	ElegantBlade        = "7sb"
	Espadon             = "92h"
	EttinAxe            = "72a"
	ExecutionerSword    = "9gd"
	ExplodingPotion     = "opm"
	Falcata             = "7ss"
	Falchion            = "flc"
	FangedKnife         = "7kr"
	Fascia              = "9xf"
	FeralAxe            = "7la"
	FeralClaws          = "7lw"
	Flail               = "fla"
	Flamberge           = "flb"
	FlangedMace         = "9ma"
	FlyingAxe           = "7ta"
	FlyingKnife         = "7tk"
	Francisca           = "9ta"
	FulmatingPotion     = "opl"
	Fuscina             = "9tr"
	GhostGlaive         = "7gl"
	GhostSpear          = "7st"
	GhostWand           = "7yw"
	GiantAxe            = "gix"
	GiantSword          = "gis"
	GiantThresher       = "7wc"
	Gidbinn             = "g33"
	Gladius             = "9ss"
	Glaive              = "glv"
	GloriousAxe         = "7gi"
	GlowingOrb          = "ob6"
	GnarledStaff        = "cst"
	GorgonCrossbow      = "6mx"
	GothicAxe           = "9ga"
	GothicBow           = "8lw"
	GothicStaff         = "8bs"
	GothicSword         = "9b9"
	GrandMatronBow      = "amc"
	GrandScepter        = "gsc"
	GraveWand           = "9gw"
	GreatAxe            = "gax"
	GreatBow            = "6cb"
	GreatMaul           = "gma"
	GreatPilum          = "9pi"
	GreatPoleaxe        = "7h7"
	GreatSword          = "gsd"
	GreaterClaws        = "9lw"
	GreaterTalons       = "9tw"
	GrimScythe          = "9wc"
	GrimWand            = "gwn"
	Halberd             = "hal"
	HandAxe             = "hax"
	HandScythe          = "9cs"
	Harpoon             = "9ts"
	Hatchet             = "9ha"
	HatchetHands        = "axf"
	HeavenlyStone       = "obb"
	HeavyCrossbow       = "hxb"
	HellforgeHammer     = "hfh"
	HighlandBlade       = "7cm"
	HolyWaterSprinkler  = "9qs"
	HoradricMalus       = "hdm"
	HoradricStaff       = "hst"
	HuntersBow          = "hbw"
	Hurlbat             = "9b8"
	HydraBow            = "6lw"
	HydraEdge           = "7fc"
	HyperionJavelin     = "7ja"
	HyperionSpear       = "7sr"
	JaggedStar          = "9mt"
	Javelin             = "jav"
	JoStaff             = "8ss"
	Katar               = "ktr"
	KhalimFlail         = "qf1"
	Knout               = "9fl"
	Kriss               = "kri"
	Lance               = "9p9"
	LargeAxe            = "lax"
	LegendSpike         = "7bl"
	LegendSword         = "72h"
	LegendaryMallet     = "7wh"
	LichWand            = "7bw"
	LightCrossbow       = "lxb"
	LochaberAxe         = "9b7"
	LongBattleBow       = "lbb"
	LongBow             = "lbw"
	LongSiegeBow        = "8l8"
	LongStaff           = "lst"
	LongSword           = "lsd"
	LongWarBow          = "lwb"
	Mace                = "mac"
	MaidenJavelin       = "am5"
	MaidenPike          = "am4"
	MaidenSpear         = "am3"
	Mancatcher          = "7br"
	MartelDeFar         = "9gm"
	MatriarchalBow      = "amb"
	MatriarchalPike     = "ame"
	MatriarchalSpear    = "amd"
	MatriarchalJavelin  = "amf"
	Maul                = "mau"
	MightyScepter       = "7sc"
	MilitaryAxe         = "9la"
	MilitaryPick        = "mpi"
	MithralPoint        = "7di"
	MorningStar         = "mst"
	MythicalSword       = "7wd"
	Naga                = "9wa"
	OgreAxe             = "7o7"
	OgreMaul            = "7m7"
	OilPotion           = "ops"
	Partizan            = "9pa"
	PelletBow           = "6lx"
	PetrifiedWand       = "9yw"
	PhaseBlade          = "7cr"
	Pike                = "pik"
	Pilum               = "pil"
	Poignard            = "9dg"
	Poleaxe             = "pax"
	PolishedWand        = "7wn"
	QuarterStaff        = "8ls"
	Quhab               = "9ar"
	RancidGasPotion     = "gps"
	RazorBow            = "8hb"
	ReflexBow           = "am2"
	ReinforcedMace      = "7ma"
	RepeatingCrossbow   = "rxb"
	Rondel              = "9di"
	RuneBow             = "8sw"
	RuneScepter         = "9sc"
	RuneStaff           = "8ws"
	RuneSword           = "9ls"
	RunicTalons         = "7tw"
	Sabre               = "sbr"
	SacredGlobe         = "ob2"
	Scepter             = "scp"
	Scimitar            = "scm"
	ScissorsKatar       = "skr"
	ScissorsQuhab       = "9qr"
	ScissorsSuwayyah    = "7qr"
	Scourge             = "7fl"
	Scythe              = "scy"
	SeraphRod           = "7qs"
	ShadowBow           = "6lb"
	Shamshir            = "9sb"
	Shillelah           = "6bs"
	ShortBattleBow      = "sbb"
	ShortBow            = "sbw"
	ShortSiegeBow       = "8s8"
	ShortSpear          = "ssp"
	ShortStaff          = "sst"
	ShortSword          = "ssd"
	ShortWarBow         = "swb"
	SiegeCrossbow       = "8mx"
	SilverEdgedAxe      = "7ba"
	Simbilan            = "9s9"
	SmallCrescent       = "7ax"
	SmokedSphere        = "ob3"
	SparklingBall       = "ob9"
	Spear               = "spr"
	Spetum              = "spt"
	Spiculum            = "9gl"
	SpiderBow           = "6sb"
	SpikedClub          = "spc"
	StaffOfTheKings     = "msf"
	StagBow             = "am1"
	Stalagmite          = "6ls"
	Stilleto            = "9bl"
	StranglingGasPotion = "gpl"
	StygianPike         = "7tr"
	StygianPilum        = "7pi"
	SuperKhalimFlail    = "qf2"
	Suwayyah            = "7ar"
	SwirlingCrystal     = "oba"
	Tabar               = "9bt"
	Thresher            = "7s8"
	ThrowingAxe         = "tax"
	ThrowingKnife       = "tkf"
	ThrowingSpear       = "tsp"
	ThunderMaul         = "7gm"
	Tomahawk            = "7ha"
	TombWand            = "9bw"
	Trident             = "tri"
	Truncheon           = "7cl"
	Tulwar              = "9fc"
	TuskSword           = "9gs"
	TwinAxe             = "92a"
	TwoHandedAxe        = "2hs"
	TyrantClub          = "7sp"
	UnearthedWand       = "7gw"
	VortexOrb           = "obe"
	Voulge              = "vou"
	WalkingStick        = "6ss"
	Wand                = "wnd"
	WarAxe              = "wax"
	WarClub             = "9m9"
	WarDart             = "9bk"
	WarFist             = "7xf"
	WarFork             = "9br"
	WarHammer           = "whm"
	WarJavelin          = "9ja"
	WarPike             = "7p7"
	WarScepter          = "wsp"
	WarScythe           = "wsc"
	WarSpear            = "9sr"
	WarSpike            = "7mp"
	WarStaff            = "wst"
	WarSword            = "wsd"
	WardBow             = "6sw"
	WingedAxe           = "7b8"
	WingedHarpoon       = "7ts"
	WingedKnife         = "7bk"
	WirtsLeg            = "leg"
	WristBlade          = "wrb"
	WristSpike          = "9wb"
	WristSword          = "7wb"
	Yari                = "9st"
	YewWand             = "ywn"
	Zweihander          = "9fb"
)
*/
type magicAttribute struct {
	ID     uint64
	Name   string
	Values []uint64
}

type magicalProperty struct {
	Bits []uint
	Bias uint64
	Name string
}

var magicalProperties = map[uint64]magicalProperty{
	0:  {Bits: []uint{7}, Bias: 32, Name: "+X to Strength"},
	1:  {Bits: []uint{7}, Bias: 32, Name: "+X to Energy"},
	2:  {Bits: []uint{7}, Bias: 32, Name: "+X to Dexterity"},
	3:  {Bits: []uint{7}, Bias: 32, Name: "+X to Vitality"},
	7:  {Bits: []uint{9}, Bias: 32, Name: "+X to Life"},
	9:  {Bits: []uint{8}, Bias: 32, Name: "+X to Mana"},
	11: {Bits: []uint{8}, Bias: 32, Name: "+X to Maximum Stamina"},
	16: {Bits: []uint{9}, Name: "+X% Enhanced Defense"},
	17: {Bits: []uint{9, 9}, Name: "+X% Enhanced Damage"},
	19: {Bits: []uint{10}, Name: "+X to Attack rating"},
	20: {Bits: []uint{6}, Name: "+X% Increased chance of blocking"},
	21: {Bits: []uint{6}, Name: "+X to Minimum damage"},
	22: {Bits: []uint{7}, Name: "+X to Maximum damage"},
	// Possible duplicate of id: 21, usually seen together
	23: {Bits: []uint{6}, Name: "+X to Minimum damage"},
	// Possible duplicate of id: 22, usually seen together
	24: {Bits: []uint{7}, Name: "+X to Maximum damage"},
	27: {Bits: []uint{8}, Name: "Regenerate Mana X%"},
	28: {Bits: []uint{8}, Name: "Heal Stamina X%"},
	31: {Bits: []uint{10}, Bias: 10, Name: "+X Defense"},
	32: {Bits: []uint{10}, Bias: 10, Name: "+X vs. Missile"},
	33: {Bits: []uint{10}, Bias: 10, Name: "+X vs. Melee"},
	34: {Bits: []uint{6}, Name: "Damage Reduced by X"},
	35: {Bits: []uint{6}, Name: "Magic Damage Reduced by X"},
	36: {Bits: []uint{8}, Name: "Damage Reduced by X%"},
	37: {Bits: []uint{8}, Name: "Magic Resist +X%"},
	38: {Bits: []uint{8}, Name: "+X% to Maximum Magic Resist"},
	39: {Bits: []uint{8}, Name: "Fire Resist +X%"},
	40: {Bits: []uint{5}, Name: "+X% to Maximum Fire Resist"},
	41: {Bits: []uint{8}, Name: "Lightning Resist +X%"},
	42: {Bits: []uint{5}, Name: "+X% to Maximum Lightning Resist"},
	43: {Bits: []uint{8}, Name: "Cold Resist +X%"},
	44: {Bits: []uint{5}, Name: "+X% to Maximum Cold Resist"},
	45: {Bits: []uint{8}, Name: "Poison Resist +X%"},
	46: {Bits: []uint{5}, Name: "+X% to Maximum Poison Resist"},
	48: {Bits: []uint{8, 8}, Name: "Adds X-Y Fire Damage"},
	50: {Bits: []uint{6, 9}, Name: "Adds X-Y Lightning Damage"},
	52: {Bits: []uint{6, 7}, Name: "Adds X-Y Magic Damage"},
	54: {Bits: []uint{6, 8, 8}, Name: "Adds X-Y Cold Damage"},
	57: {Bits: []uint{9, 9, 8}, Name: "Adds X-Y Poison Damage over Z Seconds"},
	60: {Bits: []uint{7}, Name: "X% Life stolen per hit"},
	73: {Bits: []uint{8}, Bias: 30, Name: "+X Maximum Durability"},
	74: {Bits: []uint{6}, Bias: 30, Name: "Replenish Life +X"},
	75: {Bits: []uint{7}, Bias: 20, Name: "Increase Maximum Durability X%"},
	76: {Bits: []uint{6}, Bias: 10, Name: "Increase Maximum Life X%"},
	77: {Bits: []uint{6}, Bias: 10, Name: "Increase Maximum Mana X%"},
	78: {Bits: []uint{7}, Name: "Attacker Takes Damage of X"},
	79: {Bits: []uint{9}, Bias: 100, Name: "X% Extra Gold from Monsters"},
	80: {Bits: []uint{8}, Bias: 100, Name: "X% Better Chance of Getting Magic Items"},
	81: {Bits: []uint{7}, Name: "Knockback"},
	83: {Bits: []uint{3}, Name: "+X to Amazon Skill Levels"},
	84: {Bits: []uint{3}, Name: "+X to Paladin Skill Levels"},
	85: {Bits: []uint{3}, Name: "+X to Necromancer Skill Levels"},
	86: {Bits: []uint{3}, Name: "+X to Sorceress Skill Levels"},
	87: {Bits: []uint{3}, Name: "+X to Barbarian Skill Levels"},
	89: {Bits: []uint{4}, Bias: 4, Name: "+X to Light Radius"},
	// This property is not displayed on the item, but its effect is to alter
	// the color of the ambient light.
	90: {Bits: []uint{5}, Name: "Ambient light"},
	// After subtracting the bias, this is usually a negative number.
	91:  {Bits: []uint{8}, Bias: 100, Name: "Requirements -X%"},
	93:  {Bits: []uint{7}, Bias: 20, Name: "X% Increased Attack Speed"},
	96:  {Bits: []uint{7}, Bias: 20, Name: "X% Faster Run/Walk"},
	99:  {Bits: []uint{7}, Bias: 20, Name: "X% Faster Hit Recovery"},
	102: {Bits: []uint{7}, Bias: 20, Name: "X% Faster Block Rate"},
	105: {Bits: []uint{7}, Bias: 20, Name: "X% Faster Cast Rate"},

	// These properties usually applied to class specific items,
	// first value selects the skill, the second determines how many
	// additional skill points are given.
	107: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	108: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	109: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	181: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	182: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	183: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	184: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	185: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	186: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},
	187: {Bits: []uint{9, 5}, Name: "+Y to spell X (char_class Only)"},

	110: {Bits: []uint{8}, Bias: 20, Name: "Poison Length Reduced by X%"},
	111: {Bits: []uint{7}, Bias: 20, Name: "Damage +X"},
	112: {Bits: []uint{7}, Name: "Hit Causes Monsters to Flee X%"},
	113: {Bits: []uint{7}, Name: "Hit Blinds Target +X"},
	114: {Bits: []uint{6}, Name: "X% Damage Taken Goes to Mana"},
	115: {Bits: []uint{1}, Name: "Ignore Target Defense"},
	// The value of the data field is always 1.
	116: {Bits: []uint{7}, Name: "X% Target Defense"},
	// The value of the data field is always 1.
	117: {Bits: []uint{7}, Name: "Prevent Monster Heal"},
	// The value of the data field is always 1.
	118: {Bits: []uint{1}, Name: "Half Freeze Duration"},
	119: {Bits: []uint{9}, Bias: 20, Name: "X% Bonus to Attack Rating"},
	120: {Bits: []uint{7}, Bias: 128, Name: "X to Monster Defense Per Hit"},
	121: {Bits: []uint{9}, Bias: 20, Name: "+X% Damage to Demons"},
	122: {Bits: []uint{9}, Bias: 20, Name: "+X% Damage to Undead"},
	123: {Bits: []uint{10}, Bias: 128, Name: "+X to Attack Rating against Demons"},
	124: {Bits: []uint{10}, Bias: 128, Name: "+X to Attack Rating against Undead"},
	126: {Bits: []uint{4}, Name: "+X to Fire Skills"},
	127: {Bits: []uint{3}, Name: "+X to All Skill Levels"},
	128: {Bits: []uint{5}, Name: "Attacker Takes Lightning Damage of X"},
	134: {Bits: []uint{5}, Name: "Freezes Target"},
	135: {Bits: []uint{7}, Name: "X% Chance of Open Wounds"},
	136: {Bits: []uint{7}, Name: "X% Change of Crushing Blow"},
	137: {Bits: []uint{7}, Name: "+X Kick Damage"},
	138: {Bits: []uint{7}, Name: "+X to Mana After Each Kill"},
	139: {Bits: []uint{7}, Name: "+X Life after each Demon Kill"},
	// Unknown property, shows up on Swordback Hold Spiked Shield.
	140: {Bits: []uint{7}, Name: "Unknown"},
	141: {Bits: []uint{7}, Name: "X% Deadly Strike"},
	142: {Bits: []uint{7}, Name: "Fire Absorb X%"},
	143: {Bits: []uint{7}, Name: "X Fire Absorb"},
	144: {Bits: []uint{7}, Name: "Lightning Absorb X%"},
	145: {Bits: []uint{7}, Name: "X Lightning Absorb"},
	146: {Bits: []uint{7}, Name: "Magic Absorb X%"},
	147: {Bits: []uint{7}, Name: "X Magic Absorb"},
	148: {Bits: []uint{7}, Name: "Cold Absorb X%"},
	149: {Bits: []uint{7}, Name: "X Cold Absorb"},
	150: {Bits: []uint{7}, Name: "Slows Target by X%"},
	151: {Bits: []uint{7}, Name: "Blessed Aim"},
	152: {Bits: []uint{7}, Name: "Defiance"},
	153: {Bits: []uint{1}, Name: "Cannot Be Frozen"},
	154: {Bits: []uint{7}, Name: "X% Slower Stamina Drain"},
	155: {Bits: []uint{7}, Name: "X% Chance to Reanimate Target"},
	156: {Bits: []uint{7}, Name: "Piercing Attack"},
	157: {Bits: []uint{7}, Name: "Fires Magic Arrows"},
	158: {Bits: []uint{7}, Name: "Fires Explosive Arrows or Bolts"},
	// Never seen this by itself, always accompanied by properties 21 and 23
	159: {Bits: []uint{6}, Name: "+X to Minimum Damage"},
	// Never seen this by itself, always accompanied by properties 22 and 24
	160: {Bits: []uint{7}, Name: "+X to Maximum Damage"},
	179: {Bits: []uint{3}, Name: "+X to Druid Skill Levels"},
	180: {Bits: []uint{3}, Name: "+X to Assassin Skill Levels"},

	// A skill set is a class specific skill tree id, e.g bow and crossbow skills,
	// traps or war cries. ID's are described below.
	188: {Bits: []uint{5, 5}, Name: "+Y to skill_set Skills (char_class Only)"},
	189: {Bits: []uint{5, 5}, Name: "+Y to skill_set Skills (char_class Only)"},
	190: {Bits: []uint{5, 5}, Name: "+Y to skill_set Skills (char_class Only)"},
	191: {Bits: []uint{5, 5}, Name: "+Y to skill_set Skills (char_class Only)"},
	192: {Bits: []uint{5, 5}, Name: "+Y to skill_set Skills (char_class Only)"},
	193: {Bits: []uint{5, 5}, Name: "+Y to skill_set Skills (char_class Only)"},

	194: {Bits: []uint{4}, Name: "Adds X extra sockets to the item"},

	// First order is spell id, % chance and level.
	195: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id on attack"},
	196: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id on attack"},
	197: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id on attack"},

	// First order is spell id, % chance and level.
	198: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id on striking"},
	199: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id on striking"},
	200: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id on striking"},

	// First order is spell id, % chance and level.
	201: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id when struck"},
	202: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id when struck"},
	203: {Bits: []uint{9, 5, 7}, Name: "Z% Chance to Cast Level Y skill_id when struck"},

	// First value selects the spell id, second value is level, third is remaining charges
	// and the last is the total number of charges.
	204: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	205: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	206: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	207: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	208: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	209: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	210: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	211: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	212: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},
	213: {Bits: []uint{9, 5, 8, 8}, Name: "Level X spell_id (Y/Z Charges)"},

	// All values based on character level are stored in eights, so take
	// the number divide by 8 and multiply by the character level and round down.
	214: {Bits: []uint{6}, Name: "+X to Defense (Based on Character Level)"},
	215: {Bits: []uint{6}, Name: "X% Enhanced Defense (Based on Character Level)"},
	216: {Bits: []uint{6}, Name: "+X to Life (Based on Character Level)"},
	217: {Bits: []uint{6}, Name: "+X to Mana (Based on Character Level)"},
	218: {Bits: []uint{6}, Name: "+X to Maximum Damage (Based on Character Level)"},
	219: {Bits: []uint{6}, Name: "X% Enhanced Maximum Damage (Based on Character Level)"},
	220: {Bits: []uint{6}, Name: "+X to Strength (Based on Character Level)"},
	221: {Bits: []uint{6}, Name: "+X to Dexterity (Based on Character Level)"},
	222: {Bits: []uint{6}, Name: "+X to Energy (Based on Character Level)"},
	223: {Bits: []uint{6}, Name: "+X to Vitality (Based on Character Level)"},
	224: {Bits: []uint{6}, Name: "+X to Attack Rating (Based on Character Level)"},
	225: {Bits: []uint{6}, Name: "X% Bonus to Attack Rating (Based on Character Level)"},
	226: {Bits: []uint{6}, Name: "+X Cold Damage (Based on Character Level)"},
	227: {Bits: []uint{6}, Name: "+X Fire Damage (Based on Character Level)"},
	228: {Bits: []uint{6}, Name: "+X Lightning Damage (Based on Character Level)"},
	229: {Bits: []uint{6}, Name: "+X Poison Damage (Based on Character Level)"},
	230: {Bits: []uint{6}, Name: "Cold Resist +X% (Based on Character Level)"},
	231: {Bits: []uint{6}, Name: "Fire Resist +X% (Based on Character Level)"},
	232: {Bits: []uint{6}, Name: "Lightning Resist +X% (Based on Character Level)"},
	233: {Bits: []uint{6}, Name: "Poison Resist +X% (Based on Character Level)"},
	234: {Bits: []uint{6}, Name: "+X Cold Absorb (Based on Character Level)"},
	235: {Bits: []uint{6}, Name: "+X Fire Absorb (Based on Character Level)"},
	236: {Bits: []uint{6}, Name: "+X Lightning Absorb (Based on Character Level)"},
	237: {Bits: []uint{6}, Name: "X Poison Absorb (Based on Character Level)"},
	238: {Bits: []uint{6}, Name: "Attacker Takes Damage of X (Based on Character Level)"},
	239: {Bits: []uint{6}, Name: "X% Extra Gold from Monsters (Based on Character Level)"},
	240: {Bits: []uint{6}, Name: "X% Better Chance of Getting Magic Items (Based on Character Level)"},
	241: {Bits: []uint{6}, Name: "Heal Stamina Plus X% (Based on Character Level)"},
	243: {Bits: []uint{6}, Name: "X% Damage to Demons (Based on Character Level)"},
	244: {Bits: []uint{6}, Name: "X% Damage to Undead (Based on Character Level)"},
	245: {Bits: []uint{6}, Name: "+X to Attack Rating against Demons (Based on Character Level)"},
	246: {Bits: []uint{6}, Name: "+X to Attack Rating against Undead (Based on Character Level)"},
	247: {Bits: []uint{6}, Name: "X% Chance of Crushing Blow (Based on Character Level)"},
	248: {Bits: []uint{6}, Name: "X% Chance of Open Wounds (Based on Character Level)"},
	249: {Bits: []uint{6}, Name: "+X Kick Damage (Based on Character Level)"},
	250: {Bits: []uint{6}, Name: "X% to Deadly Strike (Based on Character Level)"},

	// The value of the data field is not actually a time period, but a frequency in terms
	// of the number of times durability is repaired over a period of 100 seconds.
	// For example, if the value is 5, then this property repairs 1 durability in 100 / 5 = 20 seconds.
	252: {Bits: []uint{5}, Name: "Repairs 1 Durability in X Seconds"},

	// As in the previous property, the value of the data field is a frequency in terms of the number
	// replenished over a period of 100 seconds. For example if the value is 4, then this property
	// replenishes 1 item in 100 / 4 = 25 seconds.
	253: {Bits: []uint{5}, Name: "Replenishes Quantity"},

	// Number of additional items beyond the base limit, for example if the base
	// is 50 and additional is 30, then the total is 50 + 30.
	254: {Bits: []uint{8}, Name: "Increased Stack Size"},

	// These are some weird values, no idea if they are used in game or not.
	// But these values change depending on the time of day in the game.
	// The format of the bit fields are the same in all cases, the first 2 bits
	// specifies the the of time when the value is at its maximum.
	//
	// The second and third are respectively the minimum and maximum values of the property.
	// The maximum value at the time specified and the minimum at the opposite.

	// TODO: Add ids 268 - 303 if they prove to exist.
}

// All item types that contain the quantity bits will exist in here,
// we'll use this when reading items to make sure we only read quantity bits
// when they exist, or we'll ruin the rest of the bit offsets for the item.
var quantityMap = map[string]bool{
	Ring:         false,
	CompositeBow: false,
}
