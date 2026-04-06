// SPDX-FileCopyrightText: Copyright 2026 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"io"

	"github.com/goccy/go-yaml"

	"github.com/ossf/security-baseline/pkg/types"
)

// ExportGemara writes the assembled ControlCatalog as Gemara-compliant YAML.
func (g *Generator) ExportGemara(b *types.Baseline, w io.Writer) error {
	out, err := yaml.Marshal(&b.Catalog)
	if err != nil {
		return fmt.Errorf("marshaling catalog: %w", err)
	}
	if _, err := w.Write(out); err != nil {
		return fmt.Errorf("writing YAML: %w", err)
	}
	return nil
}
