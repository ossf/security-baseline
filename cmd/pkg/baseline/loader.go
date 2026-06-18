// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gemaraproj/go-gemara"
	"github.com/gemaraproj/go-gemara/fetcher"
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

// toFileURI converts local paths to RFC-8089 file:/// URIs, handling Windows drive-letter paths correctly.
func toFileURI(path string) string {
	switch {
	case strings.HasPrefix(path, "https://"), strings.HasPrefix(path, "http://"):
		return path
	case strings.HasPrefix(path, "file:///"):
		return path
	case strings.HasPrefix(path, "file://"):
		path = strings.TrimPrefix(path, "file://")
	}
	cleaned := filepath.Clean(path)
	cleaned = strings.ReplaceAll(cleaned, "\\", "/")
	isAbs := strings.HasPrefix(cleaned, "/") ||
		filepath.IsAbs(path) ||
		(len(cleaned) >= 3 && cleaned[1] == ':' && cleaned[2] == '/')
	if isAbs {
		return "file:///" + strings.TrimPrefix(cleaned, "/")
	}
	return cleaned
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
		familyPath := toFileURI(filepath.Join(absData, fmt.Sprintf("OSPS-%s.yaml", familyID)))
		familyPaths = append(familyPaths, familyPath)
	}

	if err := catalog.LoadFiles(context.Background(), &fetcher.URI{}, familyPaths); err != nil {
		return fmt.Errorf("error loading controls families: %w", err)
	}

	return nil
}
