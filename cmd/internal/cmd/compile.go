// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/ossf/security-baseline/pkg/baseline"
	"github.com/spf13/cobra"
)

type compileOptions struct {
	outPath               string
	checklistOutPath      string
	baselinePath          string
	checklistTemplatePath string
	templatePath          string
	checklist             bool
	validate              bool
}

// Validate the options in context with arguments
func (o *compileOptions) Validate() error {
	errs := []error{}

	if o.baselinePath == "" {
		errs = append(errs, errors.New("baseline path not specified"))
	}
	return errors.Join(errs...)
}

func (o *compileOptions) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(
		&o.baselinePath, "baseline", "b", "../baseline", "path to directory containing the baseline YAML data",
	)

	cmd.PersistentFlags().StringVarP(
		&o.checklistOutPath, "checklist-output", "", "",
		"path to checklist output file (checklist only generated if specified)",
	)

	cmd.PersistentFlags().StringVarP(
		&o.checklistTemplatePath, "checklist-template", "", "template-checklist.md",
		"path to the checklist template file",
	)

	cmd.PersistentFlags().StringVarP(
		&o.outPath, "output", "o", "", "path to output file",
	)

	cmd.PersistentFlags().StringVarP(
		&o.templatePath, "template", "t", "template.md", "path to the markdown template file",
	)

	cmd.PersistentFlags().BoolVarP(
		&o.validate, "validate", "v", true, "validate data inegrity before rendering",
	)
}

// addCompile adds the compile subcommand to the parent command
func addCompile(parentCmd *cobra.Command) {
	opts := compileOptions{}
	compileCmd := &cobra.Command{
		Use:           "compile [file]",
		Short:         "Compile a YAML file of security controls",
		SilenceUsage:  false,
		SilenceErrors: true,
		PreRunE: func(_ *cobra.Command, args []string) error {
			if opts.outPath != "" && len(args) > 1 && opts.outPath != args[1] {
				return fmt.Errorf("output path specified twice")
			}

			if len(args) > 1 {
				opts.outPath = args[1]
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.Validate(); err != nil {
				return err
			}

			cmd.SilenceUsage = true

			// Load the baseline data
			loader := baseline.NewLoader()
			loader.DataPath = opts.baselinePath

			bline, err := loader.Load()
			if err != nil {
				return err
			}

			// Validate the data
			validator := baseline.NewValidator()
			err = validator.Check(bline)
			if err != nil {
				if opts.validate {
					fmt.Fprint(os.Stderr, "\n❌ Error validating the baseline data:\n")
					return err
				}

				// if the validation flag is off, we still validate but only warn
				// the user about it
				fmt.Fprint(os.Stderr, "\n⚠️ Error validating the baseline data (still rendering)")
			}

			// Generate the rendered version
			gen := baseline.NewGenerator()

			if opts.outPath == "" {
				fmt.Fprintf(os.Stderr, "\n⚠️  No output path specified. Not rendering Baseline.")
			} else {
				if err := gen.ExportMarkdown(bline, opts.templatePath, opts.outPath); err != nil {
					return fmt.Errorf("writing markdown render: %w", err)
				}
				fmt.Fprintf(os.Stderr, "\n✅ Baseline rendered to %s\n", opts.outPath)
			}

			fmt.Fprintf(os.Stderr, "\nℹ️  Counts\n")
			for c := range bline.Categories {
				fmt.Fprintf(os.Stderr, " OSPS-%s: %d controls\n", c, len(bline.Categories[c].Controls))
			}
			fmt.Fprintf(os.Stderr, "\n+ %d lexicon entries\n", len(bline.Lexicon))

			// Print a checklist if they asked for it
			if opts.checklistOutPath != "" {
				if err := gen.ExportMarkdown(bline, opts.checklistTemplatePath, opts.checklistOutPath); err != nil {
					return fmt.Errorf("checklist creation: %w", err)
				}

				fmt.Fprintf(os.Stderr, "\n✅ Checklist rendered to %s",
					opts.checklistOutPath)
			}

			return nil
		},
	}
	opts.AddFlags(compileCmd)
	parentCmd.AddCommand(compileCmd)
}
