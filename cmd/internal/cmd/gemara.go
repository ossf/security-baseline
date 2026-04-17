// SPDX-FileCopyrightText: Copyright 2026 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ossf/security-baseline/pkg/baseline"
)

type gemaraOptions struct {
	outPath      string
	baselinePath string
}

func (o *gemaraOptions) Validate() error {
	if o.baselinePath == "" {
		return errors.New("baseline path not specified")
	}
	return nil
}

func (o *gemaraOptions) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(
		&o.baselinePath, "baseline", "b", defaultBaselinePath, "path to directory containing the baseline YAML data",
	)
	cmd.PersistentFlags().StringVarP(
		&o.outPath, "out", "o", "", "path to output file (defaults to STDOUT)",
	)
}

func addGemara(parentCmd *cobra.Command) {
	opts := gemaraOptions{}
	gemaraCmd := &cobra.Command{
		Use:   "gemara -o baseline.gemara.yaml",
		Short: "Export the assembled baseline as Gemara YAML",
		Long: `Assembles the baseline YAML sources into a single Gemara ControlCatalog
and writes it to a file or STDOUT. The output can be validated externally
with cue vet.`,
		SilenceUsage:  false,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
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

			w := os.Stdout
			if opts.outPath != "" {
				f, err := os.Create(opts.outPath)
				if err != nil {
					return fmt.Errorf("creating output file: %w", err)
				}
				defer f.Close() //nolint:errcheck
				w = f
			}

			gen := baseline.NewGenerator()
			if err := gen.ExportGemara(bline, w); err != nil {
				return err
			}

			if opts.outPath != "" {
				fmt.Fprintf(os.Stderr, "Gemara YAML written to %s\n", opts.outPath)
			}
			return nil
		},
	}
	opts.AddFlags(gemaraCmd)
	parentCmd.AddCommand(gemaraCmd)
}
