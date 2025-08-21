// SPDX-FileCopyrightText: Copyright 2025 The OSPS Authors
// SPDX-License-Identifier: Apache-2.0

package baseline

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func collapseNewlines(s string) string {
	return strings.ReplaceAll(s, "\n", " ")
}

func containsSynonym(list []string, entryTerm, term string) bool {
	if strings.EqualFold(entryTerm, term) {
		return true
	}
	for _, item := range list {
		if strings.EqualFold(item, term) {
			return true
		}
	}
	return false
}

// Check whether there's an unmatched open bracket before the term
func isWrapped(text, matched string) bool {
	beforeIndex := strings.Index(text, matched)
	if beforeIndex == -1 {
		return true
	}
	substrBeforeTerm := text[:beforeIndex]

	openBrackets := strings.Count(substrBeforeTerm, "[")
	closeBrackets := strings.Count(substrBeforeTerm, "]")

	return openBrackets > closeBrackets
}

// Function to add links by wrapping terms with square brackets
func addLinks(lexicon []LexiconEntry, text, term string) string {
	// Escape any special characters in the term to use in regex
	escapedTerm := regexp.QuoteMeta(term)

	// Create a regular expression to match the term as a whole word
	// The `(?i)` part makes it case-insensitive, and `\b` ensures whole word matching
	termRegex := regexp.MustCompile(`(?i)\b` + escapedTerm + `(?:s)?\b`)

	// Replace the term with the same term wrapped in brackets, and avoid rewrapping terms
	return termRegex.ReplaceAllStringFunc(text, func(matched string) string {
		if isWrapped(text, matched) {
			return matched // Skip wrapping if already wrapped in brackets
		}

		for i, entry := range lexicon {
			if entry.Term == term && !containsSynonym(entry.Synonyms, entry.Term, matched) {
				lexicon[i].Synonyms = append(entry.Synonyms, matched) //nolint:gocritic
			}
		}
		return fmt.Sprintf("[%s]", matched)
	})
}

// Main function to apply the term replacements
func addLinksTemplateFunction(lexicon []LexiconEntry, text string) string {
	// Iterate over the lexicon and replace terms with brackets
	for _, entry := range lexicon {
		text = addLinks(lexicon, text, entry.Term)
		for _, synonym := range entry.Synonyms {
			text = addLinks(lexicon, text, synonym)
		}
	}

	return text
}

func asLinkTemplateFunction(text string) string {
	return "#" + strings.ToLower(
		strings.ReplaceAll(
			strings.ReplaceAll(text, " ", "-"), ".", ""))
}

// loop through maturityLevels
// to see if any are higher than the targetMaturity
func maxLevel(maturityLevels []string, targetMaturity int) bool {
	var out bool
	for _, maturity := range maturityLevels {
		maturityInt, err := strconv.Atoi(maturity[len(maturity)-1:])
		if err != nil {
			fmt.Println(err)
			return false
		}
		if maturityInt == targetMaturity {
			out = true
		}
		if maturityInt < targetMaturity {
			return false
		}
	}
	return out
}
