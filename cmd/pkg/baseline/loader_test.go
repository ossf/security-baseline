// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"strings"
	"testing"
)

func TestToFileURI(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Unix absolute path",
			input: "/home/user/baseline/OSPS-AC.yaml",
			want:  "file:///home/user/baseline/OSPS-AC.yaml",
		},
		{
			name:  "Windows absolute path",
			input: `C:\Users\foo\baseline\OSPS-AC.yaml`,
			want:  "file:///C:/Users/foo/baseline/OSPS-AC.yaml",
		},
		{
			name:  "Already a file URI",
			input: "file:///already/a/uri.yaml",
			want:  "file:///already/a/uri.yaml",
		},
		{
			name:  "Already an http URI",
			input: "https://example.com/catalog.yaml",
			want:  "https://example.com/catalog.yaml",
		},
		{
			name:  "Relative path (unchanged)",
			input: "baseline/OSPS-AC.yaml",
			want:  "baseline/OSPS-AC.yaml",
		},
		{
			name:  "Malformed file:// (two slashes) is repaired",
			input: "file://C:/Users/foo/baseline/OSPS-AC.yaml",
			want:  "file:///C:/Users/foo/baseline/OSPS-AC.yaml",
		},
		{
			name:  "file:/// (three slashes) is returned unchanged",
			input: "file:///home/user/baseline/OSPS-AC.yaml",
			want:  "file:///home/user/baseline/OSPS-AC.yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toFileURI(tt.input)
			if got != tt.want {
				t.Errorf("toFileURI(%q)\n  got:  %q\n  want: %q", tt.input, got, tt.want)
			}
			// Ensure no Windows drive letter ends up as a host:port in the URI
			if strings.Contains(got, "file://C:") || strings.Contains(got, "file://D:") {
				t.Errorf("toFileURI produced an invalid URI with drive letter as host: %q", got)
			}
		})
	}
}
