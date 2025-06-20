package loaders

import "drawn-by-fate/pkg/config"

type AssetLoader struct {
	Paths    config.CommonPaths
	ModPaths map[string]*config.CommonPaths
}

func NewAssetLoader() *AssetLoader {
	return &AssetLoader{
		Paths: config.CorePaths,
	}
}
