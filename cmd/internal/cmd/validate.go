// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ossf/security-baseline/pkg/baseline"
)

type validateOptions struct {
	baselinePath string
}

// Validate the options in context with arguments
func (o *validateOptions) Validate() error {
	errs := []error{}

	if o.baselinePath == "" {
		errs = append(errs, errors.New("baseline data path not specified"))
	}
	return errors.Join(errs...)
}

func (o *validateOptions) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(
		&o.baselinePath, "baseline", "b", defaultBaselinePath, "path to directory containing the baseline YAML definitions",
	)
}

// addValidate adds the compile subcommand to the parent command
func addValidate(parentCmd *cobra.Command) {
	opts := validateOptions{}
	validateCmd := &cobra.Command{
		Use:           "validate",
		Short:         "Validate the baseline data files",
		SilenceUsage:  false,
		SilenceErrors: true,
		PreRunE: func(_ *cobra.Command, args []string) error {
			if opts.baselinePath != "" && len(args) > 0 && opts.baselinePath != args[0] {
				return fmt.Errorf("baseline data path specified twice")
			}

			if len(args) > 0 {
				opts.baselinePath = args[0]
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.Validate(); err != nil {
				return err
			}

			cmd.SilenceUsage = true

			// Parse the data files
			loader := baseline.NewLoader()
			err := loader.Load(opts.baselinePath)
			if err != nil {
				return err
			}

			// Generate the rendered version
			err = loader.Validate()
			if err != nil {
				fmt.Fprint(os.Stderr, "\n❌ Error validating the baseline data:\n")
				return err
			}
			fmt.Printf("\n✅ Baseline YAML data OK\n\n")
			return nil
		},
	}
	opts.AddFlags(validateCmd)
	parentCmd.AddCommand(validateCmd)
}
