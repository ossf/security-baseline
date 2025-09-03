// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

// ExportMarkdown runs the baseline data through the markdown template
func (l *Loader) ExportMarkdown(templatePath, outPath string) error {
	// Read the markdown template from the external file
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading markdown template: %w", err)
	}

	// Open or create the output file
	if err := os.MkdirAll(filepath.Dir(outPath), os.FileMode(0o750)); err != nil {
		return fmt.Errorf("error creating output directory %s: %w", filepath.Dir(outPath), err)
	}

	outputFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("error creating output file %s: %w", outPath, err)
	}
	defer outputFile.Close() //nolint:errcheck

	// Create and parse the template
	tmpl, err := template.New("baseline").Funcs(template.FuncMap{
		// Template function to remove newlines and collapse text
		"collapseNewlines": collapseNewlines,
		"addLinks": func(s string) string {
			return addLinksTemplateFunction(l.Lexicon, s)
		},
		"asLink":   asLinkTemplateFunction,
		"maxLevel": maxLevel,
	}).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Execute the template and write to the output file
	if err := tmpl.Execute(outputFile, l); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
