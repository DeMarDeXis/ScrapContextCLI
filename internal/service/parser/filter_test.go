package parser

import "testing"

func TestFilter_ShouldSkip(t *testing.T) {
	patterns := []string{
		".git/**", "node_modules/**", "vendor/**",
	}
	f := NewFilter(patterns)

	tests := []struct {
		path     string
		isDir    bool
		wantSkip bool
	}{
		{"file.txt", false, false},
		{".git/objects", true, true},
		{"vendor/github.com/pkg", true, true},
		{"src/main.go", false, false},
		{"docs/README.md", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			if got := f.ShouldSkip(tt.path, tt.isDir); got != tt.wantSkip {
				t.Errorf("ShouldSkip(%q, %v) = %v, want %v",
					tt.path, tt.isDir, got, tt.wantSkip)
			}
		})
	}
}
