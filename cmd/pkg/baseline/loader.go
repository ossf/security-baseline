// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/revanite-io/sci/layer2"

	"github.com/ossf/security-baseline/pkg/types"
)

const (
	LexiconFilename    = "lexicon.yaml"
	FrameworksFilename = "frameworks.yaml"
)

// Loader is an object that reads the baseline data
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
	frameworks, err := l.loadFramework()
	if err != nil {
		return nil, fmt.Errorf("error reading frameworks: %w", err)
	}

	b.Lexicon = lexicon
	b.Frameworks = frameworks
	controlFamilies := []layer2.ControlFamily{}
	familyIDs := map[string]string{}

	for _, familyID := range types.ControlFamilies {
		cf, err := l.loadControlFamily(familyID)
		if err != nil {
			return nil, fmt.Errorf("loading control family %s: %w", familyID, err)
		}
		familyIDs[cf.Title] = familyID
		controlFamilies = append(controlFamilies, *cf)
	}
	b.ControlFamilyIDs = familyIDs
	b.Catalog = layer2.Layer2{
		ControlFamilies: controlFamilies,
	}
	return b, nil
}

// loadLexicon
func (l *Loader) loadLexicon() ([]types.LexiconEntry, error) {
	file, err := os.Open(filepath.Join(l.DataPath, LexiconFilename))
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() //nolint:errcheck

	var lexicon []types.LexiconEntry

	decoder := yaml.NewDecoder(file, yaml.DisallowUnknownField())
	if err := decoder.Decode(&lexicon); err != nil {
		return nil, fmt.Errorf("error decoding YAML: %w", err)
	}
	return lexicon, nil
}

// loadLexicon
func (l *Loader) loadFramework() ([]types.FrameworkEntry, error) {
	file, err := os.Open(filepath.Join(l.DataPath, FrameworksFilename))
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() //nolint:errcheck

	var frameworks types.Frameworks

	decoder := yaml.NewDecoder(file, yaml.DisallowUnknownField())
	if err := decoder.Decode(&frameworks); err != nil {
		return nil, fmt.Errorf("error decoding YAML: %w", err)
	}
	return frameworks.Frameworks, nil
}

// loadControlFamily loads a ControlFamily definition from its YAML source
func (l *Loader) loadControlFamily(familyID string) (*layer2.ControlFamily, error) {
	file, err := os.Open(filepath.Join(l.DataPath, fmt.Sprintf("OSPS-%s.yaml", familyID)))
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() //nolint:errcheck

	controlFamily := &layer2.ControlFamily{}

	decoder := yaml.NewDecoder(file, yaml.Strict())
	if err := decoder.Decode(controlFamily); err != nil {
		return nil, fmt.Errorf("error decoding %s YAML: %w", familyID, err)
	}
	return controlFamily, nil
}
