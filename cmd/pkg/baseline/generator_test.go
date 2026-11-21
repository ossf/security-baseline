// SPDX-FileCopyrightText: Copyright 2026 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ossf/security-baseline/pkg/types"
)

func TestReleaseVersion(t *testing.T) {
	for _, tc := range []struct {
		artifactVersion string
		want            string
		release         bool
	}{
		{"", "devel", false},
		{"v2026.08.28", "2026.08.28", true},
		{"2026.08.28", "2026.08.28", true},
	} {
		g := &Generator{ArtifactVersion: tc.artifactVersion}
		if got := g.ReleaseVersion(); got != tc.want {
			t.Errorf("ReleaseVersion() for %q = %q, want %q", tc.artifactVersion, got, tc.want)
		}
		if got := g.IsRelease(); got != tc.release {
			t.Errorf("IsRelease() for %q = %v, want %v", tc.artifactVersion, got, tc.release)
		}
	}
}

func TestExportMarkdownStampsVersion(t *testing.T) {
	dir := t.TempDir()
	tmpl := filepath.Join(dir, "template.md")
	content := "Version: {{ releaseVersion }}\n\n{{ if not isRelease -}}\nNot for production use.\n{{ end -}}\nbody\n"
	if err := os.WriteFile(tmpl, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		artifactVersion string
		want            string
	}{
		{"", "Version: devel\n\nNot for production use.\nbody\n"},
		{"v2026.08.28", "Version: 2026.08.28\n\nbody\n"},
	} {
		out := filepath.Join(dir, "out.md")
		g := &Generator{ArtifactVersion: tc.artifactVersion}
		if err := g.ExportMarkdown(&types.Baseline{}, tmpl, out); err != nil {
			t.Fatal(err)
		}
		got, err := os.ReadFile(out)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != tc.want {
			t.Errorf("artifact version %q rendered %q, want %q", tc.artifactVersion, got, tc.want)
		}
	}
}
