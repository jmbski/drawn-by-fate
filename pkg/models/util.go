package models

import (
	"drawn-by-fate/pkg/config"
	"drawn-by-fate/pkg/models/auto_types/entity_defs"
	"drawn-by-fate/pkg/models/auto_types/pack_data"
	"drawn-by-fate/pkg/utils"
	"path/filepath"
)

func LoadEntityTemplates(modName string) (*entity_defs.EntityDefs, error) {

	modPaths, err := config.NewCommonPaths(modName)
	if err != nil {
		return nil, err
	}

	path := filepath.Join(modPaths.Configs, "entity_defs.yaml")

	templates := &entity_defs.EntityDefs{}

	if _, err := utils.ReadDataFile(path, templates); err != nil {
		return nil, err
	}

	return templates, nil
}

func LoadPackData(packName string) (*pack_data.PackData, error) {
	packPaths, err := config.NewCommonPaths(packName)
	if err != nil {
		return nil, err
	}

	packData := &pack_data.PackData{}

	if err = packPaths.LoadConfig("pack_data.yaml", packData); err != nil {
		return nil, err
	}

	return packData, nil
}
