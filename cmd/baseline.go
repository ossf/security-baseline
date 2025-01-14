package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

// Struct for representing each entry
type Criterion struct {
	ID                    string            `yaml:"id"`
	MaturityLevel         int               `yaml:"maturity_level"`
	Category              string            `yaml:"category"`
	CriterionText         string            `yaml:"criterion"`
	Rationale             string            `yaml:"rationale"`
	Implementation        string            `yaml:"implementation"`
	Details               string            `yaml:"details"`
	ControlMappings       map[string]string `yaml:"control_mappings"`
	SecurityInsightsValue string            `yaml:"security_insights_value"`
	// If ReplacedBy is set, no other fields (beyond ID) should be set
	ReplacedBy string `yaml:"replaced_by"`
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
	var errs []error
	for _, categoryName := range hardcodedCategories() {
		category, err := newCategory(categoryName)
		if err != nil {
			errs = append(errs, fmt.Errorf("error reading category %s: %w", categoryName, err))
		}
		b.Categories[categoryName] = category
	}
	categoryErr := errors.Join(errs...)
	if categoryErr != nil {
		return b, categoryErr
	}
	return b, b.validate()
}

func (b Baseline) validate() error {
	var entryIDs []string
	retiredIDs := map[string]string{}
	errs := []error{}
	for _, category := range b.Categories {
		for _, entry := range category.Criteria {
			if slices.Contains(entryIDs, entry.ID) {
				errs = append(errs, fmt.Errorf("duplicate ID for 'criterion' for %s", entry.ID))
			}
			entryIDs = append(entryIDs, entry.ID)
			if entry.ID == "" {
				errs = append(errs, fmt.Errorf("missing ID for 'criterion' %s", entry.ID))
			}
			if entry.ReplacedBy != "" {
				retiredIDs[entry.ID] = entry.ReplacedBy
				if !reflect.DeepEqual(entry, Criterion{ID: entry.ID, ReplacedBy: entry.ReplacedBy}) {
					errs = append(errs, fmt.Errorf("retired criterion %s has additional fields", entry.ID))
				}
				continue
			}
			if entry.CriterionText == "" {
				errs = append(errs, fmt.Errorf("missing 'criterion' text for %s", entry.ID))
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
		}
	}
	for retired, replacement := range retiredIDs {
		if !slices.Contains(entryIDs, replacement) {
			errs = append(errs, fmt.Errorf("retired criterion %s has invalid replacement %s", retired, replacement))
		}
		if _, ok := retiredIDs[replacement]; ok {
			errs = append(errs, fmt.Errorf("retired criterion %s references another retired criterion %s", retired, replacement))
		}
	}
	return errors.Join(errs...)
}

func newCategory(categoryName string) (Category, error) {
	file, err := os.Open(filename(categoryName))
	if err != nil {
		return Category{}, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var category Category

	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)
	if err := decoder.Decode(&category); err != nil {
		return category, fmt.Errorf("error decoding YAML: %v", err)
	}
	slices.SortFunc(category.Criteria, func(a, b Criterion) int {
		return strings.Compare(a.ID, b.ID)
	})
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
	decoder.KnownFields(true)
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
