package minecraft

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// spigot.yml
type Spigot struct {
	Settings      SpigotSettings                 `yaml:"settings" json:"settings"`
	Messages      SpigotMessages                 `yaml:"messages" json:"messages"`
	Commands      SpigotCommands                 `yaml:"commands" json:"commands"`
	WorldSettings map[string]SpigotWorldSettings `yaml:"world-settings" json:"world-settings"`
	Advancements  SpigotAdvancements             `yaml:"advancements" json:"advancements"`
	Players       SpigotPlayers                  `yaml:"players" json:"players"`
	ConfigVersion int                            `yaml:"config-version" json:"config-version"`
	Stats         SpigotStats                    `yaml:"stats" json:"stats"`
}

type SpigotSettings struct {
	Debug                     bool            `yaml:"debug" json:"debug" json:"debug" json:"debug"`
	Bungeecord                bool            `yaml:"bungeecord" json:"bungeecord"`
	SampleCount               int             `yaml:"sample-count" json:"sample-count"`
	PlayerShuffle             int             `yaml:"player-shuffle" json:"player-shuffle"`
	UserCacheSize             int             `yaml:"user-cache-size" json:"user-cache-size"`
	SaveUserCacheOnStopOnly   bool            `yaml:"save-user-cache-on-stop-only" json:"save-user-cache-on-stop-only"`
	MovedWronglyThreshold     float64         `yaml:"moved-wrongly-threshold" json:"moved-wrongly-threshold"`
	MovedTooQuicklyMultiplier float64         `yaml:"moved-too-quickly-multiplier" json:"moved-too-quickly-multiplier"`
	TimeoutTime               int             `yaml:"timeout-time" json:"timeout-time"`
	RestartOnCrash            bool            `yaml:"restart-on-crash" json:"restart-on-crash"`
	RestartScript             string          `yaml:"restart-script" json:"restart-script"`
	NettyThreads              int             `yaml:"netty-threads" json:"netty-threads"`
	Attribute                 SpigotAttribute `yaml:"attribute" json:"attribute"`
	LogVillagerDeaths         bool            `yaml:"log-villager-deaths" json:"log-villager-deaths"`
	LogNamedDeaths            bool            `yaml:"log-named-deaths" json:"log-named-deaths"`
}

type SpigotAttribute struct {
	MaxHealth     Max `yaml:"maxHealth" json:"maxHealth"`
	MovementSpeed Max `yaml:"movementSpeed" json:"movementSpeed"`
	AttackDamage  Max `yaml:"attackDamage" json:"attackDamage"`
}

type Max struct {
	Max float64 `yaml:"max" json:"max"`
}

type SpigotMessages struct {
	Whitelist      string `yaml:"whitelist" json:"whitelist"`
	UnknownCommand string `yaml:"unknown-command" json:"unknown-command"`
	ServerFull     string `yaml:"server-full" json:"server-full"`
	OutdatedClient string `yaml:"outdated-client" json:"outdated-client"`
	OutdatedServer string `yaml:"outdated-server" json:"outdated-server"`
	Restart        string `yaml:"restart" json:"restart"`
}

type SpigotCommands struct {
	ReplaceCommand            []string `yaml:"replace-command" json:"replace-command"`
	SpamExclusions            []string `yaml:"spam-exclusions" json:"spam-exclusions"`
	SilentCommandblockConsole bool     `yaml:"silent-commandblock-console" json:"silent-commandblock-console"`
	Log                       bool     `yaml:"log" json:"log"`
	TabComplete               int      `yaml:"tab-complete" json:"tab-complete"`
	SendNamespaced            bool     `yaml:"send-namespaced" json:"send-namespaced"`
}

type SpigotWorldSettings struct {
	BelowZeroGenerationInExistingChunks bool                        `yaml:"below-zero-generation-in-existing-chunks" json:"below-zero-generation-in-existing-chunks"`
	Verbose                             bool                        `yaml:"verbose" json:"verbose"`
	Growth                              SpigotWorldGrowth           `yaml:"growth" json:"growth"`
	MergeRadius                         SpigotMergeRadius           `yaml:"merge-radius" json:"merge-radius"`
	MobSpawnRange                       int                         `yaml:"mob-spawn-range" json:"mob-spawn-range"`
	EntityActivationRange               SpigotEntityActivationRange `yaml:"entity-activation-range" json:"entity-activation-range"`
	EntityTrackingRange                 SpigotEntityTrackingRange   `yaml:"entity-tracking-range" json:"entity-tracking-range"`
	TicksPer                            SpigotTicksPer              `yaml:"ticks-per" json:"ticks-per"`
	HopperAmount                        int                         `yaml:"hopper-amount" json:"hopper-amount"`
	HopperCanLoadChunks                 bool                        `yaml:"hopper-can-load-chunks" json:"hopper-can-load-chunks"`
	DragonDeathSoundRadius              int                         `yaml:"dragon-death-sound-radius" json:"dragon-death-sound-radius"`
	SeedVillage                         int                         `yaml:"seed-village" json:"seed-village"`
	SeedDesert                          int                         `yaml:"seed-desert" json:"seed-desert"`
	SeedIgloo                           int                         `yaml:"seed-igloo" json:"seed-igloo"`
	SeedJungle                          int                         `yaml:"seed-jungle" json:"seed-jungle"`
	SeedSwamp                           int                         `yaml:"seed-swamp" json:"seed-swamp"`
	SeedMonument                        int                         `yaml:"seed-monument" json:"seed-monument"`
	SeedShipwreck                       int                         `yaml:"seed-shipwreck" json:"seed-shipwreck"`
	SeedOcean                           int                         `yaml:"seed-ocean" json:"seed-ocean"`
	SeedOutpost                         int                         `yaml:"seed-outpost" json:"seed-outpost"`
	SeedEndcity                         int                         `yaml:"seed-endcity" json:"seed-endcity"`
	SeedSlime                           int                         `yaml:"seed-slime" json:"seed-slime"`
	SeedNether                          int                         `yaml:"seed-nether" json:"seed-nether"`
	SeedMansion                         int                         `yaml:"seed-mansion" json:"seed-mansion"`
	SeedFossil                          int                         `yaml:"seed-fossil" json:"seed-fossil"`
	SeedPortal                          int                         `yaml:"seed-portal" json:"seed-portal"`
	Hunger                              SpigotHunger                `yaml:"hunger" json:"hunger"`
	MaxTntPerTick                       int                         `yaml:"max-tnt-per-tick" json:"max-tnt-per-tick"`
	MaxTickTime                         SpigotMaxTickTime           `yaml:"max-tick-time" json:"max-tick-time"`
	ItemDespawnRate                     int                         `yaml:"item-despawn-rate" json:"item-despawn-rate"`
	ViewDistance                        int                         `yaml:"view-distance" json:"view-distance"`
	SimulationDistance                  int                         `yaml:"simulation-distance" json:"simulation-distance"`
	ThunderChance                       int                         `yaml:"thunder-chance" json:"thunder-chance"`
	EnableZombiePigmenPortalSpawns      bool                        `yaml:"enable-zombie-pigmen-portal-spawns" json:"enable-zombie-pigmen-portal-spawns"`
	WitherSpawnSoundRadius              int                         `yaml:"wither-spawn-sound-radius" json:"wither-spawn-sound-radius"`
	HangingTickFrequency                int                         `yaml:"hanging-tick-frequency" json:"hanging-tick-frequency"`
	ArrowDespawnRate                    int                         `yaml:"arrow-despawn-rate" json:"arrow-despawn-rate"`
	ZombieAggressiveTowardsVillager     bool                        `yaml:"zombie-aggressive-toward-villager" json:"zombie-aggressive-toward-villager"`
	TridentDespawnRate                  int                         `yaml:"trident-despawn-rate" json:"trident-despawn-rate"`
	NerfSpawnerMobs                     bool                        `yaml:"nerf-spawner-mobs" json:"nerf-spawner-mobs"`
	EndPortalSoundRadius                int                         `yaml:"end-portal-sound-radius" json:"end-portal-sound-radius"`
}

type SpigotWorldGrowth struct {
	CactusModifier     int `yaml:"cactus-modifier" json:"cactus-modifier"`
	CaneModifier       int `yaml:"cane-modifier" json:"cane-modifier"`
	MelonModifier      int `yaml:"Melon-modifier" json:"Melon-modifier"`
	MushroomModifier   int `yaml:"mushroom-modifier" json:"mushroom-modifier"`
	PumpkinModifier    int `yaml:"pumpkin-modifier" json:"pumpkin-modifier"`
	SaplingModifier    int `yaml:"sapling-modifier" json:"sapling-modifier"`
	BeetrootModifier   int `yaml:"Beetroot-modifier" json:"Beetroot-modifier"`
	CarrotModifier     int `yaml:"carrot-modifier" json:"carrot-modifier"`
	PotatoModifier     int `yaml:"potato-modifier" json:"potato-modifier"`
	WheatModifier      int `yaml:"wheat-modifier" json:"wheat-modifier"`
	NetherwartModifier int `yaml:"netherwart-modifier" json:"netherwart-modifier"`
	VineModifier       int `yaml:"vine-modifier" json:"vine-modifier"`
	CocoaModifier      int `yaml:"cocoa-modifier" json:"cocoa-modifier"`
	BambooModifier     int `yaml:"bamboo-modifier" json:"bamboo-modifier"`
	SweetberryModifier int `yaml:"sweetberry-modifier" json:"sweetberry-modifier"`
	KelpModifier       int `yaml:"kelp-modifier" json:"kelp-modifier"`
}

type SpigotMergeRadius struct {
	Exp  float64 `yaml:"exp" json:"exp"`
	Item float64 `yaml:"item" json:"item"`
}

type SpigotEntityActivationRange struct {
	Animals               int  `yaml:"animals" json:"animals"`
	Monsters              int  `yaml:"monsters" json:"monsters"`
	Raiders               int  `yaml:"raiders" json:"raiders"`
	Misc                  int  `yaml:"misc" json:"misc"`
	TickInactiveVillagers bool `yaml:"tick-inactive-villagers" json:"tick-inactive-villagers"`
	IgnoreSpectators      bool `yaml:"ignore-spectators" json:"ignore-spectators"`
}

type SpigotEntityTrackingRange struct {
	Players  int `yaml:"players" json:"players"`
	Monsters int `yaml:"monsters" json:"monsters"`
	Animals  int `yaml:"animals" json:"animals"`
	Misc     int `yaml:"misc" json:"misc"`
	Other    int `yaml:"other" json:"other"`
}

type SpigotTicksPer struct {
	HopperTransfer int `yaml:"hopper-transfer" json:"hopper-transfer"`
	HopperCheck    int `yaml:"hopper-check" json:"hopper-check"`
}

type SpigotHunger struct {
	JumpWalkExhaustion   float64 `yaml:"jump-walk-exhaustion" json:"jump-walk-exhaustion"`
	JumpSprintExhaustion float64 `yaml:"jump-sprint-exhaustion" json:"jump-sprint-exhaustion"`
	CombatExhaustion     float64 `yaml:"combat-exhaustion" json:"combat-exhaustion"`
	RegenExhaustion      float64 `yaml:"regen-exhaustion" json:"regen-exhaustion"`
	SwimMultiplier       float64 `yaml:"swim-multiplier" json:"swim-multiplier"`
	SprintMultiplier     float64 `yaml:"sprint-multiplier" json:"sprint-multiplier"`
	OtherMultiplier      float64 `yaml:"other-multiplier" json:"other-multiplier"`
}

type SpigotMaxTickTime struct {
	Tile   int `yaml:"tile" json:"tile"`
	Entity int `yaml:"entity" json:"entity"`
}

type SpigotAdvancements struct {
	DisableSaving bool     `yaml:"disable-saving" json:"disable-saving"`
	Disabled      []string `yaml:"disabled" json:"disabled"`
}

type SpigotPlayers struct {
	DisableSaving bool `yaml:"disable-saving" json:"disable-saving"`
}

type SpigotStats struct {
	DisableSaving bool   `yaml:"disable-saving" json:"disable-saving"`
	ForcedStats   string `yaml:"forced-stats" json:"forced-stats"`
}

func (s Spigot) ToYaml() (string, error) {
	b, err := yaml.Marshal(s)
	if err != nil {
		fmt.Println("Failed to send spigot.yaml toYaml()")
		return "", err
	}
	// str := strings.ReplaceAll(string(b), ": ", "=")
	// str = strings.ReplaceAll(str, "'", "")
	// str = strings.ReplaceAll(str, "\"", "")

	return string(b), nil
}

func BoilerplateSpigotYaml() Spigot {
	return Spigot{
		Settings: SpigotSettings{
			Debug:                     false,
			Bungeecord:                true,
			SampleCount:               12,
			PlayerShuffle:             0,
			UserCacheSize:             1000,
			SaveUserCacheOnStopOnly:   false,
			MovedWronglyThreshold:     0.0625,
			MovedTooQuicklyMultiplier: 10.0,
			TimeoutTime:               60,
			RestartOnCrash:            true,
			RestartScript:             "./start",
			NettyThreads:              4,
			Attribute: SpigotAttribute{
				MaxHealth: Max{
					Max: 2048.0,
				},
				MovementSpeed: Max{
					Max: 2048.0,
				},
				AttackDamage: Max{
					Max: 2048.0,
				},
			},
			LogVillagerDeaths: true,
			LogNamedDeaths:    true,
		},
		Messages: SpigotMessages{
			Whitelist:      "You are not whitelisted on this server!",
			UnknownCommand: "Unknown command. Type /help for help.",
			ServerFull:     "The server is full!",
			OutdatedClient: "Outdated client! Please use {0}",
			OutdatedServer: "Outdated server! I'm on still on {0}",
			Restart:        "Server is restarting",
		},
		Commands: SpigotCommands{
			ReplaceCommand: []string{
				"setBlock",
				"summon",
				"testforblock",
				"tellraw",
			},
			SpamExclusions: []string{
				"/skill",
			},
			SilentCommandblockConsole: false,
			Log:                       true,
			TabComplete:               0,
			SendNamespaced:            true,
		},
		WorldSettings: map[string]SpigotWorldSettings{
			"default": SpigotWorldSettings{
				BelowZeroGenerationInExistingChunks: true,
				Verbose:                             true,
				Growth: SpigotWorldGrowth{
					CactusModifier:     100,
					CaneModifier:       100,
					MelonModifier:      100,
					MushroomModifier:   100,
					PumpkinModifier:    100,
					SaplingModifier:    100,
					BeetrootModifier:   100,
					CarrotModifier:     100,
					PotatoModifier:     100,
					WheatModifier:      100,
					NetherwartModifier: 100,
					VineModifier:       100,
					CocoaModifier:      100,
					BambooModifier:     100,
					SweetberryModifier: 100,
					KelpModifier:       100,
				},
				MergeRadius: SpigotMergeRadius{
					Exp:  3.0,
					Item: 2.5,
				},
				MobSpawnRange: 6,
				EntityActivationRange: SpigotEntityActivationRange{
					Animals:               32,
					Monsters:              32,
					Raiders:               32,
					Misc:                  16,
					TickInactiveVillagers: true,
					IgnoreSpectators:      false,
				},
				EntityTrackingRange: SpigotEntityTrackingRange{
					Players:  48,
					Animals:  48,
					Monsters: 48,
					Misc:     32,
					Other:    64,
				},
				TicksPer: SpigotTicksPer{
					HopperTransfer: 8,
					HopperCheck:    1,
				},
				HopperAmount:           1,
				HopperCanLoadChunks:    false,
				DragonDeathSoundRadius: 0,
				SeedVillage:            10387312,
				SeedDesert:             14357617,
				SeedIgloo:              14357618,
				SeedJungle:             14357619,
				SeedSwamp:              14357620,
				SeedMonument:           10387313,
				SeedShipwreck:          165745295,
				SeedOcean:              14357621,
				SeedOutpost:            165745296,
				SeedEndcity:            10387313,
				SeedSlime:              987234911,
				SeedNether:             30084232,
				SeedMansion:            10387319,
				SeedFossil:             14357921,
				SeedPortal:             34222645,
				Hunger: SpigotHunger{
					JumpWalkExhaustion:   0.05,
					JumpSprintExhaustion: 0.2,
					CombatExhaustion:     0.1,
					RegenExhaustion:      6.0,
					SwimMultiplier:       0.01,
					SprintMultiplier:     0.1,
					OtherMultiplier:      0.0,
				},
				MaxTntPerTick: 100,
				MaxTickTime: SpigotMaxTickTime{
					Tile:   50,
					Entity: 50,
				},
				ItemDespawnRate:                 6000,
				ViewDistance:                    10,
				SimulationDistance:              10,
				ThunderChance:                   100000,
				EnableZombiePigmenPortalSpawns:  true,
				WitherSpawnSoundRadius:          0,
				HangingTickFrequency:            100,
				ArrowDespawnRate:                1200,
				ZombieAggressiveTowardsVillager: true,
				NerfSpawnerMobs:                 false,
				EndPortalSoundRadius:            0,
			},
		},
		Advancements: SpigotAdvancements{
			DisableSaving: false,
			Disabled: []string{
				"minecraft:story/disabled",
			},
		},
		Players: SpigotPlayers{
			DisableSaving: false,
		},
		ConfigVersion: 12,
		Stats: SpigotStats{
			DisableSaving: false,
			ForcedStats:   "{}",
		},
	}
}
