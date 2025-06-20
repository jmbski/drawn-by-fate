package models

import (
	"drawn-by-fate/pkg/config"
	"drawn-by-fate/pkg/models/auto_types/entity_defs"
	"drawn-by-fate/pkg/utils"
	"path/filepath"
)

func LoadEntityTemplates(modName string) (*entity_defs.EntityDefs, error) {

	modPaths, err := config.NewCommonPaths(modName, modName != "core")
	if err != nil {
		return nil, err
	}

	path := filepath.Join(modPaths.Configs, "entity_defs.json")

	templates := &entity_defs.EntityDefs{}

	if _, err := utils.ReadJSON(path, templates); err != nil {
		return nil, err
	}

	return templates, nil
}
