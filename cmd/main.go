package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

var Data Baseline

// Command to process YAML file
var compileCmd = &cobra.Command{
	Use:   "compile [file]",
	Short: "Compile a YAML file of security criteria",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		err := readYAMLFile(filePath)
		if err != nil {
			log.Fatalf("Error reading YAML file: %v", err)
		}
		err = generateBaselineMdFile(Data)
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
func readYAMLFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var baseline Baseline
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&baseline); err != nil {
		return fmt.Errorf("error decoding YAML: %v", err)
	}
	Data = baseline
	return nil
}

func containsSynonym(list []string, entryTerm, term string) bool {
	if strings.EqualFold(entryTerm, term) {
		return true
	}
	for _, item := range list {
		if strings.EqualFold(item, term) {
			return true
		}
	}
	return false
}

// Function to add links by wrapping terms with square brackets
func addLinks(text, term string) string {
	// Escape any special characters in the term to use in regex
	escapedTerm := regexp.QuoteMeta(term)

	// Create a regular expression to match the term as a whole word
	// The `(?i)` part makes it case-insensitive, and `\b` ensures whole word matching
	termRegex := regexp.MustCompile(`(?i)\b` + escapedTerm + `(?:s)?\b`)

	// Replace the term with the same term wrapped in brackets, avoiding rewrapping terms
	return termRegex.ReplaceAllStringFunc(text, func(matched string) string {
		// Only wrap the term in brackets if it's not already wrapped
		if strings.HasPrefix(matched, "[") && strings.HasSuffix(matched, "]") {
			return matched // Term is already wrapped, skip it
		}
		for i, entry := range Data.Lexicon {
			if entry.Term == term && !containsSynonym(entry.Synonyms, entry.Term, matched) {
				fmt.Printf("Adding synonym %s to %s\n", matched, term)
				Data.Lexicon[i].Synonyms = append(entry.Synonyms, matched)
			}
		}
		return fmt.Sprintf("[%s]", matched)
	})
}

// Main function to apply the term replacements
func addLinksTemplateFunction(text string) string {
	// Iterate over the lexicon and replace terms with brackets
	for _, entry := range Data.Lexicon {
		text = addLinks(text, entry.Term)
		for _, synonym := range entry.Synonyms {
			text = addLinks(text, synonym)
		}
	}

	return text
}

func linkPrepTemplateFunction(text string) string {
	return "#" + strings.ToLower(strings.ReplaceAll(text, " ", "-"))
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
		"addLinks": func(s string) string {
			return addLinksTemplateFunction(s)
		},
		"linkPrep": func(s string) string {
			return linkPrepTemplateFunction(s)
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
