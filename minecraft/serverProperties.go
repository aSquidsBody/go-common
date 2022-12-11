package minecraft

import (
	"fmt"
	"strings"

	"github.com/aSquidsBody/go-common/logs"
	"gopkg.in/yaml.v2"
)

// ServerProperties contains the data for a server.properties file
type ServerProperties struct {
	EnableJmxMonitoring            bool   `yaml:"enable-jmx-monitoring" json:"enable-jmx-monitoring,omitempty"`
	LevelSeed                      string `yaml:"level-seed" json:"level-seed,omitempty"`
	RconPort                       int    `yaml:"rcon.port" json:"rcon-port,omitempty"`
	EnableCommandBlock             bool   `yaml:"enable-command-block" json:"enable-command-block,omitempty"`
	Gamemode                       string `yaml:"gamemode" json:"gamemode,omitempty"`
	EnableQuery                    bool   `yaml:"enable-query" json:"enable-query,omitempty"`
	GeneratorSettings              string `yaml:"generator-settings" json:"generator-settings,omitempty"`
	EnforceSecureProfile           bool   `yaml:"enforce-secure-profile" json:"enforce-secure-profile,omitempty"`
	LevelName                      string `yaml:"level-name" json:"level-name,omitempty"`
	Motd                           string `yaml:"motd" json:"motd,omitempty"`
	QueryPort                      int    `yaml:"query.port" json:"query-port,omitempty"`
	Pvp                            bool   `yaml:"pvp" json:"pvp,omitempty"`
	GenerateStructures             bool   `yaml:"generate-structures" json:"generate-structures,omitempty"`
	MaxChainedNeighborUpdates      int    `yaml:"max-chained-neighbor-updates" json:"max-chained-neighbor-updates,omitempty"`
	Difficulty                     string `yaml:"difficulty" json:"difficulty,omitempty"`
	NetworkCompressionThreshold    int    `yaml:"network-compression-threshold" json:"network-compression-threshold,omitempty"`
	MaxTickTime                    int    `yaml:"max-tick-time" json:"max-tick-time,omitempty"`
	RequireResourcePack            bool   `yaml:"require-resource-pack" json:"require-resource-pack,omitempty"`
	MaxPlayers                     int    `yaml:"max-players" json:"max-players,omitempty"`
	UseNativeTransport             bool   `yaml:"use-native-transport" json:"use-native-transport,omitempty"`
	OnlineMode                     bool   `yaml:"online-mode" json:"online-mode,omitempty"`
	EnableStatus                   bool   `yaml:"enable-status" json:"enable-status,omitempty"`
	AllowFlight                    bool   `yaml:"allow-flight" json:"allow-flight,omitempty"`
	BroadcastRconToOps             bool   `yaml:"broadcast-rcon-to-ops" json:"broadcast-rcon-to-ops,omitempty"`
	ViewDistance                   int    `yaml:"view-distance" json:"view-distance,omitempty"`
	MaxBuildHeight                 int    `yaml:"max-build-height" json:"max-build-height,omitempty"`
	ServerIp                       string `yaml:"server-ip" json:"server-ip,omitempty"`
	ResourcePackPrompt             string `yaml:"resource-pack-prompt" json:"resource-pack-prompt,omitempty"`
	AllowNether                    bool   `yaml:"allow-nether" json:"allow-nether,omitempty"`
	ServerPort                     int    `yaml:"server-port" json:"server-port,omitempty"`
	EnableRcon                     bool   `yaml:"enable-rcon" json:"enable-rcon,omitempty"`
	SyncChunkWrites                bool   `yaml:"sync-chunk-writes" json:"sync-chunk-writes,omitempty"`
	OpPermissionLevel              int    `yaml:"op-permission-level" json:"op-permission-level,omitempty"`
	PreventProxyConnections        bool   `yaml:"prevent-proxy-connections" json:"prevent-proxy-connections,omitempty"`
	HideOnlinePlayers              bool   `yaml:"hide-online-players" json:"hide-online-players,omitempty"`
	ResourcePack                   string `yaml:"resource-pack" json:"resource-pack,omitempty"`
	TexturePack                    string `yaml:"texture-pack" json:"texture-pack,omitempty"`
	EntityBroadcastRangePercentage int    `yaml:"entity-broadcast-range-percentage" json:"entity-broadcast-range-percentage,omitempty"`
	SimulationDistance             int    `yaml:"simulation-distance" json:"simulation-distance,omitempty"`
	PlayerIdleTimeout              int    `yaml:"player-idle-timeout" json:"player-idle-timeout,omitempty"`
	RconPassword                   string `yaml:"rcon.password" json:"rcon-password,omitempty"`
	ForceGamemode                  bool   `yaml:"force-gamemode" json:"force-gamemode,omitempty"`
	Debug                          bool   `yaml:"debug" json:"debug,omitempty"`
	RateLimit                      int    `yaml:"rate-limit" json:"rate-limit,omitempty"`
	Hardcore                       bool   `yaml:"hardcore" json:"hardcore,omitempty"`
	WhiteList                      bool   `yaml:"white-list" json:"white-list,omitempty"`
	BroadcastConsoleToOps          bool   `yaml:"broadcast-console-to-ops" json:"broadcast-console-to-ops,omitempty"`
	SpawnNpcs                      bool   `yaml:"spawn-npcs" json:"spawn-npcs,omitempty"`
	PreviewsChat                   bool   `yaml:"previews-chat" json:"previews-chat,omitempty"`
	SpawnAnimals                   bool   `yaml:"spawn-animals" json:"spawn-animals,omitempty"`
	SnooperEnabled                 bool   `yaml:"snooper-enabled" json:"snooper-enabled,omitempty"`
	FunctionPermissionLevel        int    `yaml:"function-permission-level" json:"function-permission-level,omitempty"`
	LevelType                      string `yaml:"level-type" json:"level-type,omitempty"`
	TextFilteringConfig            string `yaml:"text-filtering-config" json:"text-filtering-config,omitempty"`
	SpawnMonsters                  bool   `yaml:"spawn-monsters" json:"spawn-monsters,omitempty"`
	EnforceWhitelist               bool   `yaml:"enforce-whitelist" json:"enforce-whitelist,omitempty"`
	ResourcePackSha1               string `yaml:"resource-pack-sha1" json:"resource-pack-sha1,omitempty"`
	SpawnProtection                int    `yaml:"spawn-protection" json:"spawn-protection,omitempty"`
	MaxWorldSize                   int    `yaml:"max-world-size" json:"max-world-size,omitempty"`
}

func castBool(key string, value interface{}) (bool, error) {
	if v, ok := value.(bool); !ok {
		err := fmt.Errorf("Invalid value: %s", value)
		logs.Errorf(err, "Failed to set %s in server.properties", key)
		return false, err
	} else {
		return v, nil
	}
}

func castString(key string, value interface{}) (string, error) {
	if v, ok := value.(string); !ok {
		err := fmt.Errorf("Invalid value: %s", value)
		logs.Errorf(err, "Failed to set %s in server.properties", key)
		return "", err
	} else {
		return v, nil
	}
}

func castInt(key string, value interface{}) (int, error) {
	if v, ok := value.(int); !ok {
		err := fmt.Errorf("Invalid value: %s", value)
		logs.Errorf(err, "Failed to set %s in server.properties", key)
		return 0, err
	} else {
		return v, nil
	}
}

// Set the value of a property in server.properties
func (s *ServerProperties) Set(key string, value interface{}) error {
	switch key {
	case "enable-jmx-monitoring":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.EnableJmxMonitoring = v

	case "level-seed":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.LevelSeed = v

	case "rcon.port":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.RconPort = v

	case "enable-command-block":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.EnableCommandBlock = v
	case "gamemode":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.Gamemode = v
	case "enable-query":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.EnableQuery = v
	case "generator-settings":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.GeneratorSettings = v
	case "enforce-secure-profile":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.EnforceSecureProfile = v
	case "level-name":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.LevelName = v
	case "motd":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.Motd = v
		fmt.Println("IN SET", s.Motd)
	case "query.port":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.QueryPort = v
	case "texture-pack":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.TexturePack = v
	case "pvp":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.Pvp = v
	case "generate-structures":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.GenerateStructures = v
	case "max-chained-neighbor-updates":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.MaxChainedNeighborUpdates = v
	case "difficulty":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.Difficulty = v
	case "network-compression-threshold":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.NetworkCompressionThreshold = v
	case "max-tick-time":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.MaxTickTime = v
	case "require-resource-pack":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.RequireResourcePack = v
	case "max-players":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.MaxPlayers = v
	case "use-native-transport":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.UseNativeTransport = v
	case "online-mode":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.OnlineMode = v
	case "enable-status":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.EnableStatus = v
	case "allow-flight":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.AllowFlight = v
	case "broadcast-rcon-to-ops":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.BroadcastRconToOps = v
	case "view-distance":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.ViewDistance = v
	case "max-build-height":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.MaxBuildHeight = v
	case "server-ip":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.ServerIp = v
	case "resource-pack-prompt":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.ResourcePackPrompt = v
	case "allow-nether":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.AllowNether = v
	case "server-port":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.ServerPort = v
	case "enable-rcon":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.EnableRcon = v
	case "sync-chunk-writes":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.SyncChunkWrites = v
	case "op-permission-level":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.OpPermissionLevel = v
	case "prevent-proxy-connections":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.PreventProxyConnections = v
	case "hide-online-players":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.HideOnlinePlayers = v
	case "resource-pack":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.ResourcePack = v
	case "entity-broadcast-range-percentage":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.EntityBroadcastRangePercentage = v

	case "simulation-distance":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.SimulationDistance = v
	case "player-idle-timeout":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.PlayerIdleTimeout = v
	case "rcon.password":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.RconPassword = v
	case "force-gamemode":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.ForceGamemode = v
	case "debug":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.Debug = v
	case "rate-limit":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.RateLimit = v
	case "hardcore":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.Hardcore = v
	case "white-list":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.WhiteList = v
	case "broadcast-console-to-ops":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.BroadcastConsoleToOps = v
	case "spawn-npcs":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.SpawnNpcs = v
	case "previews-chat":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.PreviewsChat = v
	case "spawn-animals":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.SpawnAnimals = v
	case "snooper-enabled":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.SnooperEnabled = v

	case "function-permission-level":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.FunctionPermissionLevel = v
	case "level-type":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.LevelType = v
	case "text-filtering-config":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.TextFilteringConfig = v
	case "spawn-monsters":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.SpawnMonsters = v
	case "enforce-whitelist":
		v, err := castBool(key, value)
		if err != nil {
			return err
		}
		s.EnforceWhitelist = v
	case "resource-pack-sha1":
		v, err := castString(key, value)
		if err != nil {
			return err
		}
		s.ResourcePackSha1 = v
	case "spawn-protection":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.SpawnProtection = v
	case "max-world-size":
		v, err := castInt(key, value)
		if err != nil {
			return err
		}
		s.MaxWorldSize = v
	default:
		err := fmt.Errorf("Invalid key: %s", key)
		logs.Error(err, "Could not set server property")
		return err
	}
	return nil
}

func (s ServerProperties) ToYaml() (string, error) {
	b, err := yaml.Marshal(s)
	if err != nil {
		fmt.Println("Failed to send server.properties toString()")
		return "", err
	}
	str := strings.ReplaceAll(string(b), ": ", "=")
	str = strings.ReplaceAll(str, "'", "")
	str = strings.ReplaceAll(str, "\"", "")

	return str, nil
}

// Creates a new ServerProperties object and populates it with the default values
func NewServerProperties() ServerProperties {
	return ServerProperties{
		EnableJmxMonitoring:            false,
		LevelSeed:                      "",
		RconPort:                       25575,
		EnableCommandBlock:             true,
		Gamemode:                       "survival",
		EnableQuery:                    false,
		GeneratorSettings:              "{}",
		EnforceSecureProfile:           true,
		LevelName:                      "",
		Motd:                           "A Spigot Server powered by Docker",
		QueryPort:                      25565,
		TexturePack:                    "",
		Pvp:                            true,
		GenerateStructures:             true,
		MaxChainedNeighborUpdates:      1_000_000,
		Difficulty:                     "normal",
		NetworkCompressionThreshold:    256,
		MaxTickTime:                    60_000,
		RequireResourcePack:            false,
		MaxPlayers:                     10,
		UseNativeTransport:             true,
		OnlineMode:                     false,
		EnableStatus:                   true,
		AllowFlight:                    false,
		BroadcastRconToOps:             true,
		ViewDistance:                   10,
		MaxBuildHeight:                 256,
		ServerIp:                       "",
		ResourcePackPrompt:             "",
		AllowNether:                    true,
		ServerPort:                     25565,
		EnableRcon:                     true,
		SyncChunkWrites:                true,
		OpPermissionLevel:              4,
		PreventProxyConnections:        false,
		HideOnlinePlayers:              false,
		ResourcePack:                   "",
		EntityBroadcastRangePercentage: 100,
		SimulationDistance:             10,
		PlayerIdleTimeout:              0,
		RconPassword:                   "minecraft",
		ForceGamemode:                  false,
		Debug:                          false,
		RateLimit:                      0,
		Hardcore:                       false,
		WhiteList:                      false,
		BroadcastConsoleToOps:          true,
		SpawnNpcs:                      true,
		PreviewsChat:                   false,
		SpawnAnimals:                   true,
		SnooperEnabled:                 true,
		FunctionPermissionLevel:        2,
		LevelType:                      "default",
		TextFilteringConfig:            "",
		SpawnMonsters:                  true,
		EnforceWhitelist:               false,
		ResourcePackSha1:               "",
		SpawnProtection:                16,
		MaxWorldSize:                   29_999_984,
	}
}

func oneOf(v string, arr []string) bool {
	for _, elem := range arr {
		if v == elem {
			return true
		}
	}

	return false
}
