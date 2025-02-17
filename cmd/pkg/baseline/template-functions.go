package main

import (
	"fmt"
	"regexp"
	"strings"
)

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
func isWrapped(text string, matched string) bool {
	beforeIndex := strings.Index(text, matched)
	substrBeforeTerm := text[:beforeIndex]

	openBrackets := strings.Count(substrBeforeTerm, "[")
	closeBrackets := strings.Count(substrBeforeTerm, "]")

	return openBrackets > closeBrackets
}

// Function to add links by wrapping terms with square brackets
func addLinks(text, term string) string {
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

		for i, entry := range Lexicon {
			if entry.Term == term && !containsSynonym(entry.Synonyms, entry.Term, matched) {
				Lexicon[i].Synonyms = append(entry.Synonyms, matched)
			}
		}
		return fmt.Sprintf("[%s]", matched)
	})
}

// Main function to apply the term replacements
func addLinksTemplateFunction(text string) string {
	// Iterate over the lexicon and replace terms with brackets
	for _, entry := range Lexicon {
		text = addLinks(text, entry.Term)
		for _, synonym := range entry.Synonyms {
			text = addLinks(text, synonym)
		}
	}

	return text
}

func asLinkTemplateFunction(text string) string {
	return "#" + strings.ToLower(strings.ReplaceAll(text, " ", "-"))
}
