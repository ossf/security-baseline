// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import "github.com/spf13/cobra"

func Execute() error {
	rootCmd := &cobra.Command{
		Use:  "baseline-compiler",
		Long: `Baseline Compiler reads the Basline YAML and outputs it as a markdown document.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help() //nolint
		},
	}
	addCompile(rootCmd)

	return rootCmd.Execute()
}
