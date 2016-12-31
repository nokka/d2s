package d2s

var skillMap = [...]string{
	"attack",
	"kick",
	"throw_item",
	"unsummon",
	"left_hand_throw",
	"left_hand_swing",
	"magic_arrow",
	"fire_arrow",
	"inner_sight",
	"critical_strike",
	"jab",
	"cold_arrow",
	"multiple_shot",
	"dodge",
	"power_strike",
	"poison_javelin",
	"exploding_arrow",
	"slow_missiles",
	"avoid",
	"impale",
	"lightning_bolt",
	"ice_arrow",
	"guided_arrow",
	"penetrate",
	"charged_strike",
	"plague_javelin",
	"strafe",
	"immolation_arrow",
	"dopplezon",
	"evade",
	"fend",
	"freezing_arrow",
	"valkyrie",
	"pierce",
	"lightning_strike",
	"lightning_fury",
	"fire_bolt",
	"warmth",
	"charged_bolt",
	"ice_bolt",
	"frozen_armor",
	"inferno",
	"static_field",
	"telekinesis",
	"frost_nova",
	"ice_blast",
	"blaze",
	"fire_ball",
	"nova",
	"lightning",
	"shiver_armor",
	"fire_wall",
	"enchant",
	"chain_lightning",
	"teleport",
	"glacial_spike",
	"meteor",
	"thunder_storm",
	"energy_shield",
	"blizzard",
	"chilling_armor",
	"fire_mastery",
	"hydra",
	"lightning_mastery",
	"frozen_orb",
	"cold_mastery",
	"amplify_damage",
	"teeth",
	"bone_armor",
	"skeleton_mastery",
	"raise_skeleton",
	"dim_vision",
	"weaken",
	"poison_dagger",
	"corpse_explosion",
	"clay_golem",
	"iron_maiden",
	"terror",
	"bone_wall",
	"golem_mastery",
	"raise_skeletal_mage",
	"confuse",
	"life_tap",
	"poison_explosion",
	"bone_spear",
	"bloodgolem",
	"attract",
	"decrepify",
	"bone_prison",
	"summon_resist",
	"irongolem",
	"lower_resist",
	"poison_nova",
	"bone_spirit",
	"firegolem",
	"revive",
	"sacrifice",
	"smite",
	"might",
	"prayer",
	"resist_fire",
	"holy_bolt",
	"holy_fire",
	"thorns",
	"defiance",
	"resist_cold",
	"zeal",
	"charge",
	"blessed_aim",
	"cleansing",
	"resist_lightning",
	"vengeance",
	"blessed_hammer",
	"concentration",
	"holy_freeze",
	"vigor",
	"conversion",
	"holy_shield",
	"holy_shock",
	"sanctuary",
	"meditation",
	"fist_of_the_heavens",
	"fanaticism",
	"conviction",
	"redemption",
	"salvation",
	"bash",
	"sword_mastery",
	"axe_mastery",
	"mace_mastery",
	"howl",
	"find_potion",
	"leap",
	"double_swing",
	"pole_arm_mastery",
	"throwing_mastery",
	"spear_mastery",
	"taunt",
	"shout",
	"stun",
	"double_throw",
	"increased_stamina",
	"find_item",
	"leap_attack",
	"concentrate",
	"iron_skin",
	"battle_cry",
	"frenzy",
	"increased_speed",
	"battle_orders",
	"grim_ward",
	"whirlwind",
	"berserk",
	"natural_resistance",
	"war_cry",
	"battle_command",
	"fire_hit",
	"unholybolt",
	"skeletonraise",
	"maggotegg",
	"shamanfire",
	"magottup",
	"magottdown",
	"magottlay",
	"andrialspray",
	"jump",
	"swarm_move",
	"nest",
	"quick_strike",
	"vampirefireball",
	"vampirefirewall",
	"vampiremeteor",
	"gargoyletrap",
	"spiderlay",
	"vampireheal",
	"vampireraise",
	"submerge",
	"fetishaura",
	"fetishinferno",
	"zakarumheal",
	"emerge",
	"resurrect",
	"bestow",
	"missileskill1",
	"monteleport",
	"primelightning",
	"primebolt",
	"primeblaze",
	"primefirewall",
	"primespike",
	"primeicenova",
	"primepoisonball",
	"primepoisonnova",
	"diablight",
	"diabcold",
	"diabfire",
	"fingermagespider",
	"diabwall",
	"diabrun",
	"diabprison",
	"poisonballtrap",
	"andypoisonbolt",
	"hireablemissile",
	"desertturret",
	"arcanetower",
	"monblizzard",
	"mosquito",
	"cursedballtrapright",
	"cursedballtrapleft",
	"monfrozenarmor",
	"monbonearmor",
	"monbonespirit",
	"moncursecast",
	"hellmeteor",
	"regurgitatoreat",
	"monfrenzy",
	"queendeath",
	"scroll_of_identify",
	"book_of_identify",
	"scroll_of_townportal",
	"book_of_townportal",
	/*221	raven
	  222	plague_poppy
	  223	wearwolf
	  224	shape_shifting
	  225	firestorm
	  226	oak_sage
	  227	summon_spirit_wolf
	  228	wearbear
	  229	molten_boulder
	  230	arctic_blast
	  231	cycle_of_life
	  232	feral_rage
	  233	maul
	  234	eruption
	  235	cyclone_armor
	  236	heart_of_wolverine
	  237	summon_fenris
	  238	rabies
	  239	fire_claws
	  240	twister
	  241	vines
	  242	hunger
	  243	shock_wave
	  244	volcano
	  245	tornado
	  246	spirit_of_barbs
	  247	summon_grizzly
	  248	fury
	  249	armageddon
	  250	hurricane
	  251	fire_trauma
	  252	claw_mastery
	  253	psychic_hammer
	  254	tiger_strike
	  255	dragon_talon
	  256	shock_field
	  257	blade_sentinel
	  258	quickness
	  259	fists_of_fire
	  260	dragon_claw
	  261	charged_bolt_sentry
	  262	wake_of_fire_sentry
	  263	weapon_block
	  264	cloak_of_shadows
	  265	cobra_strike
	  266	blade_fury
	  267	fade
	  268	shadow_warrior
	  269	claws_of_thunder
	  270	dragon_tail
	  271	lightning_sentry
	  272	inferno_sentry
	  273	mind_blast
	  274	blades_of_ice
	  275	dragon_flight
	  276	death_sentry
	  277	blade_shield
	  278	venom
	  279	shadow_master
	  280	royal_strike
	  281	wake_of_destruction_sentry
	  282	imp_inferno
	  283	imp_fireball
	  284	baal_taunt
	  285	baal_corpse_explode
	  286	baal_monster_spawn
	  287	catapult_charged_ball
	  288	catapult_spike_ball
	  289	suck_blood
	  290	cry_help
	  291	healing_vortex
	  292	teleport_2
	  293	self-resurrect
	  294	vine_attack
	  295	overseer_whip
	  296	barbs_aura
	  297	wolverine_aura
	  298	oak_sage_aura
	  299	imp_fire_missile
	  300	impregnate
	  301	siege_beast_stomp
	  302	minionspawner
	  303	catapultblizzard
	  304	catapultplague
	  305	catapultmeteor
	  306	boltsentry
	  307	corpsecycler
	  308	deathmaul
	  309	defense_curse
	  310	blood_mana
	  311	mon_inferno_sentry
	  312	mon_death_sentry
	  313	sentry_lightning
	  314	fenris_rage
	  315	baal_tentacle
	  316	baal_nova
	  317	baal_inferno
	  318	baal_cold_missiles
	  319	megademoninferno
	  320	evilhutspawner
	  321	countessfirewall
	  322	impbolt
	  323	horror_arctic_blast
	  324	death_sentry_ltng
	  325	vinecycler
	  326	bearsmite
	  327	resurrect2
	  328	bloodlordfrenzy
	  329	baal_teleport
	  330	imp_teleport
	  331	baal_clone_teleport
	  332	zakarumlightning
	  333	vampiremissile
	  334	mephistomissile
	  335	doomknightmissile
	  336	roguemissile
	  337	hydramissile
	  338	necromagemissile
	  339	monbow
	  340	monfirearrow
	  341	moncoldarrow
	  342	monexplodingarrow
	  343	monfreezingarrow
	  344	monpowerstrike
	  345	succubusbolt
	  346	mephfrostnova
	  347	monicespear
	  348	shamanice
	  349	diablogeddon
	  350	delerium_change
	  351	nihlathakcorpseexplosion
	  352	serpentcharge
	  353	trap_nova
	  354	unholyboltex
	  355	shamanfireex
	  356	imp_fire_missile_ex*/
}
