// SPDX-FileCopyrightText: Copyright 2026 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import "testing"

func TestAsLink(t *testing.T) {
	tests := map[string]string{
		"CI/CD Pipeline":              "#cicd-pipeline", // kramdown deletes "/", not "-"
		"Multi-factor Authentication": "#multi-factor-authentication",
		"OSPS-AC-01.01":               "#osps-ac-0101", // "." dropped, matching published anchors
	}
	for in, want := range tests {
		if got := asLinkTemplateFunction(in); got != want {
			t.Errorf("asLink(%q) = %q, want %q", in, got, want)
		}
	}
}
