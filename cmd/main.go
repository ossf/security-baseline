package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// Struct for representing each criteria entry
type Criteria struct {
	ID                    string   `yaml:"id"`
	MaturityLevel         int      `yaml:"maturity_level"`
	Category              string   `yaml:"category"`
	CriteriaText          string   `yaml:"criteria"`
	Objective             string   `yaml:"objective"`
	Implementation        string   `yaml:"implementation"`
	ControlMappings       string   `yaml:"control_mappings"`
	SecurityInsightsValue string   `yaml:"security_insights_value"`
	ScorecardProbe        []string `yaml:"scorecard_probe"`
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
var (
	Data         Baseline
	templatePath = "template.md"
	YAMLpath     string
	OutputPath   string

	rootCmd = &cobra.Command{
		Use:  "baseline-compiler",
		Long: `Baseline Compiler reads the Basline YAML and outputs it as a markdown document.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	compileCmd = &cobra.Command{
		Use:   "compile [file]",
		Short: "Compile a YAML file of security criteria",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			err := readYAMLFile()
			if err != nil {
				log.Fatalf("Error reading YAML file: %v", err)
			}

			err = generateBaselineMdFile()
			if err != nil {
				log.Fatalf("Error generating output: %v", err)
			}
			fmt.Println("---")
			fmt.Printf("Output generated to %s\n", OutputPath)
			fmt.Println("Please verify the contents before committing.")
			fmt.Println("Known issues exist with links where one term is a substring of another, such as 'release' and 'release pipeline'")
			fmt.Println("---")
		},
	}
)

func init() {
	compileCmd.Flags().StringVarP(&OutputPath, "output", "o", filepath.Join("..", "baseline.md"), "Output file path")
	viper.BindPFlag("output", compileCmd.Flags().Lookup("output"))

	compileCmd.Flags().StringVarP(&YAMLpath, "file", "f", filepath.Join("..", "baseline.yaml"), "Path to the YAML input file")
	viper.BindPFlag("file", compileCmd.Flags().Lookup("file"))
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
func readYAMLFile() error {
	file, err := os.Open(YAMLpath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var baseline Baseline
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&baseline); err != nil {
		return fmt.Errorf("error decoding YAML: %v", err)
	}
	var entryIDs []string
	for i, entry := range baseline.Criteria {
		// if entry in entryIDs
		if slices.Contains(entryIDs, entry.ID) {
			return fmt.Errorf("duplicate ID for criteria entry %d: %s", i, entry.ID)
		}
		if entry.ID == "" {
			return fmt.Errorf("missing ID for criteria entry %d: %s", i, entry.ID)
		}
		if entry.CriteriaText == "" {
			return fmt.Errorf("missing criteria text for entry #%d: %s", i, entry.ID)
		}
		if entry.Objective == "" {
			return fmt.Errorf("missing objective for entry #%d: %s", i, entry.ID)
		}
		if entry.Implementation == "" {
			return fmt.Errorf("missing implementation for entry #%d: %s", i, entry.ID)
		}
		entryIDs = append(entryIDs, entry.ID)
	}
	Data = baseline
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

// Check whether there's an unmatched open bracket before the term
func isWrapped(text string, matched string) bool {
	beforeIndex := strings.Index(text, matched)
	substrBeforeTerm := text[:beforeIndex]

	openBrackets := strings.Count(substrBeforeTerm, "[")
	closeBrackets := strings.Count(substrBeforeTerm, "]")

	return openBrackets > closeBrackets
}

// Function to add links by wrapping terms with square brackets
func addLinks(text, term string) string {
	// Escape any special characters in the term to use in regex
	escapedTerm := regexp.QuoteMeta(term)

	// Create a regular expression to match the term as a whole word
	// The `(?i)` part makes it case-insensitive, and `\b` ensures whole word matching
	termRegex := regexp.MustCompile(`(?i)\b` + escapedTerm + `(?:s)?\b`)

	// Replace the term with the same term wrapped in brackets, and avoid rewrapping terms
	return termRegex.ReplaceAllStringFunc(text, func(matched string) string {
		if isWrapped(text, matched) {
			return matched // Skip wrapping if already wrapped in brackets
		}

		for i, entry := range Data.Lexicon {
			if entry.Term == term && !containsSynonym(entry.Synonyms, entry.Term, matched) {
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

func asLinkTemplateFunction(text string) string {
	return "#" + strings.ToLower(strings.ReplaceAll(text, " ", "-"))
}

// Function to generate the markdown file
func generateBaselineMdFile() (err error) {
	// Open or create the output file
	oDir := filepath.Dir(OutputPath)
	err = os.MkdirAll(oDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating output directory %s: %w", oDir, err)
	}
	outputFile, err := os.Create(OutputPath)
	if err != nil {
		return fmt.Errorf("error creating output file %s: %w", OutputPath, err)
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
		"asLink": func(s string) string {
			return asLinkTemplateFunction(s)
		},
	}).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Execute the template and write to the output file
	err = tmpl.Execute(outputFile, Data)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
