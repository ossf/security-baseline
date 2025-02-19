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
	ID                    string            `yaml:"id"`
	MaturityLevel         int               `yaml:"maturity_level"`
	Category              string            `yaml:"category"`
	ControlText           string            `yaml:"control"`
	Rationale             string            `yaml:"rationale"`
	Implementation        string            `yaml:"implementation"`
	Details               string            `yaml:"details"`
	ControlMappings       map[string]string `yaml:"control_mappings"`
	SecurityInsightsValue string            `yaml:"security_insights_value"`
}

// Struct for holding the entire YAML structure
type Baseline struct {
	Categories map[string]Category
	Lexicon    []LexiconEntry
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
