package parser

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	markdownHeader   = "# Project Context Dump\n\n"
	codeBlockStart   = "```text\n"
	codeBlockEnd     = "\n```\n\n"
	fileHeaderFormat = "## 📄 %s\n\n"
	dirHeaderFormat  = "## 📁 %s\n\n"
)

type ParserService struct {
	log    *slog.Logger
	filter *Filter
}

func NewParserService(log *slog.Logger, excludePatterns []string) *ParserService {
	return &ParserService{
		log:    log,
		filter: NewFilter(excludePatterns),
	}
}

func (ps *ParserService) Parse(rootDir, outputPath string) error {
	const op = "internal.service.parser.Parse"

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("in [%s] failed to create output directory: %w", op, err)
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("in [%s] failed to create output file: %w", op, err)
	}
	defer outFile.Close()

	_, err = outFile.WriteString("# Project Context Dump\n\n")
	if err != nil {
		return fmt.Errorf("in [%s] failed to write [HEADER] in output file: %w", op, err)
	}

	absOutput, _ := filepath.Abs(outputPath)

	return filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("%s: walk error at %s: %w", op, path, err)
		}

		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			ps.log.Warn("relative path failed", slog.String("path", path))
			return nil
		}

		if abs, _ := filepath.Abs(path); abs == absOutput {
			return nil
		}

		// Add new filter
		if ps.filter.ShouldSkip(relPath, d.IsDir()) {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		if d.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			ps.log.Warn("read file failed", slog.String("path", relPath))
			return nil
		}

		return writeFileContent(outFile, relPath, content)
	})
}

// writeFileContent — writes file content to output file
func writeFileContent(out *os.File, relPath string, content []byte) error {
	_, err := fmt.Fprintf(out, "## 📄 %s\n\n```text\n%s\n```\n\n", relPath, content)
	return err
}
