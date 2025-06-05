// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package types

import "github.com/revanite-io/sci/layer2"

var ControlFamilies = []string{
	"AC",
	"BR",
	"DO",
	"GV",
	"LE",
	"QA",
	"SA",
	"VM",
}

// Struct for holding the entire YAML structure
type Baseline struct {
	// map of family names to IDs to support OSCAL groups
	ControlFamilyIDs map[string]string
	Catalog          layer2.Catalog
	Lexicon          []LexiconEntry
}

type LexiconEntry struct {
	Term       string   `yaml:"term"`
	Definition string   `yaml:"definition"`
	Synonyms   []string `yaml:"synonyms"`
	References []string `yaml:"references"`
}
