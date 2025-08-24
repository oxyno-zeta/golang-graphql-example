package main

import (
	"os"
	"path"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert/yaml"
)

type Config struct {
	NeededPackages *NeededPackagesCfg `validate:"required" yaml:"neededPackages"`
	Daos           []*DaoCfg          `validate:"required" yaml:"daos"`
}

type NeededPackagesCfg struct {
	DB         string `validate:"required" yaml:"db"`
	Pagination string `validate:"required" yaml:"pagination"`
	Helpers    string `validate:"required" yaml:"helpers"`
}

type DaoCfg struct {
	Path          string         `validate:"required" yaml:"path"`
	PackageName   string         `validate:"required" yaml:"packageName"`
	InterfaceName string         `                    yaml:"interfaceName"`
	Models        []*DaoModelCfg `                    yaml:"models"        validated:"required"`
}

type DaoModelCfg struct {
	Package                 string `validate:"required" yaml:"package"`
	StructureName           string `validate:"required" yaml:"structureName"`
	ProjectionStructureName string `                    yaml:"projectionStructureName"`
	SortOrderStructureName  string `                    yaml:"sortOrderStructureName"`
	FilterStructureName     string `                    yaml:"filterStructureName"`
}

func loadConfig() (*Config, error) {
	// Get cwd
	cwd, err := os.Getwd()
	// Check error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Build configuration path
	cfgPath := path.Join(cwd, "daogen.yml")

	// Read file
	yamlFile, err := os.ReadFile(cfgPath)
	// Check error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var config Config

	// Parse
	err = yaml.Unmarshal(yamlFile, &config)
	// Check error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Validate
	err = validator.New().Struct(&config)
	// Check error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Default
	return &config, nil
}
