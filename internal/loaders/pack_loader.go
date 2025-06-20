package loaders

import (
	"drawn-by-fate/pkg/config"
	"drawn-by-fate/pkg/logging"
	"drawn-by-fate/pkg/models"
	"drawn-by-fate/pkg/models/auto_types/pack_data"
)

type PackData = pack_data.PackData

type AssetMaps struct {
	Audio  map[string]string
	Fonts  map[string]string
	Images map[string]string
	Tarot  map[string]string
	// TODO: Add more sections as more assets are implemented
}

type PackLoader struct {
	ModPacks  []PackData
	CorePack  PackData
	AssetMaps *AssetMaps
}

func (p *PackLoader) OrderPacks() error {
	
	return nil
}

func NewPackLoader() (*PackLoader, error) {

	activePacks := config.CurrentSettings.EnabledModIDS
	modPacks := make([]PackData, 0, len(activePacks))

	for _, pack := range activePacks {
		packData, err := models.LoadPackData(pack)
		if err != nil {
			logging.LogWarning(err)
			continue
		}

		modPacks = append(modPacks, *packData)

	}

	core, err := models.LoadPackData("core")
	if err != nil {
		return nil, err
	}

	loader := &PackLoader{
		ModPacks:  modPacks,
		CorePack:  *core,
		AssetMaps: &AssetMaps{},
	}

	return loader, nil
}
