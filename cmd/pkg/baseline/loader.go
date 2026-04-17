// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gemaraproj/go-gemara"
	"github.com/goccy/go-yaml"

	"github.com/ossf/security-baseline/pkg/types"
)

const (
	LexiconFilename  = "lexicon.yaml"
	MetadataFilename = "metadata.yaml"
)

// Loader reads baseline data from a directory of YAML source files
type Loader struct {
	DataPath string
}

func NewLoader() *Loader {
	return &Loader{}
}

// Load reads the baseline data and returns the representation types
func (l *Loader) Load() (*types.Baseline, error) {
	b := &types.Baseline{}

	// Load the lexicon:
	lexicon, err := l.loadLexicon()
	if err != nil {
		return nil, fmt.Errorf("error reading lexicon: %w", err)
	}
	b.Lexicon = lexicon

	catalog, err := l.loadMetadata()
	if err != nil {
		return nil, fmt.Errorf("error reading metadata: %w", err)
	}

	if err := l.loadControlFamilies(catalog); err != nil {
		return nil, fmt.Errorf("error reading families: %w", err)
	}

	b.Catalog = *catalog
	return b, nil
}

// decodeYAMLFile opens a YAML file and decodes it into target with strict
// field checking.
func decodeYAMLFile(path string, target any) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() //nolint:errcheck

	decoder := yaml.NewDecoder(file, yaml.DisallowUnknownField())
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("error decoding YAML: %w", err)
	}
	return nil
}

// loadLexicon decodes the controlled vocabulary used across the baseline.
func (l *Loader) loadLexicon() ([]types.LexiconEntry, error) {
	var lexicon []types.LexiconEntry
	if err := decodeYAMLFile(filepath.Join(l.DataPath, LexiconFilename), &lexicon); err != nil {
		return nil, err
	}
	return lexicon, nil
}

// loadMetadata decodes the catalog-level metadata.
func (l *Loader) loadMetadata() (*gemara.ControlCatalog, error) {
	var catalog gemara.ControlCatalog
	if err := decodeYAMLFile(filepath.Join(l.DataPath, MetadataFilename), &catalog); err != nil {
		return nil, err
	}
	return &catalog, nil
}

// loadControlFamilies decodes per-family YAML files and appends their groups
// and controls onto the provided catalog.
func (l *Loader) loadControlFamilies(catalog *gemara.ControlCatalog) error {
	absData, err := filepath.Abs(l.DataPath)
	if err != nil {
		return fmt.Errorf("resolving absolute path: %w", err)
	}

	familyPaths := make([]string, 0, len(types.ControlFamilies))
	for _, familyID := range types.ControlFamilies {
		familyPath := "file://" + filepath.Join(absData, fmt.Sprintf("OSPS-%s.yaml", familyID))
		familyPaths = append(familyPaths, familyPath)
	}

	if err := catalog.LoadFiles(familyPaths); err != nil {
		return fmt.Errorf("error loading controls families: %w", err)
	}

	return nil
}
