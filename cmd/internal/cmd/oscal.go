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

var appname = "baseline"

type oscalOptions struct {
	outPath      string
	baselinePath string
}

// Validate the options in context with arguments
func (o *oscalOptions) Validate() error {
	errs := []error{}

	if o.baselinePath == "" {
		errs = append(errs, errors.New("baseline path not specified"))
	}
	return errors.Join(errs...)
}

func (o *oscalOptions) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(
		&o.baselinePath, "baseline", "b", "", "path to directory containing the baseline YAML data",
	)

	cmd.PersistentFlags().StringVarP(
		&o.outPath, "out", "o", "", "path to output file (defaults to STDOUT)",
	)
}

func addOscal(parentCmd *cobra.Command) {
	opts := oscalOptions{}
	packCmd := &cobra.Command{
		Short: "writes the baseline definition to an oscal json catalog",
		Long: fmt.Sprintf(`
%s oscal: Write the OSPS Baseline to oscal definitions.

This subcommand exports the OSPS Baseline data to OSCAL (Open Security Controls
Assessment Language). This lets automated tools understand the criteria set as 
OSCAL controls.

`, appname),
		Use:           "oscal -o osps.oscal.json",
		SilenceUsage:  false,
		SilenceErrors: true,
		PreRunE: func(_ *cobra.Command, args []string) error {
			if opts.outPath != "" && len(args) > 1 && opts.outPath != args[1] {
				return fmt.Errorf("out path specified twice")
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

			loader := baseline.NewLoader()
			loader.DataPath = opts.baselinePath

			bline, err := loader.Load()
			if err != nil {
				return err
			}

			// TODO: Open the output file

			gen := baseline.NewGenerator()
			if err := gen.ExportOSCAL(bline, os.Stdout); err != nil {
				return err
			}

			return nil
		},
	}
	opts.AddFlags(packCmd)
	parentCmd.AddCommand(packCmd)
}
