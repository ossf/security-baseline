// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import "github.com/spf13/cobra"

const (
	// defaultBaselinePath is the default location where the CLI will look for
	// the baseline YAML files
	defaultBaselinePath = "../baseline"
)

func Execute() error {
	rootCmd := &cobra.Command{
		Use:  "baseline-compiler",
		Long: `Baseline Compiler reads the Basline YAML and outputs it as a markdown document.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// Add the subcommands
	addCompile(rootCmd)
	addValidate(rootCmd)
	addOscal(rootCmd)

	return rootCmd.Execute()
}
