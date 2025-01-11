package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

// Struct for representing each entry
type Criterion struct {
	ID                    string `yaml:"id"`
	MaturityLevel         int    `yaml:"maturity_level"`
	Category              string `yaml:"category"`
	CriterionText         string `yaml:"criterion"`
	Rationale             string `yaml:"rationale"`
	Details               string `yaml:"details"`
	ControlMappings       string `yaml:"control_mappings"`
	SecurityInsightsValue string `yaml:"security_insights_value"`
}

// Struct for holding the entire YAML structure
type Baseline struct {
	Categories map[string]Category
	Lexicon    []LexiconEntry
}

type Category struct {
	CategoryName string      `yaml:"category"`
	Description  string      `yaml:"description"`
	Criteria     []Criterion `yaml:"criteria"`
}

type LexiconEntry struct {
	Term       string   `yaml:"term"`
	Definition string   `yaml:"definition"`
	Synonyms   []string `yaml:"synonyms"`
}

func hardcodedCategories() []string {
	return []string{
		"AC",
		"BR",
		"DO",
		"GV",
		"LE",
		"QA",
		"SA",
		"VM",
	}
}

func filename(name string) string {
	return filepath.Join(ContentDir, fmt.Sprintf("OSPS-%s.yaml", name))
}

func newBaseline() (Baseline, error) {
	lexicon, err := newLexicon()
	if err != nil {
		return Baseline{}, fmt.Errorf("error reading lexicon: %w", err)
	}
	b := Baseline{
		Categories: make(map[string]Category),
		Lexicon:    lexicon,
	}
	var failed bool
	for _, categoryName := range hardcodedCategories() {
		category, err := newCategory(categoryName)
		if err != nil {
			failed = true
			log.Printf("error reading category %s: %s", categoryName, err.Error())
		}
		b.Categories[categoryName] = category
	}
	if failed {
		return b, fmt.Errorf("error setting up baseline")
	}
	return b, b.validate()
}

func (b Baseline) validate() error {
	var entryIDs []string
	var failed bool
	for _, category := range b.Categories {
		for _, entry := range category.Criteria {
			if contains(entryIDs, entry.ID) {
				failed = true
				log.Printf("duplicate ID for 'criterion' for %s", entry.ID)
			}
			if entry.ID == "" {
				failed = true
				log.Printf("missing ID for 'criterion' %s", entry.ID)
			}
			if entry.CriterionText == "" {
				failed = true
				log.Printf("missing 'criterion' text for %s", entry.ID)
			}
			// For after all fields are populated:
			// if entry.Rationale == "" {
			// 	failed = true
			// 	log.Printf("missing 'rationale' for %s", entry.ID)
			// }
			// if entry.Details == "" {
			// 	failed = true
			// 	log.Printf("missing 'details' for %s", entry.ID)
			// }
			entryIDs = append(entryIDs, entry.ID)
		}
	}
	if failed {
		return fmt.Errorf("error validating baseline")
	}
	return nil
}

func contains(list []string, term string) bool {
	for _, item := range list {
		if item == term {
			return true
		}
	}
	return false
}

func newCategory(categoryName string) (Category, error) {
	file, err := os.Open(filename(categoryName))
	if err != nil {
		return Category{}, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var category Category

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&category); err != nil {
		return category, fmt.Errorf("error decoding YAML: %v", err)
	}
	return category, nil
}

func newLexicon() ([]LexiconEntry, error) {
	file, err := os.Open(LexiconPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var lexicon []LexiconEntry

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&lexicon); err != nil {
		return nil, fmt.Errorf("error decoding YAML: %v", err)
	}
	return lexicon, nil
}

func (b *Baseline) Generate() error {
	// Open or create the output file
	oDir := filepath.Dir(OutputPath)
	err := os.MkdirAll(oDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating output directory %s: %w", oDir, err)
	}
	outputFile, err := os.Create(OutputPath)
	if err != nil {
		return fmt.Errorf("error creating output file %s: %w", OutputPath, err)
	}
	defer outputFile.Close()

	// Read the markdown template from the external file
	templateContent, err := os.ReadFile(TemplatePath)
	if err != nil {
		return fmt.Errorf("error reading template file: %w", err)
	}

	// Create and parse the template
	tmpl, err := template.New("baseline").Funcs(template.FuncMap{
		// Template function to remove newlines and collapse text
		"collapseNewlines": func(s string) string {
			return strings.ReplaceAll(s, "\n", " ")
		},
		"addLinks": func(s string) string {
			return addLinksTemplateFunction(s)
		},
		"asLink": func(s string) string {
			return asLinkTemplateFunction(s)
		},
	}).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Execute the template and write to the output file
	err = tmpl.Execute(outputFile, b)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
