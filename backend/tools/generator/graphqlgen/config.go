package main

import (
	"os"
	"path"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"go.yaml.in/yaml/v3"
)

type Config struct {
	Path           string             `validate:"required" yaml:"path"`
	PackageName    string             `validate:"required" yaml:"packageName"`
	NeededPackages *NeededPackagesCfg `validate:"required" yaml:"neededPackages"`
	Connections    []*ConnectionCfg   `validate:"required" yaml:"connections"`
}

type NeededPackagesCfg struct {
	GqlgenModelPackage string `validate:"required" yaml:"gqlgenModelPackage"`
	Pagination         string `validate:"required" yaml:"pagination"`
	GraphqlUtils       string `validate:"required" yaml:"graphqlUtils"`
}

type ConnectionCfg struct {
	Package       string `validate:"required" yaml:"package"`
	StructureName string `validate:"required" yaml:"structureName"`
}

func loadConfig() (*Config, error) {
	// Get cwd
	cwd, err := os.Getwd()
	// Check error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Build configuration path
	cfgPath := path.Join(cwd, ".graphqlgen.yml")

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
