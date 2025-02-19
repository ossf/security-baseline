// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package types

var Categories = []string{
	"AC",
	"BR",
	"DO",
	"GV",
	"LE",
	"QA",
	"SA",
	"VM",
}

// Struct for representing each entry
type Control struct {
	ID           string        `yaml:"id"`
	Title        string        `yaml:"title"`
	Objective    string        `yaml:"objective"`
	Family       string        `yaml:"family"`
	Mappings     []Mapping     `yaml:"mappings"`
	Requirements []Requirement `yaml:"assessment-requirements"`
}

type Mapping struct {
	ReferenceID string   `yaml:"reference-id"`
	Identifiers []string `yaml:"identifiers"`
}

type Requirement struct {
	ID             string   `yaml:"id"`
	Text           string   `yaml:"text"`
	Applicability  []string `yaml:"applicability"`
	Recommendation string   `yaml:"recommendation"`
}

// Struct for holding the entire YAML structure
type Baseline struct {
	Categories map[string]Category
	Lexicon    []LexiconEntry
	Frameworks []FrameworkEntry `yaml:"mapping-references"`
}

type Category struct {
	CategoryName string    `yaml:"category"`
	Description  string    `yaml:"description"`
	Controls     []Control `yaml:"controls"`
}

type LexiconEntry struct {
	Term       string   `yaml:"term"`
	Definition string   `yaml:"definition"`
	Synonyms   []string `yaml:"synonyms"`
	References []string `yaml:"references"`
}

type FrameworkEntry struct {
	ID      string `yaml:"id"`
	Title   string `yaml:"title"`
	Version string `yaml:"version"`
	URL     string `yaml:"url"`
}
