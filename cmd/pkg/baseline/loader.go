// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/ossf/gemara/layer2"
)

const (
	LexiconFilename    = "lexicon.yaml"
	FrameworksFilename = "frameworks.yaml"
)

// Loader is an object that reads the baseline data
type Loader struct {
	layer2.Catalog
	Lexicon []LexiconEntry
}

type LexiconEntry struct {
	Term       string   `yaml:"term"`
	Definition string   `yaml:"definition"`
	Synonyms   []string `yaml:"synonyms"`
	References []string `yaml:"references"`
}

func NewLoader() *Loader {
	return &Loader{}
}

// Load reads the baseline data and returns the representation types
func (l *Loader) Load(dataPath string) error {
	lexicon, err := loadLexicon(dataPath)
	if err != nil {
		return fmt.Errorf("error reading lexicon: %w", err)
	}
	l.Lexicon = lexicon

	dirFiles, err := os.ReadDir(dataPath)
	if err != nil {
		return fmt.Errorf("error reading data path: %w", err)
	}
	var filepaths []string
	for _, file := range dirFiles {
		if file.Name() != LexiconFilename {
			path := filepath.Join(dataPath, file.Name())
			filepaths = append(filepaths, path)
		}
	}

	err = l.LoadFiles(filepaths)
	if err != nil {
		return fmt.Errorf("Error loading baseline directory (%s)\n%w", dataPath, err)
	}
	return nil
}

// loadLexicon
func loadLexicon(dataPath string) ([]LexiconEntry, error) {
	file, err := os.Open(filepath.Join(dataPath, LexiconFilename))
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close() //nolint:errcheck

	var lexicon []LexiconEntry

	decoder := yaml.NewDecoder(file, yaml.DisallowUnknownField())
	if err := decoder.Decode(&lexicon); err != nil {
		return nil, fmt.Errorf("error decoding YAML: %s\n %w", file.Name(), err)
	}
	return lexicon, nil
}
