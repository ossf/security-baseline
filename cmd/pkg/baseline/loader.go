// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

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
	b := &types.Baseline{
		Categories: make(map[string]types.Category, len(types.Categories)),
	}

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

	for _, catCode := range types.Categories {
		cat, err := l.loadCategory(catCode)
		if err != nil {
			return nil, fmt.Errorf("loading category %q: %w", catCode, err)
		}
		b.Categories[catCode] = *cat
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

	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)
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

	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)
	if err := decoder.Decode(&frameworks); err != nil {
		return nil, fmt.Errorf("error decoding YAML: %w", err)
	}
	return frameworks.Frameworks, nil
}

// loadCategory loads a category definition from its YAML source
func (l *Loader) loadCategory(catCode string) (*types.Category, error) {
	file, err := os.Open(filepath.Join(l.DataPath, fmt.Sprintf("OSPS-%s.yaml", catCode)))
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() //nolint:errcheck

	category := &types.Category{}

	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)
	if err := decoder.Decode(category); err != nil {
		return nil, fmt.Errorf("error decoding %s YAML: %w", catCode, err)
	}
	return category, nil
}
