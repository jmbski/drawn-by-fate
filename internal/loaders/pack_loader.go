package loaders

import (
	"drawn-by-fate/pkg/config"
	"drawn-by-fate/pkg/models/auto_types/entity_defs"
)

type PackData struct {
	Name          string
	ID            string
	AppVersion    string
	SchemaVersion string
	ModVersion    string
	EntityDefs    *entity_defs.EntityDefs
	LoadOrder     int64
	Paths         *config.CommonPaths
	Dependencies  []string
}

func NewGameMod(modId string) PackData {
	return PackData{}
}

type ModLoader struct {
	Mods []PackData
	Core PackData
}

func NewModLoader() *ModLoader {

	return &ModLoader{}
}
