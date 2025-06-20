package config

import (
	"drawn-by-fate/pkg/logging"
	"drawn-by-fate/pkg/models/auto_types/game_config"
	"drawn-by-fate/pkg/ppaths"
	"drawn-by-fate/pkg/utils"
	"embed"
	"fmt"
	"path/filepath"
)

type GameConfig = game_config.GameConfig

//go:embed package/**
var embeddedPackage embed.FS

type CommonPaths struct {
	PkgRoot    string
	Assets     string
	Audio      string
	Fonts      string
	Images     string
	Tarot      string
	Configs    string
	GameConfig string
	EntityDefs string
	Schemas    string
	// TODO: Add more as they become necessary
}

func (p CommonPaths) LoadGameConfig() (*GameConfig, error) {

	data, err := utils.ReadDataFile[GameConfig](p.GameConfig)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p CommonPaths) LoadConfig(name string, data any) error {
	path := filepath.Join(p.Configs, name)
	if _, err := utils.ReadDataFile(path, data); err != nil {
		return err
	}

	return nil
}

func NewCommonPaths(name string) (cp CommonPaths, err error) {
	isMod := name != "core"
	appDir, err := ppaths.GetUserGameDir()
	if err != nil {
		// TODO: Better error handling/logging
		return cp, err
	}

	var rootDir string
	if isMod {
		rootDir = filepath.Join(appDir, "mods", name)
	} else {
		rootDir = filepath.Join(appDir, name)
	}
	assetsDir := filepath.Join(rootDir, "assets")

	cfgDir := filepath.Join(rootDir, "configs")
	return CommonPaths{
		PkgRoot:    rootDir,
		Assets:     assetsDir,
		Audio:      filepath.Join(assetsDir, "audio"),
		Fonts:      filepath.Join(assetsDir, "fonts"),
		Images:     filepath.Join(assetsDir, "images"),
		Tarot:      filepath.Join(assetsDir, "tarot"),
		Configs:    cfgDir,
		GameConfig: filepath.Join(cfgDir, "game_config.yaml"),
		EntityDefs: filepath.Join(cfgDir, "entity_defs.yaml"),
		Schemas:    filepath.Join(cfgDir, "schemas"),
	}, nil
}

var CorePaths CommonPaths

func NewGameConfig() *GameConfig {
	return &GameConfig{
		EnabledModIDS:    []string{},
		ScreenMode:       game_config.Window,
		ScreenResolution: game_config.HeightWidthPair{Height: 640, Width: 480},
	}
}

func GetScreenSize() (int, int) {
	return int(CurrentSettings.ScreenResolution.Width), int(CurrentSettings.ScreenResolution.Height)
}

var CurrentSettings = NewGameConfig()

func init() {

	paths, err := NewCommonPaths("core")
	if err != nil {
		// TODO: Improve error handling
		panic(err)
	}

	// Ensure the user data directory exists
	if _, err := ppaths.EnsureDataDir(); err != nil {
		panic(fmt.Errorf("failed to ensure user data directory: %w", err))
	}

	embeds, err := NewPackageDir("core")
	if err != nil {
		panic(err)
	}

	logging.LogDebug("embed", embeds.EmbedPath, "path", embeds.Path)
	if err := embeds.Embed(true); err != nil {
		panic(err)
	}

	CorePaths = paths

	// Load the game config
	if cfg, err := CorePaths.LoadGameConfig(); err != nil {
		panic(fmt.Errorf("failed to load game config: %w", err))
	} else {
		logging.LogDebug("screen res", cfg.ScreenResolution)
		// TODO: load mod paths
		CurrentSettings = cfg
	}
}
