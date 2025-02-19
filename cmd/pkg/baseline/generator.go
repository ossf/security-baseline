// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/ossf/security-baseline/pkg/types"
)

func NewGenerator() *Generator {
	return &Generator{}
}

type Generator struct {
	TemplatePath string
}

// ExportMarkdown runs the baseline data through the markdown template
func (g *Generator) ExportMarkdown(b *types.Baseline, path string) error {
	// Read the markdown template from the external file
	templateContent, err := os.ReadFile(g.TemplatePath)
	if err != nil {
		return fmt.Errorf("error reading markdown template: %w", err)
	}

	// Open or create the output file
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return fmt.Errorf("error creating output directory %s: %w", filepath.Dir(path), err)
	}

	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating output file %s: %w", path, err)
	}
	defer outputFile.Close()

	// Create and parse the template
	tmpl, err := template.New("baseline").Funcs(template.FuncMap{
		// Template function to remove newlines and collapse text
		"collapseNewlines": collapseNewlines,
		"addLinks": func(s string) string {
			return addLinksTemplateFunction(b.Lexicon, s)
		},
		"asLink":   asLinkTemplateFunction,
		"maxLevel": maxLevel,
	}).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Execute the template and write to the output file
	if err := tmpl.Execute(outputFile, b); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
