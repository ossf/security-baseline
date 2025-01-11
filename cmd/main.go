package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Lexicon      []LexiconEntry
	OutputPath   string
	TemplatePath = "template.md"
	ContentDir   = "../baseline"
	LexiconPath  = filepath.Join(ContentDir, "lexicon.yaml")

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
			baseline, err := newBaseline()
			if err != nil {
				log.Fatalf("Error reading YAML file: %v", err)
				return
			}
			Lexicon = baseline.Lexicon
			baseline.Generate()

			// err = generateBaselineMdFile()
			// if err != nil {
			// 	log.Fatalf("Error generating output: %v", err)
			// }
			// fmt.Println("---")
			// fmt.Printf("Output generated to %s\n", OutputPath)
			// fmt.Println("Please verify the contents before committing.")
			// fmt.Println("Known issues exist with links where one term is a substring of another, such as 'release' and 'release pipeline'")
			// fmt.Println("---")
		},
	}
)

func init() {
	compileCmd.Flags().StringVarP(&OutputPath, "output", "o", filepath.Join("..", "baseline.md"), "Output file path")
	viper.BindPFlag("output", compileCmd.Flags().Lookup("output"))
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
