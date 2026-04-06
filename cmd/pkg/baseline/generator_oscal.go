// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"encoding/json"
	"fmt"
	"io"

	oscal "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
	"github.com/gemaraproj/go-gemara/gemaraconv"

	"github.com/ossf/security-baseline/pkg/types"
)

const (
	VersionOSPS   = "devel"
	controlHREF   = "https://baseline.openssf.org/version/%s#%s"
	canonicalHREF = "https://baseline.openssf.org"
)

func (g *Generator) ExportOSCAL(b *types.Baseline, w io.Writer) error {

	oscalCatalog, err := gemaraconv.ControlCatalog(&b.Catalog).ToOSCAL(
		gemaraconv.WithCanonicalHrefFormat(canonicalHREF),
		gemaraconv.WithControlHref(controlHREF))
	if err != nil {
		return fmt.Errorf("converting to OSCAL: %w", err)
	}
	oscalCatalog.Metadata.Version = VersionOSPS

	models := oscal.OscalModels{
		Catalog: &oscalCatalog,
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(models); err != nil {
		return fmt.Errorf("encoding OSCAL JSON: %w", err)
	}
	return nil
}
