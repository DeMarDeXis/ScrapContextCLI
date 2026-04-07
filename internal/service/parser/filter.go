package parser

import (
	gitIgnorePKG "github.com/sabhiram/go-gitignore"
)

type Filter struct {
	ignorer gitIgnorePKG.IgnoreParser
}

// NewFilter compiles patterns to matcher
func NewFilter(patterns []string) *Filter {
	return &Filter{
		ignorer: gitIgnorePKG.CompileIgnoreLines(patterns...),
	}
}

// ShouldSkip checks if the path should be skipped
func (f *Filter) ShouldSkip(relPath string, isDir bool) bool {
	// Normalize: gitignore expects paths without "./" and with forward slashes
	normalized := relPath
	if len(normalized) > 2 && normalized[:2] == "./" {
		normalized = normalized[2:]
	}

	if isDir && len(normalized) > 0 && normalized[len(normalized)-1] != '/' {
		normalized += "/"
	}

	return f.ignorer.MatchesPath(normalized)
}
