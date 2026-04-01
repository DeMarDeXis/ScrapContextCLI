package parser

import "strings"

// shouldSkipDir — dir which we should skip
func shouldSkipDir(relPath string) bool {
	skip := []string{
		".git", "node_modules", "vendor",
		"export", "dist", "build",
	}
	for _, s := range skip {
		if relPath == s || strings.HasPrefix(relPath, s+"/") {
			return true
		}
	}
	return false
}

// shouldSkipFile — file which we should skip
func shouldSkipFile(relPath string) bool {
	skipExts := []string{
		".exe", ".dll", ".so", ".dylib",
		".png", ".jpg", ".jpeg", ".gif",
		".zip", ".tar", ".gz", ".pdf",
		".log", ".sum",
	}
	for _, ext := range skipExts {
		if strings.HasSuffix(relPath, ext) {
			return true
		}
	}
	return false
}
