package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// Struct for representing each criteria entry
type Criteria struct {
	ID                    string `yaml:"id"`
	MaturityLevel         int    `yaml:"maturity_level"`
	Category              string `yaml:"category"`
	CriteriaText          string `yaml:"criteria"`
	Objective             string `yaml:"objective"`
	Implementation        string `yaml:"implementation"`
	ControlMappings       string `yaml:"control_mappings"`
	SecurityInsightsValue string `yaml:"security_insights_value"`
	ScorecardProbe        string `yaml:"scorecard_probe"`
}

// Struct for holding the entire YAML structure
type Baseline struct {
	Criteria []Criteria     `yaml:"criteria"`
	Lexicon  []LexiconEntry `yaml:"lexicon"`
}

type LexiconEntry struct {
	Term       string   `yaml:"term"`
	Definition string   `yaml:"definition"`
	Synonyms   []string `yaml:"synonyms"`
}

// Root command for Cobra
var rootCmd = &cobra.Command{
	Use:   "baseline-compiler",
	Short: "Baseline Compiler is a tool for compiling YAML-based security criteria.",
	Long:  `Baseline Compiler reads a YAML file that defines security criteria and outputs a structured report.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Print help when no arguments are passed
		cmd.Help()
	},
}

// Command to process YAML file
var compileCmd = &cobra.Command{
	Use:   "compile [file]",
	Short: "Compile a YAML file of security criteria",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		baseline, err := readYAMLFile(filePath)
		if err != nil {
			log.Fatalf("Error reading YAML file: %v", err)
		}
		fmt.Println(baseline)
		err = generateBaselineMdFile(baseline)
		if err != nil {
			log.Fatalf("Error generating output: %v", err)
		}

	},
}

func main() {
	// Add compile command to root command
	rootCmd.AddCommand(compileCmd)

	// Execute Cobra root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Function to read the YAML file
func readYAMLFile(filePath string) (Baseline, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Baseline{}, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var baseline Baseline
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&baseline); err != nil {
		return Baseline{}, fmt.Errorf("error decoding YAML: %v", err)
	}

	return baseline, nil
}

// Function to generate the markdown file
func generateBaselineMdFile(baseline Baseline) (err error) {
	templatePath := "template.md"
	outputDir := ".."
	mdFileName := "baseline.md"
	outputPath := filepath.Join(outputDir, mdFileName)

	// Open or create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file %s: %w", outputPath, err)
	}
	defer outputFile.Close()

	// Read the markdown template from the external file
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("error reading template file: %w", err)
	}

	// Create and parse the template
	tmpl, err := template.New("baseline").Funcs(template.FuncMap{
		// Template function to remove newlines and collapse text
		"collapseNewlines": func(s string) string {
			return strings.ReplaceAll(s, "\n", " ")
		},
	}).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Execute the template and write to the output file
	err = tmpl.Execute(outputFile, baseline)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
