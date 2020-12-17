package d2s

// Item describes any type of item and all it's data.
type Item struct {
	Identified          uint64             `json:"identified"`
	Socketed            uint64             `json:"socketed"`
	New                 uint64             `json:"new"`
	IsEar               uint64             `json:"is_ear"`
	StarterItem         uint64             `json:"starter_item"`
	SimpleItem          uint64             `json:"simple_item"`
	Ethereal            uint64             `json:"ethereal"`
	Personalized        uint64             `json:"personalized"`
	PersonalizedName    string             `json:"personalized_name,omitempty"`
	GivenRuneword       uint64             `json:"given_runeword"`
	Version             uint64             `json:"version"`
	LocationID          uint64             `json:"location_id"`
	EquippedID          uint64             `json:"equipped_id,omitempty"`
	PositionX           uint64             `json:"position_x"`
	PositionY           uint64             `json:"position_y"`
	AltPositionID       uint64             `json:"alt_position_id"`
	Type                string             `json:"type"`
	TypeID              uint64             `json:"type_id"`
	TypeName            string             `json:"type_name"`
	NrOfItemsInSockets  uint64             `json:"nr_of_items_in_sockets"`
	ID                  uint64             `json:"id"`
	Level               uint64             `json:"level"`
	Quality             uint64             `json:"quality"`
	MultiplePictures    uint64             `json:"multiple_pictures"`
	PictureID           uint64             `json:"picture_id"`
	ClassSpecific       uint64             `json:"class_specific"`
	LowQualityID        uint64             `json:"low_quality_id"`
	Timestamp           uint64             `json:"timestamp"`
	EarAttributes       earAttributes      `json:"ear_attributes"`
	DefenseRating       int64              `json:"defense_rating,omitempty"`
	MaxDurability       uint64             `json:"max_durability,omitempty"`
	CurrentDurability   uint64             `json:"current_durability,omitempty"`
	TotalNrOfSockets    uint64             `json:"total_nr_of_sockets"`
	Quantity            uint64             `json:"quantity,omitempty"`
	MagicPrefix         uint64             `json:"magic_prefix,omitempty"`
	MagicPrefixName     string             `json:"magic_prefix_name,omitempty"`
	MagicSuffix         uint64             `json:"magic_suffix,omitempty"`
	MagicSuffixName     string             `json:"magic_suffix_name,omitempty"`
	RunewordID          uint64             `json:"runeword_id,omitempty"`
	RunewordName        string             `json:"runeword_name,omitempty"`
	RunewordAttributes  []magicAttribute   `json:"runeword_attributes"`
	SetID               uint64             `json:"set_id,omitempty"`
	SetName             string             `json:"set_name,omitempty"`
	SetListCount        uint64             `json:"set_list_count"`
	SetAttributes       [][]magicAttribute `json:"set_attributes"`
	SetAttributesNumReq []uint             `json:"set_attributes_num_req,omitempty"`
	SetAttributesIDsReq []uint64           `json:"set_attributes_ids_req,omitempty"`
	RareName            string             `json:"rare_name,omitempty"`
	RareName2           string             `json:"rare_name2,omitempty"`
	MagicalNameIDs      []uint64           `json:"magical_name_ids,omitempty"`
	UniqueID            uint64             `json:"unique_id,omitempty"`
	UniqueName          string             `json:"unique_name,omitempty"`
	MagicAttributes     []magicAttribute   `json:"magic_attributes"`
	SocketedItems       []Item             `json:"socketed_items"`
	BaseDamage          *weaponDamage      `json:"base_damage,omitempty"`
}

func (i *Item) getTypeID() uint64 {
	if _, ok := armorCodes[i.Type]; ok {
		return armor
	}

	if _, ok := shieldCodes[i.Type]; ok {
		return shield
	}

	if _, ok := weaponCodes[i.Type]; ok {
		return weapon
	}

	return other
}

// Item code value constant used as an internal reference or "ID".
const (
	// Armor
	AlphaHelm         string = "dr6"
	AncientArmor      string = "aar"
	Antlers           string = "dr3"
	ArchonPlate       string = "utp"
	Armet             string = "ulm"
	AssaultHelmet     string = "ba4"
	AvengerGuard      string = "ba5"
	BalrogSkin        string = "upl"
	Basinet           string = "xhl"
	BattleBelt        string = "ztb"
	BattleBoots       string = "xtb"
	BattleGauntlets   string = "xtg"
	Belt              string = "mbl"
	BloodSpirit       string = "drb"
	BoneHelm          string = "bhm"
	BoneVisage        string = "uh9"
	Boneweave         string = "uhn"
	BoneweaveBoots    string = "umb"
	Bracers           string = "mgl"
	BrambleMitts      string = "ulg"
	BreastPlate       string = "brs"
	Cap               string = "cap"
	CarnageHelm       string = "bab"
	Casque            string = "xlm"
	ChainBoots        string = "mbt"
	ChainMail         string = "chn"
	ChaosArmor        string = "xul"
	Circlet           string = "ci0"
	ColossusGirdle    string = "uhc"
	ConquerorCrown    string = "bae"
	Corona            string = "urn"
	Coronet           string = "ci1"
	Crown             string = "crn"
	CrusaderGauntlets string = "utg"
	Cuirass           string = "xrs"
	DeathMask         string = "xsk"
	Demonhead         string = "usk"
	DemonhideArmor    string = "xla"
	DemonhideBoots    string = "xlb"
	DemonhideGloves   string = "xlg"
	DemonhideSash     string = "zlb"
	DestroyerHelm     string = "bad"
	Diadem            string = "ci3"
	DiamondMail       string = "ung"
	DreamSpirit       string = "drf"
	DuskShroud        string = "uui"
	EarthSpirit       string = "drd"
	EmbossedPlate     string = "xth"
	FalconMask        string = "dr4"
	FangedHelm        string = "ba2"
	FieldPlate        string = "fld"
	FullHelm          string = "fhl"
	PlateMail         string = "plt"
	FullPlateMail     string = "ful"
	FuryVisor         string = "bac"
	Gauntlets         string = "hgl"
	GhostArmor        string = "xui"
	GiantConch        string = "uhl"
	Girdle            string = "hbl"
	Gloves            string = "lgl"
	GothicPlate       string = "gth"
	GrandCrown        string = "xrn"
	GreatHauberk      string = "urs"
	GreatHelm         string = "ghm"
	GriffonHeadress   string = "dr7"
	GrimHelm          string = "xh9"
	GuardianCrown     string = "baf"
	HardLeatherArmor  string = "hla"
	HawkHelm          string = "dr2"
	HeavyBelt         string = "tbl"
	HeavyBoots        string = "vbt"
	HeavyBracers      string = "xmg"
	HeavyGloves       string = "vgl"
	HellforgedPlate   string = "ult"
	Helm              string = "hlm"
	HornedHelm        string = "ba3"
	HuntersGuise      string = "dr8"
	Hydraskull        string = "ukp"
	Hyperion          string = "urg"
	JawboneCap        string = "ba1"
	JawboneVisor      string = "ba6"
	KrakenShell       string = "uld"
	LacqueredPlate    string = "uth"
	LeatherArmor      string = "lea"
	LeatherBoots      string = "lbt"
	LightBelt         string = "vbl"
	LightGauntlets    string = "tgl"
	LightPlate        string = "ltp"
	LightPlatedBoots  string = "tbt"
	LinkedMail        string = "xng"
	LionHelm          string = "ba7"
	LoricatedMail     string = "ucl"
	MagePlate         string = "xtp"
	Mask              string = "msk"
	MeshArmor         string = "xhn"
	MeshBelt          string = "zmb"
	MeshBoots         string = "xmb"
	MirroredBoots     string = "utb"
	MithrilCoil       string = "umc"
	MyrmidonGreaves   string = "uhb"
	OgreGauntlets     string = "uhg"
	OrnatePlate       string = "xar"
	PlateBoots        string = "hbt"
	QuiltedArmor      string = "qui"
	RageMask          string = "ba8"
	RingMail          string = "rng"
	RussetArmor       string = "xpl"
	SacredArmor       string = "uar"
	SacredFeathers    string = "dr9"
	Sallet            string = "xkp"
	Sash              string = "lbl"
	SavageHelmet      string = "ba9"
	ScaleMail         string = "scl"
	ScarabHusk        string = "ula"
	ScarabshellBoots  string = "uvb"
	SerpentskinArmor  string = "xea"
	ShadowPlate       string = "uul"
	Shako             string = "uap"
	SharkskinBelt     string = "zvb"
	SharkskinBoots    string = "xvb"
	SharkskinGloves   string = "xvg"
	SharktoothArmor   string = "xld"
	SkullCap          string = "skp"
	SkySpirit         string = "dre"
	SlayerGuard       string = "baa"
	SpiderwebSash     string = "ulc"
	SpiredHelm        string = "uhm"
	SpiritMask        string = "dr5"
	SplintMail        string = "spl"
	StuddedLeather    string = "stu"
	SunSpirit         string = "drc"
	TemplarCoat       string = "xlt"
	Tiara             string = "ci2"
	TigulatedMail     string = "xcl"
	TotemicMask       string = "dra"
	TrellisedArmor    string = "xtu"
	TrollBelt         string = "utc"
	Vambraces         string = "umg"
	VampireboneGloves string = "uvg"
	VampirefangBelt   string = "uvc"
	WarBelt           string = "zhb"
	WarBoots          string = "xhb"
	WarGauntlets      string = "xhg"
	WarHat            string = "xap"
	WingedHelm        string = "xhm"
	WireFleece        string = "utu"
	WolfHead          string = "dr1"
	Wyrmhide          string = "uea"
	WyrmhideBoots     string = "ulb"

	// Misc
	Amethyst                      string = "gsv"
	AmnRune                       string = "r11"
	Amulet                        string = "amu"
	AntidotePotion                string = "yps"
	Arrows                        string = "aqv"
	BaalsEye                      string = "bey"
	BarkScroll                    string = "bks"
	BerRune                       string = "r30"
	Bolts                         string = "cqv"
	BookofSkill                   string = "ass"
	BurningEssenceofTerror        string = "bet"
	ChamRune                      string = "r32"
	ChargedEssenceofHatred        string = "ceh"
	GrandCharm                    string = "cm3"
	LargeCharm                    string = "cm2"
	SmallCharm                    string = "cm1"
	ChippedAmethyst               string = "gcv"
	ChippedDiamond                string = "gcw"
	ChippedEmerald                string = "gcg"
	ChippedRuby                   string = "gcr"
	ChippedSapphire               string = "gcb"
	ChippedSkull                  string = "skc"
	ChippedTopaz                  string = "gcy"
	DecipheredBarkScroll          string = "bkd"
	DiablosHorn                   string = "dhn"
	Diamond                       string = "gsw"
	DolRune                       string = "r14"
	ElRune                        string = "r01"
	EldRune                       string = "r02"
	Elixir                        string = "elx"
	Emerald                       string = "gsg"
	EthRune                       string = "r05"
	FalRune                       string = "r19"
	FesteringEssenceofDestruction string = "fed"
	FlawedAmethyst                string = "gfv"
	FlawedDiamond                 string = "gfw"
	FlawedEmerald                 string = "gfg"
	FlawedRuby                    string = "gfr"
	FlawedSapphire                string = "gfb"
	FlawedSkull                   string = "skf"
	FlawedTopaz                   string = "gfy"
	FlawlessAmethyst              string = "gzv"
	FlawlessDiamond               string = "glw"
	FlawlessEmerald               string = "glg"
	FlawlessRuby                  string = "glr"
	FlawlessSapphire              string = "glb"
	FlawlessSkull                 string = "skl"
	FlawlessTopaz                 string = "gly"
	FullHealingPotion             string = "hpf"
	FullManaPotion                string = "mpf"
	FullRejuvenationPotion        string = "rvl"
	Gold                          string = "gld"
	GoldBird                      string = "g34"
	GreaterHealingPotion          string = "hp5"
	GreaterManaPotion             string = "mp5"
	GulRune                       string = "r25"
	HealingPotion                 string = "hp3"
	HealingPotionMid              string = "hpo"
	HelRune                       string = "r15"
	Herb                          string = "hrb"
	HoradricCube                  string = "box"
	IdentifyBook                  string = "ibk"
	IdentifyScroll                string = "isc"
	IoRune                        string = "r16"
	IstRune                       string = "r24"
	IthRune                       string = "r06"
	JadeFigurine                  string = "j34"
	JahRune                       string = "r31"
	Jewel                         string = "jew"
	KeyofDestruction              string = "pk3"
	KeyofHate                     string = "pk2"
	KeyofTerror                   string = "pk1"
	KhalimsBrain                  string = "qbr"
	KhalimsEye                    string = "qey"
	KhalimsHeart                  string = "qhr"
	KoRune                        string = "r18"
	LamEsensTome                  string = "bbb"
	LargeBluePotion               string = "bpl"
	LargeRedPotion                string = "rpl"
	LemRune                       string = "r20"
	LesserHealingPotion           string = "hp1"
	LesserManaPotion              string = "mp1"
	LightHealingPotion            string = "hp2"
	LightManaPotion               string = "mp2"
	LoRune                        string = "r28"
	LumRune                       string = "r17"
	Maguffin                      string = "ice"
	MalRune                       string = "r23"
	ManaPotion                    string = "mp3"
	ManaPotionMid                 string = "mpo"
	MephistosBrain                string = "mbr"
	MephistoKey                   string = "luv"
	MephistoSoulStone             string = "mss"
	NefRune                       string = "r04"
	OhmRune                       string = "r27"
	OrtRune                       string = "r09"
	PerfectAmethyst               string = "gpv"
	PerfectDiamond                string = "gpw"
	PerfectEmerald                string = "gpg"
	PerfectRuby                   string = "gpr"
	PerfectSapphire               string = "gpb"
	PerfectSkull                  string = "skz"
	PerfectTopaz                  string = "gpy"
	PlayerEar                     string = "ear"
	PulRune                       string = "r21"
	RalRune                       string = "r08"
	RejuvenationPotion            string = "rvs"
	Ring                          string = "rin"
	Ruby                          string = "gsr"
	Sapphire                      string = "gsb"
	Scroll                        string = "0sc"
	ScrollofHoradric              string = "tr1"
	ScrollofMalah                 string = "tr2"
	PotionofLife                  string = "xyz"
	ShaelRune                     string = "r13"
	SkeletonKey                   string = "key"
	Skull                         string = "sku"
	SmallBluePotion               string = "bps"
	SmallRedPotion                string = "rps"
	SolRune                       string = "r12"
	StaminaPotion                 string = "vps"
	StandardofHeroes              string = "std"
	StrongHealingPotion           string = "hp4"
	StrongManaPotion              string = "mp4"
	SurRune                       string = "r29"
	TalRune                       string = "r07"
	ThawingPotion                 string = "wms"
	ThulRune                      string = "r10"
	TirRune                       string = "r03"
	TokenofAbsolution             string = "toa"
	Topaz                         string = "gsy"
	Torch                         string = "tch"
	TownPortalBook                string = "tbk"
	TownPortalScroll              string = "tsc"
	TwistedEssenceofSuffering     string = "tes"
	UmRune                        string = "r22"
	VexRune                       string = "r26"
	ViperAmulet                   string = "vip"
	ZodRune                       string = "r33"

	// Shields
	Aegis            string = "uow"
	AerinShield      string = "pa4"
	AkaranRondache   string = "pa7"
	AkaranTarge      string = "pa6"
	KurastShield     string = "pad"
	AncientShield    string = "xts"
	BarbedShield     string = "xpk"
	BladeBarrier     string = "upk"
	BloodlordSkull   string = "nef"
	BoneShield       string = "bsh"
	Buckler          string = "buc"
	CantorTrophy     string = "ne9"
	CrownShield      string = "pa5"
	Defender         string = "xuc"
	DemonHead        string = "ne5"
	DragonShield     string = "xit"
	FetishTrophy     string = "ne7"
	GargoyleHead     string = "ne4"
	GothicShield     string = "gts"
	GrimShield       string = "xsh"
	GildedShield     string = "pa9"
	HeirophantTrophy string = "nea"
	Heater           string = "uuc"
	HellspawnSkull   string = "neg"
	HeraldicShield   string = "pa3"
	KiteShield       string = "kit"
	LargeShield      string = "lrg"
	Luna             string = "uml"
	MinionSkull      string = "neb"
	Monarch          string = "uit"
	MummifiedTrophy  string = "ne6"
	OverseerSkull    string = "ned"
	Pavise           string = "xow"
	PreservedHead    string = "ne1"
	ProtectorShield  string = "pa8"
	Rondache         string = "pa2"
	RoundShield      string = "xml"
	RoyalShield      string = "paa"
	SacredRondache   string = "pac"
	SacredTarge      string = "pab"
	Scutum           string = "xrg"
	SextonTrophy     string = "ne8"
	SmallShield      string = "sml"
	SpikedShield     string = "spk"
	SuccubusSkull    string = "nee"
	Targe            string = "pa1"
	TowerShield      string = "tow"
	TrollNest        string = "ush"
	UnravellerHead   string = "ne3"
	VortexShield     string = "paf"
	Ward             string = "uts"
	ZakarumShield    string = "pae"
	ZombieHead       string = "ne2"

	// Weapons
	AncientAxe          string = "9gi"
	AncientSword        string = "9wd"
	Arbalest            string = "8lx"
	ArchonStaff         string = "6ws"
	AshwoodBow          string = "am6"
	Ataghan             string = "7sm"
	Axe                 string = "axe"
	BalancedAxe         string = "bal"
	BalancedKnife       string = "bkf"
	Balista             string = "8hx"
	BalrogBlade         string = "7gs"
	BalrogSpear         string = "7s7"
	BarbedClub          string = "9sp"
	Bardiche            string = "bar"
	BastardSword        string = "bsw"
	BattleAxe           string = "btx"
	BattleCestus        string = "7cs"
	BattleDart          string = "9tk"
	BattleHammer        string = "9wh"
	BattleScythe        string = "9s8"
	BattleStaff         string = "bst"
	BattleSword         string = "9bs"
	BeardedAxe          string = "9ba"
	BecDeCorbin         string = "9h9"
	BerserkerAxe        string = "7wa"
	Bill                string = "9vo"
	Blade               string = "bld"
	BladeBow            string = "6hb"
	BladeTalons         string = "btl"
	BoneKnife           string = "7dg"
	BoneWand            string = "bwn"
	Brandistock         string = "brn"
	BroadAxe            string = "bax"
	BroadSword          string = "bsd"
	BurntWand           string = "9wn"
	Caduceus            string = "7ws"
	CedarBow            string = "8lb"
	CedarStaff          string = "8cs"
	CeremonialBow       string = "am7"
	CeremonialJavelin   string = "ama"
	CeremonialPike      string = "am9"
	CeremonialSpear     string = "am8"
	Cestus              string = "ces"
	ChampionAxe         string = "7ga"
	ChampionSword       string = "7b7"
	ChokingGasPotion    string = "gpm"
	ChuKoNu             string = "8rx"
	Cinquedeas          string = "9kr"
	ClaspedOrb          string = "ob4"
	Claws               string = "clw"
	Claymore            string = "clm"
	Cleaver             string = "9ax"
	CloudySphere        string = "ob8"
	Club                string = "clb"
	ColossalSword       string = "7fb"
	ColossalBlade       string = "7gd"
	ColossusCrossbow    string = "6hx"
	ColossusVoulge      string = "7vo"
	CompositeBow        string = "cbw"
	ConquestSword       string = "7bs"
	Crossbow            string = "mxb"
	Crowbill            string = "9mp"
	CrusaderBow         string = "6l7"
	CrypticAxe          string = "7pa"
	CrypticSword        string = "7ls"
	CrystalSword        string = "crs"
	CrystallineGlobe    string = "ob7"
	Cudgel              string = "9cl"
	Cutlass             string = "9sm"
	DacianFalx          string = "9cm"
	Dagger              string = "dgr"
	Decapitator         string = "7bt"
	DecoyDagger         string = "d33"
	DemonCrossbow       string = "6rx"
	DemonHeart          string = "obd"
	DevilStar           string = "7mt"
	DiamondBow          string = "6s7"
	DimensionalBlade    string = "9cr"
	DimensionalShard    string = "obf"
	Dirk                string = "dir"
	DivineScepter       string = "9ws"
	DoubleAxe           string = "2ax"
	DoubleBow           string = "8cb"
	DragonStone         string = "ob5"
	EableOrb            string = "ob1"
	EdgeBow             string = "8sb"
	ElderStaff          string = "6cs"
	EldritchOrb         string = "obc"
	ElegantBlade        string = "7sb"
	Espadon             string = "92h"
	EttinAxe            string = "72a"
	ExecutionerSword    string = "9gd"
	ExplodingPotion     string = "opm"
	Falcata             string = "7ss"
	Falchion            string = "flc"
	FangedKnife         string = "7kr"
	Fascia              string = "9xf"
	FeralAxe            string = "7la"
	FeralClaws          string = "7lw"
	Flail               string = "fla"
	Flamberge           string = "flb"
	FlangedMace         string = "9ma"
	FlyingAxe           string = "7ta"
	FlyingKnife         string = "7tk"
	Francisca           string = "9ta"
	FulmatingPotion     string = "opl"
	Fuscina             string = "9tr"
	GhostGlaive         string = "7gl"
	GhostSpear          string = "7st"
	GhostWand           string = "7yw"
	GiantAxe            string = "gix"
	GiantSword          string = "gis"
	GiantThresher       string = "7wc"
	Gidbinn             string = "g33"
	Gladius             string = "9ss"
	Glaive              string = "glv"
	GloriousAxe         string = "7gi"
	GlowingOrb          string = "ob6"
	GnarledStaff        string = "cst"
	GorgonCrossbow      string = "6mx"
	GothicAxe           string = "9ga"
	GothicBow           string = "8lw"
	GothicStaff         string = "8bs"
	GothicSword         string = "9b9"
	GrandMatronBow      string = "amc"
	GrandScepter        string = "gsc"
	GraveWand           string = "9gw"
	GreatAxe            string = "gax"
	GreatBow            string = "6cb"
	GreatMaul           string = "gma"
	GreatPilum          string = "9pi"
	GreatPoleaxe        string = "7h7"
	GreatSword          string = "gsd"
	GreaterClaws        string = "9lw"
	GreaterTalons       string = "9tw"
	GrimScythe          string = "9wc"
	GrimWand            string = "gwn"
	Halberd             string = "hal"
	HandAxe             string = "hax"
	HandScythe          string = "9cs"
	Harpoon             string = "9ts"
	Hatchet             string = "9ha"
	HatchetHands        string = "axf"
	HeavenlyStone       string = "obb"
	HeavyCrossbow       string = "hxb"
	HellforgeHammer     string = "hfh"
	HighlandBlade       string = "7cm"
	HolyWaterSprinkler  string = "9qs"
	HoradricMalus       string = "hdm"
	HoradricStaff       string = "hst"
	HuntersBow          string = "hbw"
	Hurlbat             string = "9b8"
	HydraBow            string = "6lw"
	HydraEdge           string = "7fc"
	HyperionJavelin     string = "7ja"
	HyperionSpear       string = "7sr"
	JaggedStar          string = "9mt"
	Javelin             string = "jav"
	JoStaff             string = "8ss"
	Katar               string = "ktr"
	KhalimFlail         string = "qf1"
	Knout               string = "9fl"
	Kriss               string = "kri"
	Lance               string = "9p9"
	LargeAxe            string = "lax"
	LegendSpike         string = "7bl"
	LegendSword         string = "72h"
	LegendaryMallet     string = "7wh"
	LichWand            string = "7bw"
	LightCrossbow       string = "lxb"
	LochaberAxe         string = "9b7"
	LongBattleBow       string = "lbb"
	LongBow             string = "lbw"
	LongSiegeBow        string = "8l8"
	LongStaff           string = "lst"
	LongSword           string = "lsd"
	LongWarBow          string = "lwb"
	Mace                string = "mac"
	MaidenJavelin       string = "am5"
	MaidenPike          string = "am4"
	MaidenSpear         string = "am3"
	Mancatcher          string = "7br"
	MarteldeFer         string = "9gm"
	MatriarchalBow      string = "amb"
	MatriarchalPike     string = "ame"
	MatriarchalSpear    string = "amd"
	MatriarchalJavelin  string = "amf"
	Maul                string = "mau"
	MightyScepter       string = "7sc"
	MilitaryAxe         string = "9la"
	MilitaryPick        string = "mpi"
	MithralPoint        string = "7di"
	MorningStar         string = "mst"
	MythicalSword       string = "7wd"
	Naga                string = "9wa"
	OgreAxe             string = "7o7"
	OgreMaul            string = "7m7"
	OilPotion           string = "ops"
	Partizan            string = "9pa"
	PelletBow           string = "6lx"
	PetrifiedWand       string = "9yw"
	PhaseBlade          string = "7cr"
	Pike                string = "pik"
	Pilum               string = "pil"
	Poignard            string = "9dg"
	Poleaxe             string = "pax"
	PolishedWand        string = "7wn"
	QuarterStaff        string = "8ls"
	Quhab               string = "9ar"
	RancidGasPotion     string = "gps"
	RazorBow            string = "8hb"
	ReflexBow           string = "am2"
	ReinforcedMace      string = "7ma"
	RepeatingCrossbow   string = "rxb"
	Rondel              string = "9di"
	RuneBow             string = "8sw"
	RuneScepter         string = "9sc"
	RuneStaff           string = "8ws"
	RuneSword           string = "9ls"
	RunicTalons         string = "7tw"
	Sabre               string = "sbr"
	SacredGlobe         string = "ob2"
	Scepter             string = "scp"
	Scimitar            string = "scm"
	ScissorsKatar       string = "skr"
	ScissorsQuhab       string = "9qr"
	ScissorsSuwayyah    string = "7qr"
	Scourge             string = "7fl"
	Scythe              string = "scy"
	SeraphRod           string = "7qs"
	ShadowBow           string = "6lb"
	Shamshir            string = "9sb"
	Shillelah           string = "6bs"
	ShortBattleBow      string = "sbb"
	ShortBow            string = "sbw"
	ShortSiegeBow       string = "8s8"
	ShortSpear          string = "ssp"
	ShortStaff          string = "sst"
	ShortSword          string = "ssd"
	ShortWarBow         string = "swb"
	SiegeCrossbow       string = "8mx"
	SilverEdgedAxe      string = "7ba"
	Simbilan            string = "9s9"
	SmallCrescent       string = "7ax"
	SmokedSphere        string = "ob3"
	SparklingBall       string = "ob9"
	Spear               string = "spr"
	Spetum              string = "spt"
	Spiculum            string = "9gl"
	SpiderBow           string = "6sb"
	SpikedClub          string = "spc"
	StaffOfTheKings     string = "msf"
	StagBow             string = "am1"
	Stalagmite          string = "6ls"
	Stilleto            string = "9bl"
	StranglingGasPotion string = "gpl"
	StygianPike         string = "7tr"
	StygianPilum        string = "7pi"
	SuperKhalimFlail    string = "qf2"
	Suwayyah            string = "7ar"
	SwirlingCrystal     string = "oba"
	Tabar               string = "9bt"
	Thresher            string = "7s8"
	ThrowingAxe         string = "tax"
	ThrowingKnife       string = "tkf"
	ThrowingSpear       string = "tsp"
	ThunderMaul         string = "7gm"
	Tomahawk            string = "7ha"
	TombWand            string = "9bw"
	Trident             string = "tri"
	Truncheon           string = "7cl"
	Tulwar              string = "9fc"
	TuskSword           string = "9gs"
	TwinAxe             string = "92a"
	TwoHandedSword      string = "2hs"
	TyrantClub          string = "7sp"
	UnearthedWand       string = "7gw"
	VortexOrb           string = "obe"
	Voulge              string = "vou"
	WalkingStick        string = "6ss"
	Wand                string = "wnd"
	WarAxe              string = "wax"
	WarClub             string = "9m9"
	WarDart             string = "9bk"
	WarFist             string = "7xf"
	WarFork             string = "9br"
	WarHammer           string = "whm"
	WarJavelin          string = "9ja"
	WarPike             string = "7p7"
	WarScepter          string = "wsp"
	WarScythe           string = "wsc"
	WarSpear            string = "9sr"
	WarSpike            string = "7mp"
	WarStaff            string = "wst"
	WarSword            string = "wsd"
	WardBow             string = "6sw"
	WingedAxe           string = "7b8"
	WingedHarpoon       string = "7ts"
	WingedKnife         string = "7bk"
	WirtsLeg            string = "leg"
	WristBlade          string = "wrb"
	WristSpike          string = "9wb"
	WristSword          string = "7wb"
	Yari                string = "9st"
	YewWand             string = "ywn"
	Zweihander          string = "9fb"
)

// Attributes that only exists on a player ear.
type earAttributes struct {
	Class uint64 `json:"class"`
	Level uint64 `json:"level"`
	Name  string `json:"name"`
}

// Note the values array is of the type int64, this is because some properties
// contain negative values, such as - % requirements.
type magicAttribute struct {
	ID     uint64  `json:"id"`
	Name   string  `json:"name"`
	Values []int64 `json:"values"`
}

type magicalProperty struct {
	Bits []uint
	Bias uint64
	Name string
}

type weaponDamage struct {
	Min    int `json:"min,omitempty"`
	Max    int `json:"max,omitempty"`
	TwoMin int `json:"twohand_min,omitempty"`
	TwoMax int `json:"twohand_max,omitempty"`
}

var weaponDamageMap = map[string]weaponDamage{
	//Axes
	"hax": {Min: 3, Max: 6},
	"axe": {Min: 4, Max: 11},
	"2ax": {Min: 5, Max: 13},
	"mpi": {Min: 7, Max: 11},
	"wax": {Min: 10, Max: 18},
	"lax": {Min: 6, Max: 13},
	"bax": {Min: 10, Max: 18},
	"btx": {Min: 12, Max: 32},
	"gax": {Min: 9, Max: 30},
	"gix": {Min: 22, Max: 45},
	"9ha": {Min: 10, Max: 21},
	"9ax": {Min: 10, Max: 33},
	"92a": {Min: 13, Max: 38},
	"9mp": {Min: 14, Max: 34},
	"9wa": {Min: 16, Max: 45},
	"9la": {Min: 14, Max: 34},
	"9ba": {Min: 21, Max: 49},
	"9bt": {Min: 24, Max: 77},
	"9ga": {Min: 18, Max: 70},
	"9gi": {Min: 43, Max: 85},
	"7ha": {Min: 33, Max: 58},
	"7ax": {Min: 38, Max: 60},
	"72a": {Min: 33, Max: 66},
	"7mp": {Min: 30, Max: 48},
	"7wa": {Min: 24, Max: 71},
	"7la": {Min: 25, Max: 123},
	"7ba": {Min: 62, Max: 110},
	"7bt": {Min: 49, Max: 137},
	"7ga": {Min: 59, Max: 94},
	"7gi": {Min: 60, Max: 124},

	//Maces
	"clb": {Min: 1, Max: 6},
	"spc": {Min: 5, Max: 8},
	"mac": {Min: 3, Max: 10},
	"mst": {Min: 7, Max: 16},
	"fla": {Min: 1, Max: 24},
	"whm": {Min: 19, Max: 29},
	"mau": {Min: 30, Max: 43},
	"gma": {Min: 38, Max: 58},
	"9cl": {Min: 6, Max: 21},
	"9sp": {Min: 13, Max: 25},
	"9ma": {Min: 15, Max: 23},
	"9mt": {Min: 20, Max: 31},
	"9fl": {Min: 13, Max: 35},
	"9wh": {Min: 35, Max: 58},
	"9m9": {Min: 53, Max: 78},
	"9gm": {Min: 61, Max: 99},
	"7cl": {Min: 35, Max: 43},
	"7sp": {Min: 32, Max: 58},
	"7ma": {Min: 41, Max: 49},
	"7mt": {Min: 42, Max: 53},
	"7fl": {Min: 3, Max: 80},
	"7wh": {Min: 50, Max: 61},
	"7m7": {Min: 77, Max: 106},
	"7gm": {Min: 33, Max: 180},

	//Polearms
	"bar": {TwoMin: 1, TwoMax: 27},
	"vou": {TwoMin: 6, TwoMax: 21},
	"scy": {TwoMin: 8, TwoMax: 20},
	"pax": {TwoMin: 18, TwoMax: 39},
	"hal": {TwoMin: 21, TwoMax: 45},
	"wsc": {TwoMin: 15, TwoMax: 36},
	"9b7": {TwoMin: 6, TwoMax: 58},
	"9vo": {TwoMin: 14, TwoMax: 53},
	"9s8": {TwoMin: 18, TwoMax: 45},
	"9pa": {TwoMin: 34, TwoMax: 75},
	"9h9": {TwoMin: 13, TwoMax: 85},
	"9wc": {TwoMin: 30, TwoMax: 70},
	"7o7": {TwoMin: 28, TwoMax: 145},
	"7vo": {TwoMin: 17, TwoMax: 165},
	"7s8": {TwoMin: 12, TwoMax: 141},
	"7pa": {TwoMin: 33, TwoMax: 150},
	"7h7": {TwoMin: 46, TwoMax: 127},
	"7wc": {TwoMin: 40, TwoMax: 127},

	//Swords
	"ssd": {Min: 2, Max: 7},
	"scm": {Min: 2, Max: 6},
	"sbr": {Min: 3, Max: 8},
	"flc": {Min: 9, Max: 17},
	"crs": {Min: 5, Max: 15},
	"bsd": {Min: 7, Max: 14},
	"lsd": {Min: 3, Max: 19},
	"wsd": {Min: 8, Max: 20},
	"9ss": {Min: 8, Max: 22},
	"9sm": {Min: 8, Max: 21},
	"9sb": {Min: 10, Max: 24},
	"9fc": {Min: 16, Max: 35},
	"9cr": {Min: 13, Max: 35},
	"9bs": {Min: 16, Max: 34},
	"9ls": {Min: 10, Max: 42},
	"9wd": {Min: 18, Max: 43},
	"7ss": {Min: 31, Max: 59},
	"7sm": {Min: 26, Max: 46},
	"7sb": {Min: 33, Max: 45},
	"7fc": {Min: 28, Max: 68},
	"7cr": {Min: 31, Max: 35},
	"7bs": {Min: 37, Max: 53},
	"7ls": {Min: 5, Max: 77},
	"7wd": {Min: 40, Max: 50},

	// Two handed swords
	"2hs": {Min: 2, Max: 9, TwoMin: 5, TwoMax: 17},
	"clm": {Min: 5, Max: 12, TwoMin: 13, TwoMax: 30},
	"gis": {Min: 3, Max: 16, TwoMin: 9, TwoMax: 28},
	"bsw": {Min: 7, Max: 19, TwoMin: 20, TwoMax: 28},
	"flb": {Min: 9, Max: 15, TwoMin: 13, TwoMax: 26},
	"gsd": {Min: 12, Max: 20, TwoMin: 25, TwoMax: 42},
	"92h": {Min: 8, Max: 26, TwoMin: 18, TwoMax: 40},
	"9cm": {Min: 13, Max: 30, TwoMin: 26, TwoMax: 61},
	"9gs": {Min: 10, Max: 37, TwoMin: 19, TwoMax: 58},
	"9b9": {Min: 14, Max: 40, TwoMin: 39, TwoMax: 60},
	"9fb": {Min: 19, Max: 35, TwoMin: 29, TwoMax: 54},
	"9gd": {Min: 24, Max: 40, TwoMin: 47, TwoMax: 80},
	"7sh": {Min: 22, Max: 56, TwoMin: 50, TwoMax: 94},
	"7cm": {Min: 22, Max: 62, TwoMin: 67, TwoMax: 96},
	"7gs": {Min: 15, Max: 75, TwoMin: 55, TwoMax: 118},
	"7b7": {Min: 24, Max: 54, TwoMin: 71, TwoMax: 83},
	"7fb": {Min: 26, Max: 70, TwoMin: 61, TwoMax: 121},
	"7gd": {Min: 25, Max: 65, TwoMin: 58, TwoMax: 115},

	//Bows
	"sbw": {TwoMin: 1, TwoMax: 4},
	"hbw": {TwoMin: 2, TwoMax: 6},
	"lbw": {TwoMin: 3, TwoMax: 10},
	"cbw": {TwoMin: 4, TwoMax: 8},
	"sbb": {TwoMin: 3, TwoMax: 18},
	"swb": {TwoMin: 6, TwoMax: 14},
	"lwb": {TwoMin: 3, TwoMax: 23},
	"8sb": {TwoMin: 6, TwoMax: 19},
	"8hb": {TwoMin: 8, TwoMax: 22},
	"8lb": {TwoMin: 10, TwoMax: 29},
	"8cb": {TwoMin: 11, TwoMax: 26},
	"8s8": {TwoMin: 13, TwoMax: 30},
	"8l8": {TwoMin: 10, TwoMax: 42},
	"8sw": {TwoMin: 14, TwoMax: 35},
	"8lw": {TwoMin: 10, TwoMax: 50},
	"6sb": {TwoMin: 23, TwoMax: 50},
	"6hb": {TwoMin: 21, TwoMax: 41},
	"6lb": {TwoMin: 15, TwoMax: 59},
	"6cb": {TwoMin: 12, TwoMax: 52},
	"6s7": {TwoMin: 33, TwoMax: 40},
	"6l7": {TwoMin: 15, TwoMax: 63},
	"6sw": {TwoMin: 20, TwoMax: 53},
	"6lw": {TwoMin: 10, TwoMax: 68},

	//Crossbows
	"lxb": {TwoMin: 6, TwoMax: 9},
	"mxb": {TwoMin: 9, TwoMax: 16},
	"hxb": {TwoMin: 14, TwoMax: 26},
	"rxb": {TwoMin: 6, TwoMax: 12},
	"8lx": {TwoMin: 14, TwoMax: 27},
	"8mx": {TwoMin: 20, TwoMax: 42},
	"8hx": {TwoMin: 33, TwoMax: 55},
	"8rx": {TwoMin: 14, TwoMax: 32},
	"6lx": {TwoMin: 28, TwoMax: 73},
	"6mx": {TwoMin: 25, TwoMax: 87},
	"6hx": {TwoMin: 32, TwoMax: 91},
	"6rx": {TwoMin: 26, TwoMax: 40},

	//Javelins (throw damage, should update fields)
	"jav": {Min: 6, Max: 14},
	"pil": {Min: 7, Max: 20},
	"ssp": {Min: 10, Max: 22},
	"glv": {Min: 16, Max: 22},
	"tsp": {Min: 12, Max: 30},
	"9ja": {Min: 14, Max: 32},
	"9pi": {Min: 16, Max: 42},
	"9s9": {Min: 27, Max: 50},
	"9gl": {Min: 32, Max: 60},
	"9ts": {Min: 18, Max: 54},
	"7ja": {Min: 28, Max: 55},
	"7pi": {Min: 21, Max: 75},
	"7s7": {Min: 40, Max: 62},
	"7gl": {Min: 30, Max: 85},
	"7ts": {Min: 11, Max: 77},

	// Amazon-only Weapons (throw damage, should update fields)
	"am1": {Min: 7, Max: 12},
	"am2": {Min: 9, Max: 19},
	"am3": {Min: 18, Max: 24},
	"am4": {Min: 23, Max: 55},
	"am5": {Min: 6, Max: 22},
	"am6": {Min: 16, Max: 29},
	"am7": {Min: 19, Max: 41},
	"am8": {Min: 34, Max: 51},
	"am9": {Min: 42, Max: 101},
	"ama": {Min: 18, Max: 54},
	"amb": {Min: 20, Max: 47},
	"amc": {Min: 14, Max: 72},
	"amd": {Min: 65, Max: 95},
	"ame": {Min: 37, Max: 153},
	"amf": {Min: 35, Max: 66},

	//Scepters
	"scp": {Min: 6, Max: 11},
	"gsc": {Min: 8, Max: 18},
	"wsp": {Min: 10, Max: 17},
	"9sc": {Min: 13, Max: 24},
	"9qs": {Min: 14, Max: 36},
	"9ws": {Min: 16, Max: 38},
	"7sc": {Min: 40, Max: 52},
	"7qs": {Min: 45, Max: 54},
	"7ws": {Min: 37, Max: 43},

	//Spears
	"spr": {TwoMin: 3, TwoMax: 15},
	"tri": {TwoMin: 9, TwoMax: 15},
	"brn": {TwoMin: 7, TwoMax: 17},
	"spt": {TwoMin: 15, TwoMax: 23},
	"pik": {TwoMin: 14, TwoMax: 63},
	"9sr": {TwoMin: 10, TwoMax: 37},
	"9tr": {TwoMin: 19, TwoMax: 37},
	"9br": {TwoMin: 16, TwoMax: 40},
	"9st": {TwoMin: 29, TwoMax: 59},
	"9p9": {TwoMin: 27, TwoMax: 114},
	"7sr": {TwoMin: 35, TwoMax: 119},
	"7tr": {TwoMin: 29, TwoMax: 144},
	"7br": {TwoMin: 42, TwoMax: 92},
	"7st": {TwoMin: 18, TwoMax: 155},
	"7p7": {TwoMin: 33, TwoMax: 178},

	//Throwing
	"tkf": {Min: 4, Max: 9},
	"tax": {Min: 6, Max: 11},
	"bkf": {Min: 8, Max: 12},
	"bal": {Min: 12, Max: 15},
	"9tk": {Min: 11, Max: 24},
	"9ta": {Min: 14, Max: 27},
	"9bk": {Min: 18, Max: 33},
	"9b8": {Min: 24, Max: 34},
	"7tk": {Min: 23, Max: 54},
	"7bk": {Min: 23, Max: 39},
	"7ta": {Min: 15, Max: 66},
	"7b8": {Min: 7, Max: 60},

	//Daggers
	"dgr": {Min: 1, Max: 4},
	"dir": {Min: 3, Max: 9},
	"kri": {Min: 2, Max: 11},
	"bld": {Min: 4, Max: 15},
	"9dg": {Min: 6, Max: 18},
	"9di": {Min: 10, Max: 26},
	"9kr": {Min: 15, Max: 31},
	"9bl": {Min: 19, Max: 36},
	"7dg": {Min: 23, Max: 49},
	"7di": {Min: 37, Max: 53},
	"7kr": {Min: 15, Max: 57},
	"7bl": {Min: 32, Max: 47},

	//Staves
	"sst": {TwoMin: 1, TwoMax: 5},
	"lst": {TwoMin: 2, TwoMax: 8},
	"cst": {TwoMin: 4, TwoMax: 12},
	"bst": {TwoMin: 6, TwoMax: 13},
	"wst": {TwoMin: 12, TwoMax: 28},
	"8ss": {TwoMin: 6, TwoMax: 21},
	"8ls": {TwoMin: 8, TwoMax: 26},
	"8cs": {TwoMin: 11, TwoMax: 32},
	"8bs": {TwoMin: 14, TwoMax: 34},
	"8ws": {TwoMin: 24, TwoMax: 58},
	"6ss": {TwoMin: 69, TwoMax: 85},
	"6ls": {TwoMin: 75, TwoMax: 107},
	"6cs": {TwoMin: 80, TwoMax: 93},
	"6bs": {TwoMin: 65, TwoMax: 108},
	"6ws": {TwoMin: 83, TwoMax: 99},

	//Wands
	"wnd": {Min: 2, Max: 4},
	"ywn": {Min: 2, Max: 8},
	"bwn": {Min: 3, Max: 7},
	"gwn": {Min: 5, Max: 11},
	"9wn": {Min: 8, Max: 18},
	"9yw": {Min: 8, Max: 24},
	"9bw": {Min: 10, Max: 22},
	"9gw": {Min: 13, Max: 29},
	"7wn": {Min: 18, Max: 33},
	"7yw": {Min: 20, Max: 40},
	"7bw": {Min: 10, Max: 31},
	"7gw": {Min: 22, Max: 28},

	//Orbs
	"ob1": {Min: 2, Max: 5},
	"ob2": {Min: 3, Max: 8},
	"ob3": {Min: 4, Max: 10},
	"ob4": {Min: 5, Max: 12},
	"ob5": {Min: 8, Max: 18},
	"ob6": {Min: 8, Max: 21},
	"ob7": {Min: 10, Max: 26},
	"ob8": {Min: 11, Max: 29},
	"ob9": {Min: 13, Max: 32},
	"oba": {Min: 18, Max: 42},
	"obb": {Min: 21, Max: 46},
	"obc": {Min: 18, Max: 50},
	"obd": {Min: 23, Max: 55},
	"obe": {Min: 12, Max: 66},
	"obf": {Min: 30, Max: 53},

	// Assassin Claws
	"ktr": {Min: 4, Max: 7},
	"wrb": {Min: 5, Max: 9},
	"axf": {Min: 2, Max: 15},
	"ces": {Min: 7, Max: 15},
	"clw": {Min: 8, Max: 15},
	"btl": {Min: 10, Max: 14},
	"skr": {Min: 9, Max: 17},
	"9ar": {Min: 11, Max: 24},
	"9wb": {Min: 13, Max: 27},
	"9xf": {Min: 8, Max: 37},
	"9cs": {Min: 16, Max: 37},
	"9lw": {Min: 18, Max: 37},
	"9tw": {Min: 21, Max: 35},
	"9qr": {Min: 19, Max: 40},
	"7ar": {Min: 39, Max: 52},
	"7wb": {Min: 34, Max: 45},
	"7xf": {Min: 44, Max: 53},
	"7cs": {Min: 36, Max: 42},
	"7lw": {Min: 22, Max: 53},
	"7tw": {Min: 24, Max: 44},
	"7qr": {Min: 40, Max: 51},
}

// Item types, used to decide what attribute to give an item socketed with
// gems or runes mostly.
const (
	armor  = 0x01
	shield = 0x02
	weapon = 0x03
	other  = 0x04
)

// Item location IDs.
const (
	stored   = 0x00
	equipped = 0x01
	belt     = 0x02
	cursor   = 0x04
	socketed = 0x06
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

// Each set item has 5 bits of data containing the number of set lists follow
// the magical attributes list, this map tells us how many lists to read
// depending on the value given from the 5 bits. A number of 0-5 set lists.
var setListMap = map[uint64]uint64{
	0:  0,
	1:  1,
	2:  1,
	3:  2,
	4:  1,
	6:  2,
	7:  3,
	10: 2,
	12: 2,
	15: 4,
	31: 5,
}

// Certain set items (only Civerb's Ward in unmodded D2) have bonuses
// that require certain other set items in order to be activated
// (instead of the normal requirements of just 'wearing > x of any
// items in the set'); determined by add_func=1 in SetItems.txt
var setReqIDsMap = map[uint64][]uint64{
	// Civerb's Ward: [Civerb's Icon, Civerb's Cudgel]
	0: {1, 2},
}

// All item types that contain the quantity bits will exist in here,
// we'll use this when reading items to make sure we only read quantity bits
// when they exist, or we'll ruin the rest of the bit offsets for the item.
var quantityMap = map[string]bool{
	//Misc
	"tbk": true,
	"ibk": true,
	"key": true,
	"gps": true,
	"ops": true,
	"gpm": true,
	"opm": true,
	"gpl": true,
	"opl": true,
	"aqv": true,
	"cqv": true,

	//Throwables
	"tkf": true,
	"tax": true,
	"bkf": true,
	"bal": true,
	"9tk": true,
	"9ta": true,
	"9bk": true,
	"9b8": true,
	"7tk": true,
	"7ta": true,
	"7bk": true,
	"7b8": true,

	//Javelins
	"jav": true,
	"pil": true,
	"ssp": true,
	"glv": true,
	"tsp": true,
	"9ja": true,
	"9pi": true,
	"9s9": true,
	"9gl": true,
	"9ts": true,
	"7ja": true,
	"7pi": true,
	"7s7": true,
	"7gl": true,
	"7ts": true,
	"am5": true,
	"ama": true,
	"amf": true,
}

// Items that are tomes contain 5 extra bits, so we need to keep track of what
// items are tomes, and read the bits accordingly.
var tomeMap = map[string]bool{
	"tbk": true,
	"ibk": true,
}

var magicalProperties = map[uint64]magicalProperty{
	0:  {Bits: []uint{8}, Bias: 32, Name: "+{0} to Strength"},
	1:  {Bits: []uint{7}, Bias: 32, Name: "+{0} to Energy"},
	2:  {Bits: []uint{7}, Bias: 32, Name: "+{0} to Dexterity"},
	3:  {Bits: []uint{7}, Bias: 32, Name: "+{0} to Vitality"},
	7:  {Bits: []uint{9}, Bias: 32, Name: "+{0} to Life"},
	9:  {Bits: []uint{8}, Bias: 32, Name: "+{0} to Mana"},
	11: {Bits: []uint{8}, Bias: 32, Name: "+{0} to Maximum Stamina"},
	16: {Bits: []uint{9}, Name: "+{0}% Enhanced Defense"},

	// It's weird that there's two bit fields here, but it seems like
	// the first bit field enhanced min damage, and the second enchances
	// maxium damage.
	17: {Bits: []uint{9, 9}, Name: "+{0}% Enhanced Damage"},
	19: {Bits: []uint{10}, Name: "+{0} to Attack rating"},
	20: {Bits: []uint{6}, Name: "+{0}% Increased chance of blocking"},
	21: {Bits: []uint{6}, Name: "+{0} to Minimum 1-handed damage"},
	22: {Bits: []uint{7}, Name: "+{0} to Maximum 1-handed damage"},
	23: {Bits: []uint{6}, Name: "+{0} to Minimum 2-handed damage"},
	24: {Bits: []uint{7}, Name: "+{0} to Maximum 2-handed damage"},
	25: {Bits: []uint{8}, Name: "Unknown (Invisible)"}, // damagepercent
	26: {Bits: []uint{8}, Name: "Unknown (Invisible)"}, // manarecovery
	27: {Bits: []uint{8}, Name: "Regenerate Mana {0}%"},
	28: {Bits: []uint{8}, Name: "Heal Stamina {0}%"},
	31: {Bits: []uint{11}, Bias: 10, Name: "+{0} Defense"},
	32: {Bits: []uint{9}, Name: "+{0} vs. Missile"},
	33: {Bits: []uint{8}, Bias: 10, Name: "+{0} vs. Melee"},
	34: {Bits: []uint{6}, Name: "Damage Reduced by {0}"},
	35: {Bits: []uint{6}, Name: "Magic Damage Reduced by {0}"},
	36: {Bits: []uint{8}, Name: "Damage Reduced by {0}%"},
	37: {Bits: []uint{8}, Name: "Magic Resist +{0}%"},
	38: {Bits: []uint{8}, Name: "+{0}% to Maximum Magic Resist"},
	39: {Bits: []uint{8}, Bias: 50, Name: "Fire Resist +{0}%"},
	40: {Bits: []uint{5}, Name: "+{0}% to Maximum Fire Resist"},
	41: {Bits: []uint{8}, Bias: 50, Name: "Lightning Resist +{0}%"},
	42: {Bits: []uint{5}, Name: "+{0}% to Maximum Lightning Resist"},
	43: {Bits: []uint{8}, Bias: 50, Name: "Cold Resist +{0}%"},
	44: {Bits: []uint{5}, Name: "+{0}% to Maximum Cold Resist"},
	45: {Bits: []uint{8}, Bias: 50, Name: "Poison Resist +{0}%"},
	46: {Bits: []uint{5}, Name: "+{0}% to Maximum Poison Resist"},
	48: {Bits: []uint{8, 9}, Name: "Adds {0}-{1} Fire Damage"},
	49: {Bits: []uint{9}, Name: "+{0} to Maximum Fire Damage"},
	50: {Bits: []uint{6, 10}, Name: "Adds {0}-{1} Lightning Damage"},
	52: {Bits: []uint{8, 9}, Name: "Adds {0}-{1} Magic Damage"},
	54: {Bits: []uint{8, 9, 8}, Name: "Adds {0}-{1} Cold Damage"},
	57: {Bits: []uint{10, 10, 9}, Name: "Adds {0}-{1} Poison Damage over {2} Seconds"},
	60: {Bits: []uint{7}, Name: "{0}% Life Stolen Per Hit"},
	62: {Bits: []uint{7}, Name: "{0}% Mana Stolen Per Hit"},
	67: {Bits: []uint{7}, Bias: 30, Name: "Unknown (Invisible)"},  // velocitypercent
	68: {Bits: []uint{7}, Bias: 30, Name: "Unknown (Invisible)"},  // attackrate
	71: {Bits: []uint{8}, Bias: 100, Name: "Unknown (Invisible)"}, // value
	72: {Bits: []uint{9}, Name: "Unknown (Invisible)"},            // durability
	73: {Bits: []uint{8}, Name: "+{0} Maximum Durability"},
	74: {Bits: []uint{6}, Bias: 30, Name: "Replenish Life +{0}"},
	75: {Bits: []uint{7}, Bias: 20, Name: "Increase Maximum Durability {0}%"},
	76: {Bits: []uint{6}, Bias: 10, Name: "Increase Maximum Life {0}%"},
	77: {Bits: []uint{6}, Bias: 10, Name: "Increase Maximum Mana {0}%"},
	78: {Bits: []uint{7}, Name: "Attacker Takes Damage of {0}"},
	79: {Bits: []uint{9}, Bias: 100, Name: "{0}% Extra Gold from Monsters"},
	80: {Bits: []uint{8}, Bias: 100, Name: "{0}% Better Chance of Getting Magic Items"},
	81: {Bits: []uint{7}, Name: "Knockback"},
	82: {Bits: []uint{9}, Bias: 20, Name: "Unknown (Invisible)"}, // item_timeduration

	// First value is class, second is skill level, but they're printed in reverse
	// e.g. "+3 To Sorceress Skill Levels"
	83: {Bits: []uint{3, 3}, Name: "+{1} to {0} Skill Levels"},
	84: {Bits: []uint{3, 3}, Name: "+{1} to {0} Skill Levels"},

	// TODO: Check if experience gained really have a bias of 50.
	85: {Bits: []uint{9}, Bias: 50, Name: "{0}% To Experience Gained"},
	86: {Bits: []uint{7}, Name: "+{0} Life After Each Kill"},
	87: {Bits: []uint{7}, Name: "Reduces Prices {0}%"},
	88: {Bits: []uint{1}, Name: "Unknown (Invisible)"}, // item_doubleherbduration
	89: {Bits: []uint{4}, Bias: 4, Name: "+{0} to Light Radius"},
	// This property is not displayed on the item, but its effect is to alter
	// the color of the ambient light.
	90: {Bits: []uint{5}, Name: "Ambient light"},
	// After subtracting the bias, this is usually a negative number.
	91: {Bits: []uint{8}, Bias: 100, Name: "Requirements {0}%"},
	92: {Bits: []uint{7}, Name: "Level requirements +{0} (Invisible)"},
	93: {Bits: []uint{7}, Bias: 20, Name: "{0}% Increased Attack Speed"},
	94: {Bits: []uint{7}, Bias: 64, Name: "Unknown (Invisible)"}, // item_levelreqpct
	96: {Bits: []uint{7}, Bias: 20, Name: "{0}% Faster Run/Walk"},

	// Number of levels to a certain skill, e.g. +1 To Teleport.
	97: {Bits: []uint{9, 6}, Name: "+{1} To {0}"},

	// NVSTATE Charm attributes. ID 98 only occurs on charms of a special
	// type, called NV state charms, they're basically for visual effects.
	// They're imported charms and does not occur naturally in the game.
	98: {Bits: []uint{8, 1}, Name: "{1}+ to {0} (Visual effect only)"},

	99:  {Bits: []uint{7}, Bias: 20, Name: "{0}% Faster Hit Recovery"},
	102: {Bits: []uint{7}, Bias: 20, Name: "{0}% Faster Block Rate"},
	105: {Bits: []uint{7}, Bias: 20, Name: "{0}% Faster Cast Rate"},

	// These properties usually applied to class specific items,
	// first value selects the skill, the second determines how many
	// additional skill points are given.
	107: {Bits: []uint{9, 3}, Name: "+{1} To {0}"},
	108: {Bits: []uint{1}, Name: "Rest In Peace"},
	109: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
	181: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
	182: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
	183: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
	184: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
	185: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
	186: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
	187: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},

	110: {Bits: []uint{8}, Bias: 20, Name: "Poison Length Reduced by {0}%"},
	111: {Bits: []uint{9}, Bias: 20, Name: "Damage +{0}"},
	112: {Bits: []uint{7}, Name: "Hit Causes Monsters to Flee {0}%"},
	113: {Bits: []uint{7}, Name: "Hit Blinds Target +{0}"},
	114: {Bits: []uint{6}, Name: "{0}% Damage Taken Goes to Mana"},
	// The value of the data field is always 1.
	115: {Bits: []uint{1}, Name: "Ignore Target Defense"},
	116: {Bits: []uint{7}, Name: "{0}% Target Defense"},
	// The value of the data field is always 1.
	117: {Bits: []uint{7}, Name: "Prevent Monster Heal"},
	// The value of the data field is always 1.
	118: {Bits: []uint{1}, Name: "Half Freeze Duration"},
	119: {Bits: []uint{9}, Bias: 20, Name: "{0}% Bonus to Attack Rating"},
	120: {Bits: []uint{7}, Bias: 128, Name: "{0} to Monster Defense Per Hit"},
	121: {Bits: []uint{9}, Bias: 20, Name: "+{0}% Damage to Demons"},
	122: {Bits: []uint{9}, Bias: 20, Name: "+{0}% Damage to Undead"},
	123: {Bits: []uint{10}, Bias: 128, Name: "+{0} to Attack Rating against Demons"},
	124: {Bits: []uint{10}, Bias: 128, Name: "+{0} to Attack Rating against Undead"},
	125: {Bits: []uint{1}, Name: "Throwable"},
	// First value is class id, the next one is skill tree
	126: {Bits: []uint{3, 3}, Name: "+{0} to Fire Skills"},
	127: {Bits: []uint{3}, Name: "+{0} to All Skill Levels"},
	128: {Bits: []uint{5}, Name: "Attacker Takes Lightning Damage of {0}"},
	134: {Bits: []uint{5}, Name: "Freezes Target +{0}"},
	135: {Bits: []uint{7}, Name: "{0}% Chance of Open Wounds"},
	136: {Bits: []uint{7}, Name: "{0}% Chance of Crushing Blow"},
	137: {Bits: []uint{7}, Name: "+{0} Kick Damage"},
	138: {Bits: []uint{7}, Name: "+{0} to Mana After Each Kill"},
	139: {Bits: []uint{7}, Name: "+{0} Life after each Demon Kill"},
	// Unknown property, shows up on Swordback Hold Spiked Shield.
	140: {Bits: []uint{7}, Name: "Extra Blood (Invisible)"}, // item_extrablood
	141: {Bits: []uint{7}, Name: "{0}% Deadly Strike"},
	142: {Bits: []uint{7}, Name: "Fire Absorb {0}%"},
	143: {Bits: []uint{7}, Name: "+{0} Fire Absorb"},
	144: {Bits: []uint{7}, Name: "Lightning Absorb {0}%"},
	145: {Bits: []uint{7}, Name: "+{0} Lightning Absorb"},
	146: {Bits: []uint{7}, Name: "Magic Absorb {0}%"},
	147: {Bits: []uint{7}, Name: "+{0} Magic Absorb"},
	148: {Bits: []uint{7}, Name: "Cold Absorb {0}%"},
	149: {Bits: []uint{7}, Name: "+{0} Cold Absorb"},
	150: {Bits: []uint{7}, Name: "Slows Target by {0}%"},
	151: {Bits: []uint{9, 5}, Name: "Level +{1} {0} When Equipped"},
	152: {Bits: []uint{1}, Name: "Indestructible"},
	153: {Bits: []uint{1}, Name: "Cannot Be Frozen"},
	154: {Bits: []uint{7}, Bias: 20, Name: "{0}% Slower Stamina Drain"},
	155: {Bits: []uint{10, 7}, Name: "{0}% Chance to Reanimate Target"},
	156: {Bits: []uint{7}, Name: "Piercing Attack"},
	157: {Bits: []uint{7}, Name: "Fires Magic Arrows"},
	158: {Bits: []uint{7}, Name: "Fires Explosive Arrows or Bolts"},
	159: {Bits: []uint{6}, Name: "+{0} to Minimum Throw Damage"},
	160: {Bits: []uint{7}, Name: "+{0} to Maximum Throw Damage"},
	179: {Bits: []uint{3}, Name: "+{0} to Druid Skill Levels"},
	180: {Bits: []uint{3}, Name: "+{0} to Assassin Skill Levels"},

	// Ok so get this, this is quite complicated, the id 188 is for items with
	// + x to a certain skill tree, e.g. 1 + to defensive auras (paladin).
	//
	// So here's how it works, the field is 19 bits long, here's the bits for
	// the defensive auras skiller.
	//
	// 001             0000000000             011              010
	//  ^                  ^                   ^                ^
	// levels        unknown padding        class id       skill tree offset
	//
	// So in the above example, the first 3 bits 001 are the + levels (1), we'll
	// ignore the padding, the second interesting set of 3 bits is the class id.
	// Refer to the class.go for class ids, but paladin is 011 (3), and the
	// last 3 bits 010 (2) is the offset from the start of the class skill tree.
	// Refer to skills.go to find the different tree offsets. Paladin offset is
	// 9. So remember the last 3 bits 010 (2), that means the skill tree is
	// 9 + 2 = 11, aka the defensive auras tree.
	//
	// When reading the values, remember the bits are read from the right,
	// so the values will be [2 3 1], offset 2, class id 3, 1 + to skills.
	188: {Bits: []uint{3, 13, 3}, Name: "+{2} to {0} Skills ({1} only)"},

	189: {Bits: []uint{10, 9}, Name: "+{0} to {1} Skills (char_class Only)"},
	190: {Bits: []uint{10, 9}, Name: "+{0} to {1} Skills (char_class Only)"},
	191: {Bits: []uint{10, 9}, Name: "+{0} to {1} Skills (char_class Only)"},
	192: {Bits: []uint{10, 9}, Name: "+{0} to {1} Skills (char_class Only)"},
	193: {Bits: []uint{10, 9}, Name: "+{0} to {1} Skills (char_class Only)"},

	194: {Bits: []uint{4}, Name: "Adds {0} extra sockets to the item"},

	// Order is spell id, level, % chance.
	195: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} When you die"},
	196: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} When you die"},
	197: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} When you die"},

	// Order is spell id, level, % chance.
	198: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} On Striking"},
	199: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} On Striking"},
	200: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} On Striking"},

	// Order is spell id, level, % chance.
	201: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} When Struck"},
	202: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} When Struck"},
	203: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} When Struck"},

	// First value selects the spell id, second value is level, third is remaining charges
	// and the last is the total number of charges.
	204: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
	205: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
	206: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
	207: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
	208: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
	209: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
	210: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
	211: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
	212: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
	213: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},

	// All values based on character level are stored in eights, so take
	// the number divide by 8 and multiply by the character level and round down.
	// Or, just do (value * 0.125)% per level.
	214: {Bits: []uint{6}, Name: "+{0} to Defense (Based on Character Level)"},
	215: {Bits: []uint{6}, Name: "{0}% Enhanced Defense (Based on Character Level)"},
	216: {Bits: []uint{6}, Name: "+{0} to Life (Based on Character Level)"},
	217: {Bits: []uint{6}, Name: "+{0} to Mana (Based on Character Level)"},
	218: {Bits: []uint{6}, Name: "+{0} to Maximum Damage (Based on Character Level)"},
	219: {Bits: []uint{6}, Name: "{0}% Enhanced Maximum Damage (Based on Character Level)"},
	220: {Bits: []uint{6}, Name: "+{0} to Strength (Based on Character Level)"},
	221: {Bits: []uint{6}, Name: "+{0} to Dexterity (Based on Character Level)"},
	222: {Bits: []uint{6}, Name: "+{0} to Energy (Based on Character Level)"},
	223: {Bits: []uint{6}, Name: "+{0} to Vitality (Based on Character Level)"},
	224: {Bits: []uint{6}, Name: "+{0} to Attack Rating (Based on Character Level)"},
	225: {Bits: []uint{6}, Name: "{0}% Bonus to Attack Rating (Based on Character Level)"},
	226: {Bits: []uint{6}, Name: "+{0} Cold Damage (Based on Character Level)"},
	227: {Bits: []uint{6}, Name: "+{0} Fire Damage (Based on Character Level)"},
	228: {Bits: []uint{6}, Name: "+{0} Lightning Damage (Based on Character Level)"},
	229: {Bits: []uint{6}, Name: "+{0} Poison Damage (Based on Character Level)"},
	230: {Bits: []uint{6}, Name: "Cold Resist +{0}% (Based on Character Level)"},
	231: {Bits: []uint{6}, Name: "Fire Resist +{0}% (Based on Character Level)"},
	232: {Bits: []uint{6}, Name: "Lightning Resist +{0}% (Based on Character Level)"},
	233: {Bits: []uint{6}, Name: "Poison Resist +{0}% (Based on Character Level)"},
	234: {Bits: []uint{6}, Name: "+{0} Cold Absorb (Based on Character Level)"},
	235: {Bits: []uint{6}, Name: "+{0} Fire Absorb (Based on Character Level)"},
	236: {Bits: []uint{6}, Name: "+{0} Lightning Absorb (Based on Character Level)"},
	237: {Bits: []uint{6}, Name: "{0} Poison Absorb (Based on Character Level)"},
	238: {Bits: []uint{5}, Name: "Attacker Takes Damage of {0} (Based on Character Level)"},
	239: {Bits: []uint{6}, Name: "{0}% Extra Gold from Monsters (Based on Character Level)"},
	240: {Bits: []uint{6}, Name: "{0}% Better Chance of Getting Magic Items (Based on Character Level)"},
	241: {Bits: []uint{6}, Name: "Heal Stamina Plus {0}% (Based on Character Level)"},
	242: {Bits: []uint{6}, Name: "+{0} Maxmium Stamina (Based on Character Level)"},
	243: {Bits: []uint{6}, Name: "{0}% Damage to Demons (Based on Character Level)"},
	244: {Bits: []uint{6}, Name: "{0}% Damage to Undead (Based on Character Level)"},
	245: {Bits: []uint{6}, Name: "+{0} to Attack Rating against Demons (Based on Character Level)"},
	246: {Bits: []uint{6}, Name: "+{0} to Attack Rating against Undead (Based on Character Level)"},
	247: {Bits: []uint{6}, Name: "{0}% Chance of Crushing Blow (Based on Character Level)"},
	248: {Bits: []uint{6}, Name: "{0}% Chance of Open Wounds (Based on Character Level)"},
	249: {Bits: []uint{6}, Name: "+{0} Kick Damage (Based on Character Level)"},
	250: {Bits: []uint{6}, Name: "{0}% to Deadly Strike (Based on Character Level)"},

	// The value of the data field is not actually a time period, but a frequency in terms
	// of the number of times durability is repaired over a period of 100 seconds.
	// For example, if the value is 5, then this property repairs 1 durability in 100 / 5 = 20 seconds.
	252: {Bits: []uint{6}, Name: "Repairs 1 Durability in {0} Seconds"},

	// As in the previous property, the value of the data field is a frequency in terms of the number
	// replenished over a period of 100 seconds. For example if the value is 4, then this property
	// replenishes 1 item in 100 / 4 = 25 seconds.
	253: {Bits: []uint{6}, Name: "Replenishes Quantity"},

	// Number of additional items beyond the base limit, for example if the base
	// is 50 and additional is 30, then the total is 50 + 30.
	254: {Bits: []uint{8}, Name: "Increased Stack Size"},

	// IDs 268 - 303 are some weird values that were never used in the actual game.
	// These values change depending on the time of day in the game.
	// The format of the bit fields are the same in all cases, the first 2 bits
	// specifies the the of time when the value is at its maximum.
	//
	// The second and third are respectively the minimum and maximum values of the property.
	// The maximum value at the time specified and the minimum at the opposite.

	305: {Bits: []uint{8}, Bias: 50, Name: "{0} Pierce Cold"},
	306: {Bits: []uint{8}, Bias: 50, Name: "{0} Pierce Fire"},
	307: {Bits: []uint{8}, Bias: 50, Name: "{0} Pierce Lightning"},
	308: {Bits: []uint{8}, Bias: 50, Name: "{0} Pierce Poision"},

	324: {Bits: []uint{6}, Name: "Unknown (Invisible)"}, // item_extra_charges
	329: {Bits: []uint{9}, Bias: 50, Name: "{0}% To Fire Skill Damage"},
	330: {Bits: []uint{9}, Bias: 50, Name: "{0}% To Lightning Skill Damage"},
	331: {Bits: []uint{9}, Bias: 50, Name: "{0}% To Cold Skill Damage"},
	332: {Bits: []uint{9}, Bias: 50, Name: "{0}% To Poison Skill Damage"},
	333: {Bits: []uint{8}, Name: "-{0}% To Enemy Fire Resistance"},
	334: {Bits: []uint{8}, Name: "-{0}% To Enemy Lightning Resistance"},
	335: {Bits: []uint{8}, Name: "-{0}% To Enemy Cold Resistance"},
	336: {Bits: []uint{8}, Name: "-{0}% To Enemy Poison Resistance"},
	356: {Bits: []uint{2}, Name: "Quest Item Difficulty +{0} (Invisible)"},
}

// Shield codes
var shieldCodes = map[string]string{
	"uow": "Aegis",
	"pa4": "Aerin Shield",
	"pa7": "Akaran Rondache",
	"pa6": "Akaran Targe",
	"pad": "Kurast Shield",
	"xts": "Ancient Shield",
	"xpk": "Barbed Shield",
	"upk": "Blade Barrier",
	"nef": "Bloodlord Skull",
	"bsh": "Bone Shield",
	"buc": "Buckler",
	"ne9": "Cantor Trophy",
	"pa5": "Crown Shield",
	"xuc": "Defender",
	"ne5": "Demon Head",
	"xit": "Dragon Shield",
	"ne7": "Fetish Trophy",
	"ne4": "Gargoyle Head",
	"gts": "Gothic Shield",
	"xsh": "Grim Shield",
	"pa9": "Gilded Shield",
	"nea": "Heirophant Trophy",
	"uuc": "Heater",
	"neg": "Hellspawn Skull",
	"pa3": "Heraldic Shield",
	"kit": "Kite Shield",
	"lrg": "Large Shield",
	"uml": "Luna",
	"neb": "Minion Skull",
	"uit": "Monarch",
	"ne6": "Mummified Trophy",
	"ned": "Overseer Skull",
	"xow": "Pavise",
	"ne1": "Preserved Head",
	"pa8": "Protector Shield",
	"pa2": "Rondache",
	"xml": "Round Shield",
	"paa": "Royal Shield",
	"pac": "Sacred Rondache",
	"pab": "Sacred Targe",
	"xrg": "Scutum",
	"ne8": "Sexton Trophy",
	"sml": "Small Shield",
	"spk": "Spiked Shield",
	"nee": "Succubus Skull",
	"pa1": "Targe",
	"tow": "Tower Shield",
	"ush": "Troll Nest",
	"ne3": "Unraveller Head",
	"paf": "Vortex Shield",
	"uts": "Ward",
	"pae": "Zakarum Shield",
	"ne2": "Zombie Head",
}

// Armor codes.
var armorCodes = map[string]string{
	"dr6": "Alpha Helm",
	"aar": "Ancient Armor",
	"dr3": "Antlers",
	"utp": "Archon Plate",
	"ulm": "Armet",
	"ba4": "Assault Helmet",
	"ba5": "Avenger Guard",
	"upl": "Balrog Skin",
	"xhl": "Basinet",
	"ztb": "Battle Belt",
	"xtb": "Battle Boots",
	"xtg": "Battle Gauntlets",
	"mbl": "Belt",
	"drb": "Blood Spirit",
	"bhm": "Bone Helm",
	"uh9": "Bone Visage",
	"uhn": "Boneweave",
	"umb": "Boneweave Boots",
	"mgl": "Bracers",
	"ulg": "Bramble Mitts",
	"brs": "Breast Plate",
	"cap": "Cap",
	"bab": "Carnage Helm",
	"xlm": "Casque",
	"mbt": "Chain Boots",
	"chn": "Chain Mail",
	"xul": "Chaos Armor",
	"ci0": "Circlet",
	"uhc": "Colossus Girdle",
	"bae": "Conqueror Crown",
	"urn": "Corona",
	"ci1": "Coronet",
	"crn": "Crown",
	"utg": "Crusader Gauntlets",
	"xrs": "Cuirass",
	"xsk": "Death Mask",
	"usk": "Demonhead",
	"xla": "Demonhide Armor",
	"xlb": "Demonhide Boots",
	"xlg": "Demonhide Gloves",
	"zlb": "Demonhide Sash",
	"bad": "Destroyer Helm",
	"ci3": "Diadem",
	"ung": "Diamond Mail",
	"drf": "Dream Spirit",
	"uui": "Dusk Shroud",
	"drd": "Earth Spirit",
	"xth": "Embossed Plate",
	"dr4": "Falcon Mask",
	"ba2": "Fanged Helm",
	"fld": "Field Plate",
	"fhl": "Full Helm",
	"plt": "Plate Mail",
	"ful": "Full Plate Mail",
	"bac": "Fury Visor",
	"hgl": "Gauntlets",
	"xui": "Ghost Armor",
	"uhl": "Giant Conch",
	"hbl": "Girdle",
	"lgl": "Gloves",
	"gth": "Gothic Plate",
	"xrn": "Grand Crown",
	"urs": "Great Hauberk",
	"ghm": "Great Helm",
	"dr7": "Griffon Headress",
	"xh9": "Grim Helm",
	"baf": "Guardian Crown",
	"hla": "Hard Leather Armor",
	"dr2": "Hawk Helm",
	"tbl": "Heavy Belt",
	"vbt": "Heavy Boots",
	"xmg": "Heavy Bracers",
	"vgl": "Heavy Gloves",
	"ult": "Hellforged Plate",
	"hlm": "Helm",
	"ba3": "Horned Helm",
	"dr8": "Hunter's Guise",
	"ukp": "Hydraskull",
	"urg": "Hyperion",
	"ba1": "Jawbone Cap",
	"ba6": "Jawbone Visor",
	"uld": "Kraken Shell",
	"uth": "Lacquered Plate",
	"lea": "Leather Armor",
	"lbt": "Leather Boots",
	"vbl": "Light Belt",
	"tgl": "Light Gauntlets",
	"ltp": "Light Plate",
	"tbt": "Light Plated Boots",
	"xng": "Linked Mail",
	"ba7": "Lion Helm",
	"ucl": "Loricated Mail",
	"xtp": "Mage Plate",
	"msk": "Mask",
	"xhn": "Mesh Armor",
	"zmb": "Mesh Belt",
	"xmb": "Mesh Boots",
	"utb": "Mirrored Boots",
	"umc": "Mithril Coil",
	"uhb": "Myrmidon Greaves",
	"uhg": "Ogre Gauntlets",
	"xar": "Ornate Plate",
	"hbt": "Plate Boots",
	"qui": "Quilted Armor",
	"ba8": "Rage Mask",
	"rng": "Ring Mail",
	"xpl": "Russet Armor",
	"uar": "Sacred Armor",
	"dr9": "Sacred Feathers",
	"xkp": "Sallet",
	"lbl": "Sash",
	"ba9": "Savage Helmet",
	"scl": "Scale Mail",
	"ula": "Scarab Husk",
	"uvb": "Scarabshell Boots",
	"xea": "Serpentskin Armor",
	"uul": "Shadow Plate",
	"uap": "Shako",
	"zvb": "Sharkskin Belt",
	"xvb": "Sharkskin Boots",
	"xvg": "Sharkskin Gloves",
	"xld": "Sharktooth Armor",
	"skp": "Skull Cap",
	"dre": "Sky Spirit",
	"baa": "Slayer Guard",
	"ulc": "Spiderweb Sash",
	"uhm": "Spired Helm",
	"dr5": "Spirit Mask",
	"spl": "Splint Mail",
	"stu": "Studded Leather",
	"drc": "Sun Spirit",
	"xlt": "Templar Coat",
	"ci2": "Tiara",
	"xcl": "Tigulated Mail",
	"dra": "Totemic Mask",
	"xtu": "Trellised Armor",
	"utc": "Troll Belt",
	"umg": "Vambraces",
	"uvg": "Vampirebone Gloves",
	"uvc": "Vampirefang Belt",
	"zhb": "War Belt",
	"xhb": "War Boots",
	"xhg": "War Gauntlets",
	"xap": "War Hat",
	"xhm": "Winged Helm",
	"utu": "Wire Fleece",
	"dr1": "Wolf Head",
	"uea": "Wyrmhide",
	"ulb": "Wyrmhide Boots",
}

// Weapon Codes.
var weaponCodes = map[string]string{
	"9gi": "Ancient Axe",
	"9wd": "Ancient Sword",
	"8lx": "Arbalest",
	"6ws": "Archon Staff",
	"am6": "Ashwood Bow",
	"7sm": "Ataghan",
	"axe": "Axe",
	"bal": "Balanced Axe",
	"bkf": "Balanced Knife",
	"8hx": "Balista",
	"7gs": "Balrog Blade",
	"7s7": "Balrog Spear",
	"9sp": "Barbed Club",
	"bar": "Bardiche",
	"bsw": "Bastard Sword",
	"btx": "Battle Axe",
	"7cs": "Battle Cestus",
	"9tk": "Battle Dart",
	"9wh": "Battle Hammer",
	"9s8": "Battle Scythe",
	"bst": "Battle Staff",
	"9bs": "Battle Sword",
	"9ba": "Bearded Axe",
	"9h9": "Bec De Corbin",
	"7wa": "Berserker Axe",
	"9vo": "Bill",
	"bld": "Blade",
	"6hb": "Blade Bow",
	"btl": "Blade Talons",
	"7dg": "Bone Knife",
	"bwn": "Bone Wand",
	"brn": "Brandistock",
	"bax": "Broad Axe",
	"bsd": "Broad Sword",
	"9wn": "Burnt Wand",
	"7ws": "Caduceus",
	"8lb": "Cedar Bow",
	"8cs": "Cedar Staff",
	"am7": "Ceremonial Bow",
	"ama": "Ceremonial Javelin",
	"am9": "Ceremonial Pike",
	"am8": "Ceremonial Spear",
	"ces": "Cestus",
	"7ga": "Champion Axe",
	"7b7": "Champion Sword",
	"gpm": "Choking Gas Potion",
	"8rx": "Chu-Ko-Nu",
	"9kr": "Cinquedeas",
	"ob4": "Clasped Orb",
	"clw": "Claws",
	"clm": "Claymore",
	"9ax": "Cleaver",
	"ob8": "Cloudy Sphere",
	"clb": "Club",
	"7fb": "Colossal Sword",
	"7gd": "Colossal Blade",
	"6hx": "Colossus Crossbow",
	"7vo": "Colossus Voulge",
	"cbw": "Composite Bow",
	"7bs": "Conquest Sword",
	"mxb": "Crossbow",
	"9mp": "Crowbill",
	"6l7": "Crusader Bow",
	"7pa": "Cryptic Axe",
	"7ls": "Cryptic Sword",
	"crs": "Crystal Sword",
	"ob7": "Crystalline Globe",
	"9cl": "Cudgel",
	"9sm": "Cutlass",
	"9cm": "Dacian Falx",
	"dgr": "Dagger",
	"7bt": "Decapitator",
	"d33": "Decoy Dagger",
	"6rx": "Demon Crossbow",
	"obd": "Demon Heart",
	"7mt": "Devil Star",
	"6s7": "Diamond Bow",
	"9cr": "Dimensional Blade",
	"obf": "Dimensional Shard",
	"dir": "Dirk",
	"9ws": "Divine Scepter",
	"2ax": "Double Axe",
	"8cb": "Double Bow",
	"ob5": "Dragon Stone",
	"ob1": "Eable Orb",
	"8sb": "Edge Bow",
	"6cs": "Elder Staff",
	"obc": "Eldritch Orb",
	"7sb": "Elegant Blade",
	"92h": "Espadon",
	"72a": "Ettin Axe",
	"9gd": "Executioner Sword",
	"opm": "Exploding Potion",
	"7ss": "Falcata",
	"flc": "Falchion",
	"7kr": "Fanged Knife",
	"9xf": "Fascia",
	"7la": "Feral Axe",
	"7lw": "Feral Claws",
	"fla": "Flail",
	"flb": "Flamberge",
	"9ma": "Flanged Mace",
	"7ta": "Flying Axe",
	"7tk": "Flying Knife",
	"9ta": "Francisca",
	"opl": "Fulmating Potion",
	"9tr": "Fuscina",
	"7gl": "Ghost Glaive",
	"7st": "Ghost Spear",
	"7yw": "Ghost Wand",
	"gix": "Giant Axe",
	"gis": "Giant Sword",
	"7wc": "Giant Thresher",
	"g33": "Gidbinn",
	"9ss": "Gladius",
	"glv": "Glaive",
	"7gi": "Glorious Axe",
	"ob6": "Glowing Orb",
	"cst": "Gnarled Staff",
	"6mx": "Gorgon Crossbow",
	"9ga": "Gothic Axe",
	"8lw": "Gothic Bow",
	"8bs": "Gothic Staff",
	"9b9": "Gothic Sword",
	"amc": "Grand Matron Bow",
	"gsc": "Grand Scepter",
	"9gw": "Grave Wand",
	"gax": "Great Axe",
	"6cb": "Great Bow",
	"gma": "Great Maul",
	"9pi": "Great Pilum",
	"7h7": "Great Poleaxe",
	"gsd": "Great Sword",
	"9lw": "Greater Claws",
	"9tw": "Greater Talons",
	"9wc": "Grim Scythe",
	"gwn": "Grim Wand",
	"hal": "Halberd",
	"hax": "Hand Axe",
	"9cs": "Hand Scythe",
	"9ts": "Harpoon",
	"9ha": "Hatchet",
	"axf": "HatchetHands",
	"obb": "Heavenly Stone",
	"hxb": "Heavy Crossbow",
	"hfh": "Hellforge Hammer",
	"7cm": "Highland Blade",
	"9qs": "Holy Water Sprinkler",
	"hdm": "Horadric Malus",
	"hst": "Horadric Staff",
	"hbw": "Hunters Bow",
	"9b8": "Hurlbat",
	"6lw": "Hydra Bow",
	"7fc": "Hydra Edge",
	"7ja": "Hyperion Javelin",
	"7sr": "Hyperion Spear",
	"9mt": "Jagged Star",
	"jav": "Javelin",
	"8ss": "Jo Staff",
	"ktr": "Katar",
	"qf1": "Khalim Flail",
	"9fl": "Knout",
	"kri": "Kriss",
	"9p9": "Lance",
	"lax": "Large Axe",
	"7bl": "Legend Spike",
	"72h": "Legend Sword",
	"7wh": "Legendary Mallet",
	"7bw": "Lich Wand",
	"lxb": "Light Crossbow",
	"9b7": "Lochaber Axe",
	"lbb": "Long Battle Bow",
	"lbw": "Long Bow",
	"8l8": "Long Siege Bow",
	"lst": "Long Staff",
	"lsd": "Long Sword",
	"lwb": "Long War Bow",
	"mac": "Mace",
	"am5": "Maiden Javelin",
	"am4": "Maiden Pike",
	"am3": "Maiden Spear",
	"7br": "Mancatcher",
	"9gm": "Martel de Fer",
	"amb": "Matriarchal Bow",
	"ame": "Matriarchal Pike",
	"amd": "Matriarchal Spear",
	"amf": "Matriarchal Javelin",
	"mau": "Maul",
	"7sc": "Mighty Scepter",
	"9la": "Military Axe",
	"mpi": "Military Pick",
	"7di": "Mithral Point",
	"mst": "Morning Star",
	"7wd": "Mythical Sword",
	"9wa": "Naga",
	"7o7": "Ogre Axe",
	"7m7": "Ogre Maul",
	"ops": "Oil Potion",
	"9pa": "Partizan",
	"6lx": "Pellet Bow",
	"9yw": "Petrified Wand",
	"7cr": "Phase Blade",
	"pik": "Pike",
	"pil": "Pilum",
	"9dg": "Poignard",
	"pax": "Poleaxe",
	"7wn": "Polished Wand",
	"8ls": "Quarter Staff",
	"9ar": "Quhab",
	"gps": "Rancid Gas Potion",
	"8hb": "Razor Bow",
	"am2": "Reflex Bow",
	"7ma": "Reinforced Mace",
	"rxb": "Repeating Crossbow",
	"9di": "Rondel",
	"8sw": "Rune Bow",
	"9sc": "Rune Scepter",
	"8ws": "Rune Staff",
	"9ls": "Rune Sword",
	"7tw": "Runic Talons",
	"sbr": "Sabre",
	"ob2": "Sacred Globe",
	"scp": "Scepter",
	"scm": "Scimitar",
	"skr": "Scissors Katar",
	"9qr": "Scissors Quhab",
	"7qr": "Scissors Suwayyah",
	"7fl": "Scourge",
	"scy": "Scythe",
	"7qs": "Seraph Rod",
	"6lb": "Shadow Bow",
	"9sb": "Shamshir",
	"6bs": "Shillelah",
	"sbb": "Short Battle Bow",
	"sbw": "Short Bow",
	"8s8": "Short Siege Bow",
	"ssp": "Short Spear",
	"sst": "Short Staff",
	"ssd": "Short Sword",
	"swb": "Short War Bow",
	"8mx": "Siege Crossbow",
	"7ba": "Silver Edged Axe",
	"9s9": "Simbilan",
	"7ax": "Small Crescent",
	"ob3": "Smoked Sphere",
	"ob9": "Sparkling Ball",
	"spr": "Spear",
	"spt": "Spetum",
	"9gl": "Spiculum",
	"6sb": "Spider Bow",
	"spc": "Spiked Club",
	"msf": "Staff Of The Kings",
	"am1": "Stag Bow",
	"6ls": "Stalagmite",
	"9bl": "Stilleto",
	"gpl": "Strangling Gas Potion",
	"7tr": "Stygian Pike",
	"7pi": "Stygian Pilum",
	"qf2": "Super Khalim Flail",
	"7ar": "Suwayyah",
	"oba": "Swirling Crystal",
	"9bt": "Tabar",
	"7s8": "Thresher",
	"tax": "Throwing Axe",
	"tkf": "Throwing Knife",
	"tsp": "Throwing Spear",
	"7gm": "Thunder Maul",
	"7ha": "Tomahawk",
	"9bw": "Tomb Wand",
	"tri": "Trident",
	"7cl": "Truncheon",
	"9fc": "Tulwar",
	"9gs": "Tusk Sword",
	"92a": "Twin Axe",
	"2hs": "Two Handed Sword",
	"7sp": "Tyrant Club",
	"7gw": "Unearthed Wand",
	"obe": "Vortex Orb",
	"vou": "Voulge",
	"6ss": "Walking Stick",
	"wnd": "Wand",
	"wax": "War Axe",
	"9m9": "War Club",
	"9bk": "War Dart",
	"7xf": "War Fist",
	"9br": "War Fork",
	"whm": "War Hammer",
	"9ja": "War Javelin",
	"7p7": "War Pike",
	"wsp": "War Scepter",
	"wsc": "War Scythe",
	"9sr": "War Spear",
	"7mp": "War Spike",
	"wst": "War Staff",
	"wsd": "War Sword",
	"6sw": "Ward Bow",
	"7b8": "Winged Axe",
	"7ts": "Winged Harpoon",
	"7bk": "Winged Knife",
	"leg": "Wirts Leg",
	"wrb": "Wrist Blade",
	"9wb": "Wrist Spike",
	"7wb": "Wrist Sword",
	"9st": "Yari",
	"ywn": "Yew Wand",
	"9fb": "Zweihander",
}

// Misc codes, like jewelry, gems, potions and runes.
var miscCodes = map[string]string{
	"gsv": "Amethyst",
	"r11": "Amn Rune",
	"amu": "Amulet",
	"yps": "Antidote Potion",
	"aqv": "Arrows",
	"bey": "Baal's Eye",
	"bks": "Bark Scroll",
	"r30": "Ber Rune",
	"cqv": "Bolts",
	"ass": "Book of Skill",
	"bet": "Burning Essence of Terror",
	"r32": "Cham Rune",
	"ceh": "Charged Essence of Hatred",
	"cm3": "Grand Charm",
	"cm2": "Large Charm",
	"cm1": "Small Charm",
	"gcv": "Chipped Amethyst",
	"gcw": "Chipped Diamond",
	"gcg": "Chipped Emerald",
	"gcr": "Chipped Ruby",
	"gcb": "Chipped Sapphire",
	"skc": "Chipped Skull",
	"gcy": "Chipped Topaz",
	"bkd": "Deciphered Bark Scroll",
	"dhn": "Diablo's Horn",
	"gsw": "Diamond",
	"r14": "Dol Rune",
	"r01": "El Rune",
	"r02": "Eld Rune",
	"elx": "Elixir",
	"gsg": "Emerald",
	"r05": "Eth Rune",
	"r19": "Fal Rune",
	"fed": "Festering Essence of Destruction",
	"gfv": "Flawed Amethyst",
	"gfw": "Flawed Diamond",
	"gfg": "Flawed Emerald",
	"gfr": "Flawed Ruby",
	"gfb": "Flawed Sapphire",
	"skf": "Flawed Skull",
	"gfy": "Flawed Topaz",
	"gzv": "Flawless Amethyst",
	"glw": "Flawless Diamond",
	"glg": "Flawless Emerald",
	"glr": "Flawless Ruby",
	"glb": "Flawless Sapphire",
	"skl": "Flawless Skull",
	"gly": "Flawless Topaz",
	"hpf": "Full Healing Potion",
	"mpf": "Full Mana Potion",
	"rvl": "Full Rejuvenation Potion",
	"gld": "Gold",
	"g34": "Gold Bird",
	"hp5": "Greater Healing Potion",
	"mp5": "Greater Mana Potion",
	"r25": "Gul Rune",
	"hp3": "Healing Potion",
	"hpo": "Healing Potion",
	"r15": "Hel Rune",
	"hrb": "Herb",
	"box": "Horadric Cube",
	"ibk": "Identify Book",
	"isc": "Identify Scroll",
	"r16": "Io Rune",
	"r24": "Ist Rune",
	"r06": "Ith Rune",
	"j34": "Jade Figurine",
	"r31": "Jah Rune",
	"jew": "Jewel",
	"pk3": "Key of Destruction",
	"pk2": "Key of Hate",
	"pk1": "Key of Terror",
	"qbr": "Khalim's Brain",
	"qey": "Khalim's Eye",
	"qhr": "Khalim's Heart",
	"r18": "Ko Rune",
	"bbb": "Lam Esens Tome",
	"bpl": "Large Blue Potion",
	"rpl": "Large Red Potion",
	"r20": "Lem Rune",
	"hp1": "Lesser Healing Potion",
	"mp1": "Lesser Mana Potion",
	"hp2": "Light Healing Potion",
	"mp2": "Light Mana Potion",
	"r28": "Lo Rune",
	"r17": "Lum Rune",
	"ice": "Maguffin",
	"r23": "Mal Rune",
	"mp3": "Mana Potion",
	"mpo": "Mana Potion",
	"mbr": "Mephisto's Brain",
	"luv": "Mephisto Key",
	"mss": "Mephisto Soul Stone",
	"r04": "Nef Rune",
	"r27": "Ohm Rune",
	"r09": "Ort Rune",
	"gpv": "Perfect Amethyst",
	"gpw": "Perfect Diamond",
	"gpg": "Perfect Emerald",
	"gpr": "Perfect Ruby",
	"gpb": "Perfect Sapphire",
	"skz": "Perfect Skull",
	"gpy": "Perfect Topaz",
	"ear": "Player Ear",
	"r21": "Pul Rune",
	"r08": "Ral Rune",
	"rvs": "Rejuvenation Potion",
	"rin": "Ring",
	"gsr": "Ruby",
	"gsb": "Sapphire",
	"0sc": "Scroll",
	"tr1": "Scroll of Horadric",
	"tr2": "Scroll of Malah",
	"xyz": "Potion of Life",
	"r13": "Shael Rune",
	"key": "Skeleton Key",
	"sku": "Skull",
	"bps": "Small Blue Potion",
	"rps": "Small Red Potion",
	"r12": "Sol Rune",
	"vps": "Stamina Potion",
	"std": "Standard of Heroes",
	"hp4": "Strong Healing Potion",
	"mp4": "Strong Mana Potion",
	"r29": "Sur Rune",
	"r07": "Tal Rune",
	"wms": "Thawing Potion",
	"r10": "Thul Rune",
	"r03": "Tir Rune",
	"toa": "Token of Absolution",
	"gsy": "Topaz",
	"tch": "Torch",
	"tbk": "Town Portal Book",
	"tsc": "Town Portal Scroll",
	"tes": "Twisted Essence of Suffering",
	"r22": "Um Rune",
	"r26": "Vex Rune",
	"vip": "Viper Amulet",
	"r33": "Zod Rune",
}

var rareNames = map[uint64]string{
	1:   "Bite",
	2:   "Scratch",
	3:   "Scalpel",
	4:   "Fang",
	5:   "Gutter",
	6:   "Thirst",
	7:   "Razor",
	8:   "Scythe",
	9:   "Edge",
	10:  "Saw",
	11:  "Splitter",
	12:  "Cleaver",
	13:  "Sever",
	14:  "Sunder",
	15:  "Rend",
	16:  "Mangler",
	17:  "Slayer",
	18:  "Reaver",
	19:  "Spawn",
	20:  "Gnash",
	21:  "Star",
	22:  "Blow",
	23:  "Smasher",
	24:  "Bane",
	25:  "Crusher",
	26:  "Breaker",
	27:  "Grinder",
	28:  "Crack",
	29:  "Mallet",
	30:  "Knell",
	31:  "Lance",
	32:  "Spike",
	33:  "Impaler",
	34:  "Skewer",
	35:  "Prod",
	36:  "Scourge",
	37:  "Wand",
	38:  "Wrack",
	39:  "Barb",
	40:  "Needle",
	41:  "Dart",
	42:  "Bolt",
	43:  "Quarrel",
	44:  "Fletch",
	45:  "Flight",
	46:  "Nock",
	47:  "Horn",
	48:  "Stinger",
	49:  "Quill",
	50:  "Goad",
	51:  "Branch",
	52:  "Spire",
	53:  "Song",
	54:  "Call",
	55:  "Cry",
	56:  "Spell",
	57:  "Chant",
	58:  "Weaver",
	59:  "Gnarl",
	60:  "Visage",
	61:  "Crest",
	62:  "Circlet",
	63:  "Veil",
	64:  "Hood",
	65:  "Mask",
	66:  "Brow",
	67:  "Casque",
	68:  "Visor",
	69:  "Cowl",
	70:  "Hide",
	71:  "Pelt",
	72:  "Carapace",
	73:  "Coat",
	74:  "Wrap",
	75:  "Suit",
	76:  "Cloak",
	77:  "Shroud",
	78:  "Jack",
	79:  "Mantle",
	80:  "Guard",
	81:  "Badge",
	82:  "Rock",
	83:  "Aegis",
	84:  "Ward",
	85:  "Tower",
	86:  "Shield",
	87:  "Wing",
	88:  "Mark",
	89:  "Emblem",
	90:  "Hand",
	91:  "Fist",
	92:  "Claw",
	93:  "Clutches",
	94:  "Grip",
	95:  "Grasp",
	96:  "Hold",
	97:  "Torch",
	98:  "Finger",
	99:  "Knuckle",
	100: "Shank",
	101: "Spur",
	102: "Tread",
	103: "Stalker",
	104: "Greave",
	105: "Blazer",
	106: "Nails",
	107: "Trample",
	108: "Brogues",
	109: "Track",
	110: "Slippers",
	111: "Clasp",
	112: "Buckle",
	113: "Harness",
	114: "Lock",
	115: "Fringe",
	116: "Winding",
	117: "Chain",
	118: "Strap",
	119: "Lash",
	120: "Cord",
	121: "Knot",
	122: "Circle",
	123: "Loop",
	124: "Eye",
	125: "Turn",
	126: "Spiral",
	127: "Coil",
	128: "Gyre",
	129: "Band",
	130: "Whorl",
	131: "Talisman",
	132: "Heart",
	133: "Noose",
	134: "Necklace",
	135: "Collar",
	136: "Beads",
	137: "Torc",
	138: "Gorget",
	139: "Scarab",
	140: "Wood",
	141: "Brand",
	142: "Bludgeon",
	143: "Cudgel",
	144: "Loom",
	145: "Harp",
	146: "Master",
	147: "Barl",
	148: "Hew",
	149: "Crook",
	150: "Mar",
	151: "Shell",
	152: "Stake",
	153: "Picket",
	154: "Pale",
	155: "Flange",
	156: "Beast",
	157: "Eagle",
	158: "Raven",
	159: "Viper",
	160: "Ghoul",
	161: "Skull",
	162: "Blood",
	163: "Dread",
	164: "Doom",
	165: "Grim",
	166: "Bone",
	167: "Death",
	168: "Shadow",
	169: "Storm",
	170: "Rune",
	171: "Plague",
	172: "Stone",
	173: "Wraith",
	174: "Spirit",
	175: "Storm",
	176: "Demon",
	177: "Cruel",
	178: "Empyrion",
	179: "Bramble",
	180: "Pain",
	181: "Loath",
	182: "Glyph",
	183: "Imp",
	184: "Fiendra",
	185: "Hailstone",
	186: "Gale",
	187: "Dire",
	188: "Soul",
	189: "Brimstone",
	190: "Corpse",
	191: "Carrion",
	192: "Armageddon",
	193: "Havoc",
	194: "Bitter",
	195: "Entropy",
	196: "Chaos",
	197: "Order",
	198: "Rule",
	199: "Warp",
	200: "Rift",
	201: "Corruption",
}

var setNames = map[uint64]string{
	0:   "Civerb's Ward",
	1:   "Civerb's Icon",
	2:   "Civerb's Cudgel",
	3:   "Hsaru's Iron Heel",
	4:   "Hsaru's Iron Fist",
	5:   "Hsaru's Iron Stay",
	6:   "Cleglaw's Tooth",
	7:   "Cleglaw's Claw",
	8:   "Cleglaw's Pincers",
	9:   "Iratha's Collar",
	10:  "Iratha's Cuff",
	11:  "Iratha's Coil",
	12:  "Iratha's Cord",
	13:  "Isenhart's Lightbrand",
	14:  "Isenhart's Parry",
	15:  "Isenhart's Case",
	16:  "Isenhart's Horns",
	17:  "Vidala's Barb",
	18:  "Vidala's Fetlock",
	19:  "Vidala's Ambush",
	20:  "Vidala's Snare",
	21:  "Milabrega's Orb",
	22:  "Milabrega's Rod",
	23:  "Milabrega's Diadem",
	24:  "Mialbrega's Robe",
	25:  "Cathan's Rule",
	26:  "Cathan's Mesh",
	27:  "Cathan's Visage",
	28:  "Cathan's Sigil",
	29:  "Cathan's Seal",
	30:  "Tancred's Crowbill",
	31:  "Tancred's Spine",
	32:  "Tancred's Hobnails",
	33:  "Tancred's Weird",
	34:  "Tancred's Skull",
	35:  "Sigon's Gage",
	36:  "Sigon's Visor",
	37:  "Sigon's Shelter",
	38:  "Sigon's Sabot",
	39:  "Sigon's Wrap",
	40:  "Sigon's Guard",
	41:  "Infernal Cranium",
	42:  "Infernal Torch",
	43:  "Infernal Sign",
	44:  "Berserker's Headgear",
	45:  "Berserker's Hauberk",
	46:  "Berserker's Hatchet",
	47:  "Death's Hand",
	48:  "Death's Guard",
	49:  "Death's Touch",
	50:  "Angelic Sickle",
	51:  "Angelic Mantle",
	52:  "Angelic Halo",
	53:  "Angelic Wings",
	54:  "Arctic Horn",
	55:  "Arctic Furs",
	56:  "Arctic Binding",
	57:  "Arctic Mitts",
	58:  "Arcanna's Sign",
	59:  "Arcanna's Deathwand",
	60:  "Arcanna's Head",
	61:  "Arcanna's Flesh",
	62:  "Natalya's Totem",
	63:  "Natalya's Mark",
	64:  "Natalya's Shadow",
	65:  "Natalya's Soul",
	66:  "Aldur's Stony Gaze",
	67:  "Aldur's Deception",
	68:  "Aldur's Rhythm",
	69:  "Aldur's Advance",
	70:  "Immortal King's Will",
	71:  "Immortal King's Soul Cage",
	72:  "Immortal King's Detail",
	73:  "Immortal King's Forge",
	74:  "Immortal King's Pillar",
	75:  "Immortal King's Stone Crusher",
	76:  "Tal Rasha's Fine-Spun Cloth",
	77:  "Tal Rasha's Adjudication",
	78:  "Tal Rasha's Lidless Eye",
	79:  "Tal Rasha's Guardianship",
	80:  "Tal Rasha's Horadric Crest",
	81:  "Griswold's Valor",
	82:  "Griswold's Heart",
	83:  "Griswold's Redemption",
	84:  "Griswold's Honor",
	85:  "Trang-Oul's Guise",
	86:  "Trang-Oul's Scales",
	87:  "Trang-Oul's Wing",
	88:  "Trang-Oul's Claws",
	89:  "Trang-Oul's Girth",
	90:  "M'avina's True Sight",
	91:  "M'avina's Embrace",
	92:  "M'avina's Icy Clutch",
	93:  "M'avina's Tenet",
	94:  "M'avina's Caster",
	95:  "Telling of Beads",
	96:  "Laying of Hands",
	97:  "Rite of Passage",
	98:  "Dark Adherent",
	99:  "Credendum",
	100: "Dangoon's Teaching",
	101: "Taebaek's Glory",
	102: "Haemosu's Adament",
	103: "Ondal's Almighty",
	104: "Guillaume's Face",
	105: "Wilhelm's Pride",
	106: "Magnus' Skin",
	107: "Wihtstan's Guard",
	108: "Hwanin's Splendor",
	109: "Hwanin's Refuge",
	110: "Hwanin's Blessing",
	111: "Hwanin's Justice",
	112: "Sazabi's Cobalt Redeemer",
	113: "Sazabi's Ghost Liberator",
	114: "Sazabi's Mental Sheath",
	115: "Bul-Katho's Sacred Charge",
	116: "Bul-Katho's Tribal Guardian",
	117: "Cow King's Horns",
	118: "Cow King's Hide",
	119: "Cow King's Hooves",
	120: "Naj's Puzzler",
	121: "Naj's Light Plate",
	122: "Naj's Circlet",
	123: "Sander's Paragon",
	124: "Sander's Riprap",
	125: "Sander's Taboo",
	126: "Sander's Superstition",
}

var runewordNames = map[uint64]string{
	27:   "Ancient's Pledge",
	30:   "Beast",
	32:   "Black",
	34:   "Bone",
	35:   "Bramble",
	36:   "Brand",
	37:   "Breath of the Dying",
	39:   "Call to Arms",
	40:   "Chains of Honor",
	42:   "Chaos",
	43:   "Crescent Moon",
	46:   "Death",
	51:   "Destruction",
	52:   "Doom",
	53:   "Dragon",
	55:   "Dream",
	56:   "Duress",
	57:   "Edge",
	59:   "Enigma",
	60:   "Enlightenment",
	62:   "Eternity",
	63:   "Exile",
	64:   "Faith",
	65:   "Famine",
	67:   "Fortitude",
	70:   "Fury",
	71:   "Gloom",
	73:   "Grief",
	74:   "Hand of Justice",
	75:   "Harmory",
	77:   "Heart of the Oak",
	80:   "Holy Thunder",
	81:   "Honor",
	85:   "Ice",
	86:   "Infinity",
	88:   "Insight",
	91:   "King's Grace",
	92:   "Kingslayer",
	95:   "Last Wish",
	97:   "Lawbringer",
	98:   "Leaf",
	100:  "Lionheart",
	101:  "Lore",
	106:  "Malice",
	107:  "Melody",
	108:  "Memory",
	112:  "Myth",
	113:  "Nadir",
	116:  "Oath",
	117:  "Obedience",
	120:  "Passion",
	123:  "Peace",
	124:  "Winter",
	128:  "Phoenix",
	131:  "Plague",
	134:  "Pride",
	135:  "Principle",
	137:  "Prudence",
	141:  "Radiance",
	142:  "Rain",
	145:  "Rhyme",
	146:  "Rift",
	147:  "Sanctuary",
	151:  "Silence",
	153:  "Smoke",
	155:  "Spirit",
	156:  "Splendor",
	158:  "Stealth",
	159:  "Steel",
	162:  "Stone",
	164:  "Strength",
	173:  "Treachery",
	179:  "Venom",
	185:  "Wealth",
	187:  "White",
	188:  "Wind",
	193:  "Wrath",
	195:  "Zephyr",
	2718: "Delirium",
}

var uniqueNames = map[uint64]string{
	0:   "The Gnasher",
	1:   "Deathspade",
	2:   "Bladebone",
	3:   "Skull splitter",
	4:   "Rakescar",
	5:   "Axe of Fechmar",
	6:   "Goreshovel",
	7:   "The Chiefthan",
	8:   "Brainhew",
	9:   "Humongous",
	10:  "Torch of Iros",
	11:  "Maelstorm",
	12:  "Gravenspine",
	13:  "Umes Lament",
	14:  "Felloak",
	15:  "Knell Striker",
	16:  "Rusthandle",
	17:  "Stormeye",
	18:  "Stoutnail",
	19:  "Crushflange",
	20:  "Bloodrise",
	21:  "The Generals Tan Do Li Ga",
	22:  "Ironstone",
	23:  "Bonesnap",
	24:  "Steeldriver",
	25:  "Rixot's Keen",
	26:  "Blood Crescent",
	27:  "Skewer of Krintiz",
	28:  "Gleamscythe",
	29:  "Azurewrath",
	30:  "Griswold's Edge",
	31:  "Hellplague",
	32:  "Culwens Point",
	33:  "Shadowfang",
	34:  "Soulflay",
	35:  "Kinemils Awl",
	36:  "Blacktongue",
	37:  "Ripsaw",
	38:  "The Patriarch",
	39:  "Gull",
	40:  "The Diggler",
	41:  "The Jade Tan Do",
	42:  "Spectral Shard",
	43:  "The Dragon Chang",
	44:  "Razortine",
	45:  "Bloodthief",
	46:  "Lance of Yaggai",
	47:  "The Tannr Gorerod",
	48:  "Dimoaks Hew",
	49:  "Steelgoad",
	50:  "Soul Harvest",
	51:  "The Battlebranch",
	52:  "Woestave",
	53:  "The Grim Reaper",
	54:  "Bane Ash",
	55:  "Serpent Lord",
	56:  "Spire of Lazarus",
	57:  "The Salamander",
	58:  "The Iron Jang Bong",
	59:  "Pluckeye",
	60:  "Witherstring",
	61:  "Raven Claw",
	62:  "Rogue's Bow",
	63:  "Stormstrike",
	64:  "Wizendraw",
	65:  "Hellclap",
	66:  "Blastbark",
	67:  "Leadcrow",
	68:  "Ichorsting",
	69:  "Hellcast",
	70:  "Doomslinger",
	71:  "Biggin's Bonnet",
	72:  "Tarnhelm",
	73:  "Coif of Glory",
	74:  "Duskdeep",
	75:  "Wormskull",
	76:  "Howltusk",
	77:  "Undead Crown",
	78:  "The Face of Horror",
	79:  "Greyform",
	80:  "Blinkbat's Form",
	81:  "The Centurion",
	82:  "Twitchthroe",
	83:  "Darkglow",
	84:  "Hawkmail",
	85:  "Sparking Mail",
	86:  "Venom Ward",
	87:  "Iceblink",
	88:  "Boneflesh",
	89:  "Rockfleece",
	90:  "Rattlecage",
	91:  "Goldskin",
	92:  "Victors Silk",
	93:  "Heavenly Garb",
	94:  "Pelta Lunata",
	95:  "Umbral Disk",
	96:  "Stormguild",
	97:  "Wall of the Eyeless",
	98:  "Swordback Hold",
	99:  "Steelclash",
	100: "Bverrit Keep",
	101: "The Ward",
	102: "The Hand of Broc",
	103: "Bloodfist",
	104: "Chance Guards",
	105: "Magefist",
	106: "Frostburn",
	107: "Hotspur",
	108: "Gorefoot",
	109: "Treads of Cthon",
	110: "Goblin Toe",
	111: "Tearhaunch",
	112: "Lenymo",
	113: "Snakecord",
	114: "Nightsmoke",
	115: "Goldwrap",
	116: "Bladebuckle",
	117: "Nokozan Relic",
	118: "The Eye of Etlich",
	119: "The Mahim-Oak Curio",
	120: "Nagelring",
	121: "Manald Heal",
	122: "The Stone of Jordan",
	123: "Amulet of the Viper",
	124: "Staff of Kings",
	125: "Horadric Staff",
	126: "Hell Forge Hammer",
	127: "Khalim's Flail",
	128: "Super Khalim's Flail",
	129: "Coldkill",
	130: "Butcher's Pupil",
	131: "Islestrike",
	132: "Pompe's Wrath",
	133: "Guardian Naga",
	134: "Warlord's Trust",
	135: "Spellsteel",
	136: "Stormrider",
	137: "Boneslayer Blade",
	138: "The Minataur",
	139: "Suicide Branch",
	140: "Carin Shard",
	141: "Arm of King Leoric",
	142: "Blackhand Key",
	143: "Dark Clan Crusher",
	144: "Zakarum's Hand",
	145: "The Fetid Sprinkler",
	146: "Hand of Blessed Light",
	147: "Fleshrender",
	148: "Sureshrill Frost",
	149: "Moonfall",
	150: "Baezil's Vortex",
	151: "Earthshaker",
	152: "Bloodtree Stump",
	153: "The Gavel of Pain",
	154: "Bloodletter",
	155: "Coldsteel Eye",
	156: "Hexfire",
	157: "Blade of Ali Baba",
	158: "Ginther's Rift",
	159: "Headstriker",
	160: "Plague Bearer",
	161: "The Atlantian",
	162: "Crainte Vomir",
	163: "Bing Sz Wang",
	164: "The Vile Husk",
	165: "Cloudcrack",
	166: "Todesfaelle Flamme",
	167: "Swordguard",
	168: "Spineripper",
	169: "Heart Carver",
	170: "Blackbog's Sharp",
	171: "Stormspike",
	172: "The Impaler",
	173: "Kelpie Snare",
	174: "Soulfeast Tine",
	175: "Hone Sundan",
	176: "Spire of Honor",
	177: "The Meat Scraper",
	178: "Blackleach Blade",
	179: "Athena's Wrath",
	180: "Pierre Tombale Couant",
	181: "Husoldal Evo",
	182: "Grim's Burning Dead",
	183: "Razorswitch",
	184: "Ribcracker",
	185: "Chromatic Ire",
	186: "Warpspear",
	187: "Skullcollector",
	188: "Skystrike",
	189: "Riphook",
	190: "Kuko Shakaku",
	191: "Endlesshail",
	192: "Whichwild String",
	193: "Cliffkiller",
	194: "Magewrath",
	195: "Godstrike Arch",
	196: "Langer Briser",
	197: "Pus Spiter",
	198: "Buriza-Do Kyanon",
	199: "Demon Machine",
	200: "Armor (Unknown)",
	201: "Peasent Crown",
	202: "Rockstopper",
	203: "Stealskull",
	204: "Darksight Helm",
	205: "Valkyrie Wing",
	206: "Crown of Thieves",
	207: "Blckhorn's Face",
	208: "Vampire Gaze",
	209: "The Spirit Shroud",
	210: "Skin of the Vipermagi",
	211: "Skin of the Flayed One",
	212: "Ironpelt",
	213: "Spiritforge",
	214: "Crow Caw",
	215: "Shaftstop",
	216: "Duriel's Shell",
	217: "Skullder's Ire",
	218: "Guardian Angel",
	219: "Toothrow",
	220: "Atma's Wail",
	221: "Black Hades",
	222: "Corpsemourn",
	223: "Que-Hegan's Wisdom",
	224: "Visceratuant",
	225: "Mosers Blessed Circle",
	226: "Stormchaser",
	227: "Tiamat's Rebuke",
	228: "Gerke's Sanctuary",
	229: "Radimant's Sphere",
	230: "Lidless Wall",
	231: "Lance Guard",
	232: "Venom Grip",
	233: "Gravepalm",
	234: "Ghoulhide",
	235: "Lavagout",
	236: "Hellmouth",
	237: "Infernostride",
	238: "Waterwalk",
	239: "Silkweave",
	240: "Wartraveler",
	241: "Gorerider",
	242: "String of Ears",
	243: "Razortail",
	244: "Gloomstrap",
	245: "Snowclash",
	246: "Thundergod's Vigor",
	247: "Elite unique",
	248: "Harlequin Crest",
	249: "Veil of Steel",
	250: "The Gladiator's Bane",
	251: "Arkaine's Valor",
	252: "Blackoak Shield",
	253: "Stormshield",
	254: "Hellslayer",
	255: "Messerschmidt's Reaver",
	256: "Baranar's Star",
	257: "Schaefer's Hammer",
	258: "The Cranium Basher",
	259: "Lightsabre",
	260: "Doombringer",
	261: "The Grandfather",
	262: "Wizardspike",
	263: "Constricting Ring",
	264: "Stormspire",
	265: "Eaglehorn",
	266: "Windforce",
	267: "Ring",
	268: "Bul Katho's Wedding Band",
	269: "The Cat's Eye",
	270: "The Rising Sun",
	271: "Crescent Moon",
	272: "Mara's Kaleidoscope",
	273: "Atma's Scarab",
	274: "Dwarf Star",
	275: "Raven Frost",
	276: "Highlord's Wrath",
	277: "Saracen's Chance",
	278: "Class specific",
	279: "Arreat's Face",
	280: "Homunculus",
	281: "Titan's Revenge",
	282: "Lycander's Aim",
	283: "Lycander's Flank",
	284: "The Oculus",
	285: "Herald of Zakarum",
	286: "Bartuc's Cut-Throat",
	287: "Jalal's Mane",
	288: "The Scalper",
	289: "Bloodmoon",
	290: "Djinnslayer",
	291: "Deathbit",
	292: "Warshrike",
	293: "Gutsiphon",
	294: "Razoredge",
	295: "Gore Ripper",
	296: "Demon Limb",
	297: "Steel Shade",
	298: "Tomb Reaver",
	299: "Death's Web",
	300: "Nature's Peace",
	301: "Azurewrath",
	302: "Seraph's Hymn",
	303: "Zakarum's Salvation",
	304: "Fleshripper",
	305: "Odium",
	306: "Horizon's Tornado",
	307: "Stone Crusher",
	308: "Jade Talon",
	309: "Shadow Dancer",
	310: "Cerebus' Bite",
	311: "Tyrael's Might",
	312: "Soul Drainer",
	313: "Rune Master",
	314: "Death Cleaver",
	315: "Executioner's Justice",
	316: "Stoneraven",
	317: "Leviathan",
	318: "Larzuk's Champion",
	319: "Wisp Projector",
	320: "Gargoyle's Bite",
	321: "Lacerator",
	322: "Mang Song's Lesson",
	323: "Viperfork",
	324: "Ethereal Edge",
	325: "Demonhorn's Edge",
	326: "The Reaper's Toll",
	327: "Spiritkeeper",
	328: "Hellrack",
	329: "Alma Negra",
	330: "Darkforge Spawn",
	331: "Widowmaker",
	332: "Bloodraven's Charge",
	333: "Ghostflame",
	334: "Shadowkiller",
	335: "Gimmershred",
	336: "Griffon's Eye",
	337: "Windhammer",
	338: "Thunderstroke",
	339: "Giant Maimer",
	340: "Demon's Arch",
	341: "Boneflame",
	342: "Steelpillar",
	343: "Nightwing's Veil",
	344: "Crown of Ages",
	345: "Andariel's Visage",
	346: "Darkfear",
	347: "Dragonscale",
	348: "Steel Carapice",
	349: "Medusa's Gaze",
	350: "Ravenlore",
	351: "Boneshade",
	352: "Nethercrow",
	353: "Flamebellow",
	354: "Fathom",
	355: "Wolfhowl",
	356: "Spirit Ward",
	357: "Kira's Guardian",
	358: "Ormus Robes",
	359: "Gheed's Fortune",
	360: "Stormlash",
	361: "Halaberd's Reign",
	362: "Warriv's Warder",
	363: "Spike Thorn",
	364: "Dracul's Grasp",
	365: "Frostwind",
	366: "Templar's Might",
	367: "Eschuta's Temper",
	368: "Firelizard's Talons",
	369: "Sandstorm Trek",
	370: "Marrowwalk",
	371: "Heaven's Light",
	372: "Merman's Speed",
	373: "Arachnid Mesh",
	374: "Nosferatu's Coil",
	375: "Metalgrid",
	376: "Verdugo's Hearty Cord",
	377: "Sigurd's Staunch",
	378: "Carrion Wind",
	379: "Giantskull",
	380: "Ironward",
	381: "Annihilus",
	382: "Arioc's Needle",
	383: "Cranebeak",
	384: "Nord's Tenderizer",
	385: "Earthshifter",
	386: "Wraithflight",
	387: "Bonehew",
	388: "Ondal's Wisdom",
	389: "The Reedeemer",
	390: "Headhunter's Glory",
	391: "Steelrend",
	392: "Rainbow Facet",
	393: "Rainbow Facet",
	394: "Rainbow Facet",
	395: "Rainbow Facet",
	396: "Rainbow Facet",
	397: "Rainbow Facet",
	398: "Rainbow Facet",
	399: "Rainbow Facet",
	400: "Hellfire Torch",
}

var magicalPrefixes = map[uint64]string{
	2:   "Sturdy",
	3:   "Strong",
	4:   "Glorious",
	5:   "Blessed",
	6:   "Saintly",
	7:   "Holy",
	8:   "Devious",
	9:   "Fortified",
	13:  "Jagged",
	14:  "Deadly",
	15:  "Vicious",
	16:  "Brutal",
	17:  "Massive",
	18:  "Savage",
	19:  "Merciless",
	20:  "Vulpine",
	25:  "Tireless",
	26:  "Rugged",
	27:  "Bronze",
	28:  "Iron",
	29:  "Steel",
	30:  "Silver",
	32:  "Gold",
	33:  "Platinum",
	34:  "Meteoric",
	35:  "Sharp",
	36:  "Fine",
	37:  "Warrior's",
	38:  "Soldier's",
	39:  "Knight's",
	40:  "Lord's",
	41:  "King's",
	42:  "Howling",
	43:  "Fortuitous",
	49:  "Glimmering",
	50:  "Glowing",
	53:  "Lizard's",
	55:  "Snake's",
	56:  "Serpent's",
	57:  "Serpent's",
	58:  "Drake's",
	59:  "Dragon's",
	60:  "Dragon's",
	61:  "Wyrm's",
	64:  "Prismatic",
	65:  "Prismatic",
	66:  "Azure",
	67:  "Lapis",
	68:  "Lapis",
	69:  "Cobalt",
	70:  "Cobalt",
	72:  "Sapphire",
	75:  "Crimson",
	76:  "Burgundy",
	77:  "Burgundy",
	78:  "Garnet",
	79:  "Garnet",
	81:  "Ruby",
	84:  "Ocher",
	85:  "Tangerine",
	86:  "Tangerine",
	87:  "Coral",
	88:  "Coral",
	90:  "Amber",
	93:  "Beryl",
	94:  "Jade",
	95:  "Jade",
	96:  "Viridian",
	97:  "Viridian",
	99:  "Emerald",
	101: "Fletcher's",
	102: "Archer's",
	103: "Archer's",
	104: "Monk's",
	105: "Priest's",
	106: "Priest's",
	107: "Summoner's",
	108: "Necromancer's",
	109: "Necromancer's",
	110: "Angel's",
	111: "Arch-Angel's",
	112: "Arch-Angel's",
	113: "Slayer's",
	114: "Berserker's",
	115: "Berserker's",
	118: "Triumphant",
	119: "Stout",
	120: "Stout",
	121: "Stout",
	122: "Burly",
	123: "Burly",
	124: "Burly",
	125: "Stalwart",
	126: "Stalwart",
	127: "Stalwart",
	128: "Stout",
	129: "Stout",
	130: "Stout",
	131: "Burly",
	132: "Burly",
	133: "Stalwart",
	134: "Stalwart",
	135: "Stout",
	136: "Stout",
	137: "Burly",
	138: "Stalwart",
	139: "Blanched",
	140: "Eburin",
	141: "Bone",
	142: "Ivory",
	143: "Sturdy",
	144: "Sturdy",
	145: "Strong",
	146: "Glorious",
	147: "Blessed",
	148: "Saintly",
	149: "Holy",
	150: "Godly",
	151: "Devious",
	152: "Blank",
	153: "Null",
	154: "Antimagic",
	155: "Red",
	156: "Red",
	157: "Sanguinary",
	158: "Sanguinary",
	159: "Bloody",
	160: "Red",
	161: "Sanguinary",
	162: "Bloody",
	163: "Red",
	164: "Sanguinary",
	165: "Bloody",
	166: "Scarlet",
	167: "Crimson",
	168: "Jagged",
	169: "Jagged",
	170: "Jagged",
	171: "Forked",
	172: "Forked",
	173: "Serrated",
	174: "Serrated",
	175: "Jagged",
	176: "Jagged",
	177: "Forked",
	178: "Forked",
	179: "Serrated",
	180: "Jagged",
	181: "Forked",
	182: "Serrated",
	183: "Carbuncle",
	184: "Carmine",
	185: "Vermillion",
	186: "Jagged",
	187: "Deadly",
	188: "Vicious",
	189: "Brutal",
	190: "Massive",
	191: "Savage",
	192: "Merciless",
	193: "Ferocious",
	194: "Cruel",
	195: "Cinnabar",
	196: "Rusty",
	197: "Realgar",
	198: "Ruby",
	199: "Vulpine",
	200: "Dun",
	201: "Tireless",
	202: "Tireless",
	203: "Brown",
	204: "Rugged",
	205: "Rugged",
	206: "Rugged",
	207: "Rugged",
	208: "Rugged",
	209: "Rugged",
	210: "Rugged",
	211: "Rugged",
	212: "Rugged",
	213: "Rugged",
	214: "Rugged",
	215: "Vigorous",
	216: "Chestnut",
	217: "Maroon",
	218: "Bronze",
	219: "Bronze",
	220: "Bronze",
	221: "Iron",
	222: "Iron",
	223: "Iron",
	224: "Steel",
	225: "Steel",
	226: "Steel",
	227: "Bronze",
	228: "Bronze",
	229: "Bronze",
	230: "Iron",
	231: "Iron",
	232: "Steel",
	233: "Steel",
	234: "Bronze",
	235: "Bronze",
	236: "Iron",
	237: "Steel",
	238: "Bronze",
	239: "Iron",
	240: "Steel",
	241: "Silver",
	242: "Gold",
	243: "Platinum",
	244: "Meteoric",
	245: "Strange",
	246: "Weird",
	247: "Nickel",
	248: "Tin",
	249: "Silver",
	250: "Argent",
	251: "Fine",
	252: "Fine",
	253: "Sharp",
	254: "Fine",
	255: "Sharp",
	256: "Fine",
	257: "Sharp",
	258: "Fine",
	259: "Warrior's",
	260: "Soldier's",
	261: "Knight's",
	262: "Lord's",
	263: "King's",
	264: "Master's",
	265: "Grandmaster's",
	266: "Glimmering",
	267: "Glowing",
	268: "Bright",
	269: "Screaming",
	270: "Howling",
	271: "Wailing",
	272: "Screaming",
	273: "Howling",
	274: "Wailing",
	275: "Lucky",
	276: "Lucky",
	277: "Lucky",
	278: "Lucky",
	279: "Lucky",
	280: "Lucky",
	281: "Felicitous",
	282: "Fortuitous",
	283: "Emerald",
	284: "Lizard's",
	285: "Lizard's",
	286: "Lizard's",
	287: "Snake's",
	288: "Snake's",
	289: "Snake's",
	290: "Serpent's",
	291: "Serpent's",
	292: "Serpent's",
	293: "Lizard's",
	294: "Lizard's",
	295: "Lizard's",
	296: "Snake's",
	297: "Snake's",
	298: "Serpent's",
	299: "Serpent's",
	300: "Lizard's",
	301: "Lizard's",
	302: "Snake's",
	303: "Serpent's",
	304: "Lizard's",
	305: "Snake's",
	306: "Serpent's",
	307: "Serpent's",
	308: "Drake's",
	309: "Dragon's",
	310: "Dragon's",
	311: "Wyrm's",
	312: "Great Wyrm's",
	313: "Bahamut's",
	314: "Zircon",
	315: "Jacinth",
	316: "Turquoise",
	317: "Shimmering",
	318: "Shimmering",
	319: "Shimmering",
	320: "Shimmering",
	321: "Shimmering",
	322: "Shimmering",
	323: "Shimmering",
	324: "Rainbow",
	325: "Scintillating",
	326: "Prismatic",
	327: "Chromatic",
	328: "Shimmering",
	329: "Rainbow",
	330: "Scintillating",
	331: "Prismatic",
	332: "Chromatic",
	333: "Shimmering",
	334: "Rainbow",
	335: "Scintillating",
	336: "Shimmering",
	337: "Scintillating",
	338: "Azure",
	339: "Lapis",
	340: "Cobalt",
	341: "Sapphire",
	342: "Azure",
	343: "Lapis",
	344: "Cobalt",
	345: "Sapphire",
	346: "Azure",
	347: "Lapis",
	348: "Cobalt",
	349: "Sapphire",
	350: "Azure",
	351: "Lapis",
	352: "Lapis",
	353: "Cobalt",
	354: "Cobalt",
	355: "Sapphire",
	356: "Lapis Lazuli",
	357: "Sapphire",
	358: "Crimson",
	359: "Russet",
	360: "Garnet",
	361: "Ruby",
	362: "Crimson",
	363: "Russet",
	364: "Garnet",
	365: "Ruby",
	366: "Crimson",
	367: "Russet",
	368: "Garnet",
	369: "Ruby",
	370: "Russet",
	371: "Russet",
	372: "Garnet",
	373: "Garnet",
	374: "Ruby",
	375: "Garnet",
	376: "Ruby",
	377: "Tangerine",
	378: "Ocher",
	379: "Coral",
	380: "Amber",
	381: "Tangerine",
	382: "Ocher",
	383: "Coral",
	384: "Amber",
	385: "Tangerine",
	386: "Ocher",
	387: "Coral",
	388: "Amber",
	389: "Tangerine",
	390: "Ocher",
	391: "Ocher",
	392: "Coral",
	393: "Coral",
	394: "Amber",
	395: "Camphor",
	396: " Ambergris",
	397: "Beryl",
	398: "Viridian",
	399: "Jade",
	400: "Emerald",
	401: "Beryl",
	402: "Viridian",
	403: "Jade",
	404: "Emerald",
	405: "Beryl",
	406: "Viridian",
	407: "Jade",
	408: "Emerald",
	409: "Beryl",
	410: "Viridian",
	411: "Viridian",
	412: "Jade",
	413: "Jade",
	414: "Emerald",
	415: "Beryl",
	416: "Jade",
	417: "Triumphant",
	418: "Victorious",
	419: "Aureolin",
	420: "Mechanist's",
	421: "Artificer's",
	422: "Jeweler's",
	423: "Assamic",
	424: "Arcadian",
	425: "Unearthly",
	426: "Astral",
	427: "Elysian",
	428: "Celestial",
	429: "Diamond",
	430: "Fletcher's",
	431: "Acrobat's",
	432: "Harpoonist's",
	433: "Fletcher's",
	434: "Bowyer's",
	435: "Archer's",
	436: "Acrobat's",
	437: "Gymnast's",
	438: "Athlete's",
	439: "Harpoonist's",
	440: "Spearmaiden's",
	441: "Lancer's",
	442: "Burning",
	443: "Sparking",
	444: "Chilling",
	445: "Burning",
	446: "Blazing",
	447: "Volcanic",
	448: "Sparking",
	449: "Charged",
	450: "Powered",
	451: "Chilling",
	452: "Freezing",
	453: "Glacial",
	454: "Hexing",
	455: "Fungal",
	456: "Graverobber's",
	457: "Hexing",
	458: "Blighting",
	459: "Accursed",
	460: "Fungal",
	461: "Noxious",
	462: "Venomous",
	463: "Graverobber's",
	464: "Vodoun",
	465: "Golemlord's",
	466: "Lion Branded",
	467: "Captain's",
	468: "Preserver's",
	469: "Lion Branded",
	470: "Hawk Branded",
	471: "Rose Branded",
	472: "Captain's",
	473: "Commander's",
	474: "Marshal's",
	475: "Preserver's",
	476: "Warder's",
	477: "Guardian's",
	478: "Expert's",
	479: "Fanatic",
	480: "Sounding",
	481: "Expert's",
	482: "Veteran's",
	483: "Master's",
	484: "Fanatic",
	485: "Raging",
	486: "Furious",
	487: "Sounding",
	488: "Resonant",
	489: "Echoing",
	490: "Trainer's",
	491: "Spiritual",
	492: "Nature's",
	493: "Trainer's",
	494: "Caretaker's",
	495: "Keeper's",
	496: "Spiritual",
	497: "Feral",
	498: "Communal",
	499: "Nature's",
	500: "Terra's",
	501: "Gaea's",
	502: "Entrapping",
	503: "Mentalist's",
	504: "Shogukusha's",
	505: "Entrapping",
	506: "Trickster's",
	507: "Cunning",
	508: "Mentalist's",
	509: "Psychic",
	510: "Shadow",
	511: "Shogukusha's",
	512: "Sensei's",
	513: "Kenshi's",
	514: "Miocene",
	515: "Miocene",
	516: "Oligocene",
	517: "Oligocene",
	518: "Eocene",
	519: "Eocene",
	520: "Paleocene",
	521: "Paleocene",
	522: "Knave's",
	523: "Jack's",
	524: "Jester's",
	525: "Joker's",
	526: "Trump",
	527: "Loud",
	528: "Calling",
	529: "Yelling",
	530: "Shouting",
	531: "Screaming",
	532: "Paradox",
	533: "Paradox",
	534: "Robineye",
	535: "Sparroweye",
	536: "Falconeye",
	537: "Hawkeye",
	538: "Eagleeye",
	539: "Visionary",
	540: "Mnemonic",
	541: "Snowflake",
	542: "Shivering",
	543: "Boreal",
	544: "Hibernal",
	545: "Ember",
	546: "Smoldering",
	547: "Smoking",
	548: "Flaming",
	549: "Scorching",
	550: "Static",
	551: "Glowing",
	552: "Buzzing",
	553: "Arcing",
	554: "Shocking",
	555: "Septic",
	556: "Envenomed",
	557: "Corosive",
	558: "Toxic",
	559: "Pestilent",
	560: "Maiden's",
	561: "Valkyrie's",
	562: "Maiden's",
	563: "Valkyrie's",
	564: "Monk's",
	565: "Priest's",
	566: "Monk's",
	567: "Priest's",
	568: "Monk's",
	569: "Priest's",
	570: "Summoner's",
	571: "Necromancer's",
	572: "Summoner's",
	573: "Necromancer's",
	574: "Angel's",
	575: "Arch-Angel's",
	576: "Angel's",
	577: "Arch-Angel's",
	578: "Slayer's",
	579: "Berserker's",
	580: "Slayer's",
	581: "Berserker's",
	582: "Slayer's",
	583: "Berserker's",
	584: "Shaman's",
	585: "Hierophant's",
	586: "Shaman's",
	587: "Hierophant's",
	588: "Magekiller's",
	589: "Witch-hunter's",
	590: "Magekiller's",
	591: "Witch-hunter's",
	592: "Compact",
	593: "Thin",
	594: "Dense",
	595: "Consecrated",
	596: "Pure",
	597: "Sacred",
	598: "Hallowed",
	599: "Divine",
	600: "Pearl",
	601: "Crimson",
	602: "Red",
	603: "Sanguinary",
	604: "Bloody",
	605: "Red",
	606: "Sanguinary",
	607: "Red",
	608: "Jagged",
	609: "Forked",
	610: "Serrated",
	611: "Jagged",
	612: "Forked",
	613: "Jagged",
	614: "Snowflake",
	615: "Shivering",
	616: "Boreal",
	617: "Hibernal",
	618: "Snowflake",
	619: "Shivering",
	620: "Boreal",
	621: "Hibernal",
	622: "Snowflake",
	623: "Shivering",
	624: "Boreal",
	625: "Hibernal",
	626: "Ember",
	627: "Smoldering",
	628: "Smoking",
	629: "Flaming",
	630: "Ember",
	631: "Smoldering",
	632: "Smoking",
	633: "Flaming",
	634: "Ember",
	635: "Smoldering",
	636: "Smoking",
	637: "Flaming",
	638: "Static",
	639: "Glowing",
	640: "Arcing",
	641: "Shocking",
	642: "Static",
	643: "Glowing",
	644: "Arcing",
	645: "Shocking",
	646: "Static",
	647: "Glowing",
	648: "Arcing",
	649: "Shocking",
	650: "Septic",
	651: "Envenomed",
	652: "Toxic",
	653: "Pestilent",
	654: "Septic",
	655: "Envenomed",
	656: "Toxic",
	657: "Pestilent",
	658: "Septic",
	659: "Envenomed",
	660: "Toxic",
	661: "Pestilent",
	662: "Tireless",
	663: "Lizard's",
	664: "Azure",
	665: "Crimson",
	666: "Tangerine",
	667: "Beryl",
	668: "Godly",
	669: "Cruel",
}

var magicalSuffixes = map[uint64]string{
	1:   "Health",
	2:   "Protection",
	3:   "Absorption",
	4:   "Life",
	5:   "(nothing?)",
	6:   "Warding",
	7:   "the Sentinel",
	8:   "Guarding",
	9:   "Negation",
	10:  "(nothing?)",
	11:  "Piercing",
	12:  "Bashing",
	13:  "Puncturing",
	14:  "Thorns",
	15:  "Spikes",
	16:  "Fleadiness",
	17:  "Alacrity",
	18:  "Swiitness ",
	19:  "Quickness",
	20:  "Blocking",
	21:  "Deilecting",
	22:  "the Apprentice",
	23:  "the Magus",
	24:  "Frost",
	25:  "the Glacier",
	26:  "Frost",
	27:  "Warmth",
	28:  "Flame",
	29:  "Fire",
	30:  "Burning",
	31:  "Flame",
	32:  "Shook",
	33:  "Lightning",
	34:  "Thunder",
	35:  "Shock",
	36:  "Craftsmanship",
	37:  "Quality",
	38:  "Maiming",
	39:  "Slaying",
	40:  "Gore",
	41:  "Carnage",
	42:  "Slaughter",
	43:  "Maiming",
	44:  "Worth",
	45:  "Measure",
	46:  "Excellence",
	47:  "Petlctmance",
	48:  "Measure",
	49:  "Blight",
	50:  "Venom",
	51:  "Pestilence",
	52:  "Blight",
	53:  "Dextetity",
	54:  "Dextetity",
	55:  "Skill",
	56:  "Skill",
	57:  "Accuracy",
	58:  "Precision",
	59:  "Precision",
	60:  "Petlection",
	61:  "Balance",
	62:  "Stability",
	63:  "(nothing?)",
	64:  "Regenetation",
	65:  "Regenetation",
	66:  "Regenetation",
	67:  "Regrowth",
	68:  "Regrowth",
	69:  "Vileness",
	70:  "(nothing?)",
	71:  "Greed",
	72:  "Wealth",
	73:  "Chance",
	74:  "Fortune",
	75:  "Energy",
	76:  "Energy",
	77:  "the Mind",
	78:  "Brilliance",
	79:  "Sorcery",
	80:  "Wizardry",
	81:  "the Beat",
	82:  "Light",
	83:  "Radiance",
	84:  "the Sun",
	85:  "Life",
	86:  "the Jackal",
	87:  "the Fox",
	88:  "the Wolf",
	89:  "the Wolf",
	90:  "the Tiget",
	91:  "the Mammoth",
	92:  "the Mammoth",
	93:  "the Colosuss",
	94:  "the Leech",
	95:  "the Locust",
	96:  "the Bat",
	97:  "the Vampire",
	98:  "Defiance",
	99:  "Amelioration",
	100: "Remedy",
	101: "(nothing?)",
	102: "Simplicity",
	103: "Ease",
	104: "(nothing?)",
	105: "Strength",
	106: "Might",
	107: "the Ox",
	108: "the Ox",
	109: "the Giant",
	110: "the Giant",
	111: "the Titan",
	112: "Pacing",
	113: "Haste",
	114: "Speed",
	115: "Health",
	116: "Protection",
	117: "Absorption",
	118: "Life",
	119: "Life Everlasting",
	120: "Protection",
	121: "Absorption",
	122: "Life",
	123: "Anima",
	124: "Warding",
	125: "the Sentinel",
	126: "Guarding",
	127: "Negation",
	128: "the Sentinel",
	129: "Guarding",
	130: "Negation",
	131: "Coolness",
	132: "Incombustibility",
	133: "Amianthus",
	134: "Fire Quenching",
	135: "Coolness",
	136: "Incombustibility",
	137: "Amianthus",
	138: "Fire Quenching",
	139: "Faith",
	140: "Resistance",
	141: "Insulation",
	142: "Grounding",
	143: "the Dynamo",
	144: "Resistance",
	145: "Insulation",
	146: "Grounding",
	147: "the Dynamo",
	148: "Stoicism",
	149: "Warming",
	150: "Thawing",
	151: "the Dunes",
	152: "the Sirocco",
	153: "Warming",
	154: "Thawing",
	155: "the Dunes",
	156: "the Sirocco",
	157: "Desire",
	158: "Piercing",
	159: "Bashing",
	160: "Puncturing",
	161: "Thorns",
	162: "Spikes",
	163: "Razors",
	164: "Swords",
	165: "Malice",
	166: "Readiness",
	167: "Alacrity",
	168: "Swiftness",
	169: "Quickness",
	170: "Alacrity",
	171: "Fewer",
	172: "Blocking",
	173: "Deflecting",
	174: "the Apprentice",
	175: "the Magus",
	176: "Frost",
	177: "the Icicle",
	178: "the Glacier",
	179: "Winter",
	180: "Frost",
	181: "Frigidity",
	182: "Warmth",
	183: "Flame",
	184: "Fire",
	185: "Burning",
	186: "Incineration",
	187: "Flame",
	188: "Passion",
	189: "Shock",
	190: "Lightning",
	191: "Thunder",
	192: "Storms",
	193: "Shock",
	194: "Ennui",
	195: "Craftsmanship",
	196: "Quality",
	197: "Maiming",
	198: "Slaying",
	199: "Gore",
	200: "Damage",
	201: "Slaughter",
	202: "Butchery",
	203: "Evisceration",
	204: "Maiming",
	205: "Craftsmanship",
	206: "Craftsmanship",
	207: "Craftsmanship",
	208: "Quality",
	209: "Quality",
	210: "Maiming",
	211: "Maiming",
	212: "Craftsmanship",
	213: "Craftsmanship",
	214: "Quality",
	215: "Quality",
	216: "Maiming",
	217: "Craftsmanship",
	218: "Quality",
	219: "Maiming",
	220: "Ire",
	221: "Wrath",
	222: "Damage",
	223: "Worth",
	224: "Measure",
	225: "Excellence",
	226: "Performance",
	227: "Transcendence",
	228: "Worth",
	229: "Measure",
	230: "Excellence",
	231: "Performance",
	232: "Joyfulness",
	233: "Bliss",
	234: "Blight",
	235: "Venom",
	236: "Pestilence",
	237: "Anthrax",
	238: "Blight",
	239: "Envy",
	240: "Dexterity",
	241: "Skill",
	242: "Accuracy",
	243: "Precision",
	244: "Perfection",
	245: "Nirvana",
	246: "Dexterity",
	247: "Skill",
	248: "Accuracy",
	249: "Precision",
	250: "Perfection",
	251: "Dexterity",
	252: "Skill",
	253: "Accuracy",
	254: "Precision",
	255: "Dexterity",
	256: "Dexterity",
	257: "Dexterity",
	258: "Dexterity",
	259: "Dexterity",
	260: "Dexterity",
	261: "Daring",
	262: "Balance",
	263: "Equilibrium",
	264: "Stability",
	265: "Balance",
	266: "Balance",
	267: "Balance",
	268: "Truth",
	269: "Regeneration",
	270: "Regeneration",
	271: "Regeneration",
	272: "Regrowth",
	273: "Regrowth",
	274: "Revivification",
	275: "Honor",
	276: "Vileness",
	277: "Greed",
	278: "Wealth",
	279: "Greed",
	280: "Greed",
	281: "Greed",
	282: "Greed",
	283: "Greed",
	284: "Greed",
	285: "Avarice",
	286: "Chance",
	287: "Fortune",
	288: "Fortune",
	289: "Luck",
	290: "Fortune",
	291: "Good Luck",
	292: "Prosperity",
	293: "Energy",
	294: "the Mind",
	295: "Brilliance",
	296: "Sorcery",
	297: "Wizardry",
	298: "Enlightenment",
	299: "Energy",
	300: "the Mind",
	301: "Brilliance",
	302: "Sorcery",
	303: "Wizardry",
	304: "Energy",
	305: "the Mind",
	306: "Brilliance",
	307: "Sorcery",
	308: "Knowledge",
	309: "the Bear",
	310: "Light",
	311: "Radiance",
	312: "the Sun",
	313: "the Jackal",
	314: "the Fox",
	315: "the Wolf",
	316: "the Tiger",
	317: "the Mammoth",
	318: "the Colosuss",
	319: "the Squid",
	320: "the Whale",
	321: "the Jackal",
	322: "the Fox",
	323: "the Wolf",
	324: "the Tiger",
	325: "the Mammoth",
	326: "the Colosuss",
	327: "the Jackal",
	328: "the Fox",
	329: "the Wolf",
	330: "the Tiger",
	331: "the Mammoth",
	332: "Life",
	333: "Life",
	334: "Life",
	335: "Substinence",
	336: "Substinence",
	337: "Substinence",
	338: "Vita",
	339: "Vita",
	340: "Vita",
	341: "Life",
	342: "Life",
	343: "Substinence",
	344: "Substinence",
	345: "Vita",
	346: "Vita",
	347: "Life",
	348: "Substinence",
	349: "Vita",
	350: "Spirit",
	351: "Hope",
	352: "the Leech",
	353: "the Locust",
	354: "the Lamprey",
	355: "the Leech",
	356: "the Locust",
	357: "the Lamprey",
	358: "the Leech",
	359: "the Bat",
	360: "the Wraith",
	361: "the Vampire",
	362: "the Bat",
	363: "the Wraith",
	364: "the Vampire",
	365: "the Bat",
	366: "Defiance",
	367: "Amelioration",
	368: "Remedy",
	369: "Simplicity",
	370: "Ease",
	371: "Freedom",
	372: "Strength",
	373: "Might",
	374: "the Ox",
	375: "the Giant",
	376: "the Titan",
	377: "Atlus",
	378: "Strength",
	379: "Might",
	380: "the Us",
	381: "the Giant",
	382: "the Titan",
	383: "Strength",
	384: "Might",
	385: "the Us",
	386: "the Giant",
	387: "Strength",
	388: "Strength",
	389: "Strength",
	390: "Strength",
	391: "Strength",
	392: "Strength",
	393: "Virility",
	394: "Pacing",
	395: "Haste",
	396: "Speed",
	397: "Traveling",
	398: "Acceleration",
	399: "Inertia",
	400: "Inertia",
	401: "Inertia",
	402: "Self-Repair",
	403: "Fast Repair",
	404: "Ages",
	405: "Heplenishing",
	406: "Propagation",
	407: "the Kraken",
	408: "Memory",
	409: "the Elephant",
	410: "Power",
	411: "Grace",
	412: "Grace and Power",
	413: "the Yeti",
	414: "the Phoenix",
	415: "the Efreeti",
	416: "the Cobra",
	417: "the Elements",
	418: "Firebolts",
	419: "Firebolts",
	420: "Firebolts",
	421: "Charged Shield",
	422: "Charged Shield",
	423: "Charged Shield",
	424: "Icebolt",
	425: "Frozen Armor",
	426: "Static Field",
	427: "Telekinesis",
	428: "Frost Shield",
	429: "Ice Blast",
	430: "Blaze",
	431: "Fire Ball",
	432: "Nova",
	433: "Nova",
	434: "Nova Shield",
	435: "Nova Shield",
	436: "Nova Shield",
	437: "Lightning",
	438: "Lightning",
	439: "Shiver Armor",
	440: "Fire Wall",
	441: "Enchant",
	442: "Chain Lightning",
	443: "Chain Lightning",
	444: "Chain Lightning",
	445: "Teleport Shield",
	446: "Teleport Shield",
	447: "Teleport Shield",
	448: "Glacial Spike",
	449: "Meteor",
	450: "Thunder Storm",
	451: "Energy Shield",
	452: "Blizzard",
	453: "Chilling Armor",
	454: "Hydra Shield",
	455: "Frozen ler",
	456: "Dawn",
	457: "Sunlight",
	458: "Magic Arrows",
	459: "Magic Arrows",
	460: "Fire Arrows",
	461: "Fire Arrows",
	462: "lnner Sight",
	463: "Inner Sight",
	464: "Jabbing",
	465: "Jabbing",
	466: "Cold Arrows",
	467: "Cold Arrows",
	468: "Multiple Shot",
	469: "Multiple Shot",
	470: "Power Strike",
	471: "Power Strike",
	472: "Poison Jab",
	473: "Poison Jab",
	474: "Exploding Arrows",
	475: "Exploding Arrows",
	476: "Slow Missiles",
	477: "Slow Missiles",
	478: "lmpaling Strike",
	479: "lmpaling Strike",
	480: "Lightning Javelin",
	481: "Lightning Javelin",
	482: "Ice Arrows",
	483: "Ice Arrows",
	484: "Guided Arrows",
	485: "Guided Arrows",
	486: "Charged Strike",
	487: "Charged Strike",
	488: "Plague Jab",
	489: "Plague Jab",
	490: "lmmolating Arrows",
	491: "lmmolating Arrows",
	492: "Fending",
	493: "Fending",
	494: "Freezing Arrows",
	495: "Freezing Arrows",
	496: "Lightning Strike",
	497: "Lightning Strike",
	498: "Lightning Fury",
	499: "Lightning Fury",
	500: "Fire Bolts",
	501: "Fire Bolts",
	502: "Charged Bolts",
	503: "Charged Bolts",
	504: "Ice Bolts",
	505: "Ice Bolts",
	506: "Frozen Armor",
	507: "Frozen Armor",
	508: "Static Field",
	509: "Static Field",
	510: "Telekinesis",
	511: "Telekinesis",
	512: "Frost Novas",
	513: "Frost Novas",
	514: "Ice Blasts",
	515: "Ice Blasts",
	516: "Blazing",
	517: "Blazing",
	518: "Fire Balls",
	519: "Fire Balls",
	520: "Novas",
	521: "Novas",
	522: "Lightning",
	523: "Lightning",
	524: "Shiver Armor",
	525: "Shiver Armor",
	526: "Fire Walls",
	527: "Fire Walls",
	528: "Enchantment",
	529: "Enchantment",
	530: "Chain Lightning",
	531: "Chain Lightning",
	532: "Teleportation",
	533: "Teleportation",
	534: "Glacial Spikes",
	535: "Glacial Spikes",
	536: "Meteors",
	537: "Meteors",
	538: "Thunder Storm",
	539: "Thunder Storm",
	540: "Energy Shield",
	541: "Energy Shield",
	542: "Blizzards",
	543: "Blizzards",
	544: "Chilling Armor",
	545: "Chilling Armor",
	546: "Hydras",
	547: "Hydras",
	548: "Frozen Orbs",
	549: "Frozen Orbs",
	550: "Amplify Damage",
	551: "Amplify Damage",
	552: "Teeth",
	553: "Teeth",
	554: "Bone Armor",
	555: "Bone Armor",
	556: "Raise Skeletons",
	557: "Raise Skeletons",
	558: "Dim Vision",
	559: "Dim Vision",
	560: "Weaken",
	561: "Weaken",
	562: "Poison Dagger",
	563: "Poison Dagger",
	564: "Corpse Explosions",
	565: "Corpse Explosions",
	566: "Clay Golem Summoning",
	567: "Clay Golem Summoning",
	568: "Iron Maiden",
	569: "Iron Maiden",
	570: "Terror",
	571: "Terror",
	572: "Bone Walls",
	573: "Bone Walls",
	574: "Raise Skeletal Mages",
	575: "Raise Skeletal Mages",
	576: "Confusion",
	577: "Confusion",
	578: "Life Tap",
	579: "Life Tap",
	580: "Poison Explosion",
	581: "Poison Explosion",
	582: "Bone Spears",
	583: "Bone Spears",
	584: "Blood Golem Summoning",
	585: "Blood Golem Summoning",
	586: "Attraction",
	587: "Attraction",
	588: "Decrepification",
	589: "Decrepification",
	590: "Bone Imprisonment",
	591: "Bone Imprisonment",
	592: "Iron Golem Creation",
	593: "Iron Golem Creation",
	594: "Lower Resistance",
	595: "Lower Flesistance",
	596: "Poison Novas",
	597: "Poison Novas",
	598: "Bone Spirits",
	599: "Bone Spirits",
	600: "Fire Golem Summoning",
	601: "Fire Golem Summoning",
	602: "Revivification",
	603: "Revivification",
	604: "Sacrifice",
	605: "Sacrifice",
	606: "Holy Bolts",
	607: "Holy Bolts",
	608: "Zeal",
	609: "Zeal",
	610: "Vengeance",
	611: "Vengeance",
	612: "Blessed Hammers",
	613: "Blessed Hammers",
	614: "Conversion",
	615: "Conversion",
	616: "Fist of the Heavens",
	617: "Fist of the Heavens",
	618: "Bashing",
	619: "Bashing",
	620: "Howling",
	621: "Howling",
	622: "Potion Finding",
	623: "Potion Finding",
	624: "Taunting",
	625: "Taunting",
	626: "Shouting",
	627: "Shouting",
	628: "Stunning",
	629: "Stunning",
	630: "Item Finding",
	631: "Item Finding",
	632: "Concentration",
	633: "Concentration",
	634: "Battle Cry",
	635: "Battle Cry",
	636: "Battle Orders",
	637: "Battle Orders",
	638: "Grim Ward",
	639: "Grim Ward",
	640: "War Cry",
	641: "War Cry",
	642: "Battle Command",
	643: "Battle Command",
	644: "Firestorms",
	645: "Firestorms",
	646: "Molten Boulders",
	647: "Molten Boulders",
	648: "Eruption",
	649: "Eruption",
	650: "Cyclone Armor",
	651: "Cyclone Armor",
	652: "Twister",
	653: "Twister",
	654: "Volcano",
	655: "Volcano",
	656: "Tornado",
	657: "Tornado",
	658: "Armageddon",
	659: "Armageddon",
	660: "Hurricane",
	661: "Hurricane",
	662: "Damage Amplification",
	663: "the Icicle",
	664: "the Glacier",
	665: "Fire",
	666: "Burning",
	667: "Lightning",
	668: "Thunder",
	669: "Daring",
	670: "Daring",
	671: "Knowledge",
	672: "Knowledge",
	673: "Virility",
	674: "Virility",
	675: "Readiness",
	676: "Craftsmanship",
	677: "Quality",
	678: "Maiming",
	679: "Craftsmanship",
	680: "Quality",
	681: "Craftsmanship",
	682: "Blight",
	683: "Venom",
	684: "Pestilence",
	685: "Anthrax",
	686: "Blight",
	687: "Venom",
	688: "Pestilence",
	689: "Anthrax",
	690: "Blight",
	691: "Venom",
	692: "Pestilence",
	693: "Anthrax",
	694: "Frost",
	695: "the Icicle",
	696: "the Glacier",
	697: "Winter",
	698: "Frost",
	699: "the Icicle",
	700: "the Glacier",
	701: "Winter",
	702: "Frost",
	703: "the Icicle",
	704: "the Glacier",
	705: "Winter",
	706: "Flame",
	707: "Fire",
	708: "Burning",
	709: "Incineration",
	710: "Flame",
	711: "Fire",
	712: "Burning",
	713: "Incineration",
	714: "Flame",
	715: "Fire",
	716: "Burning",
	717: "Incineration",
	718: "Shock",
	719: "Lightning",
	720: "Thunder",
	721: "Storms",
	722: "Shock",
	723: "Lightning",
	724: "Thunder",
	725: "Storms",
	726: "Shock",
	727: "Lightning",
	728: "Thunder",
	729: "Storms",
	730: "Dexterity",
	731: "Dexterity",
	732: "Strength",
	733: "Strength",
	734: "Thorns",
	735: "Frost",
	736: "Flame",
	737: "Blight",
	738: "Shock",
	739: "Regeneration",
	740: "Energy",
	741: "Light",
	742: "the Leech",
	743: "the Locust",
	744: "the Lamprey",
	745: "the Bat",
	746: "the Wraith",
	747: "the Vampire",
}
