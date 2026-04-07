package service

import (
	"TypeCatParser/internal/service/config"
	"TypeCatParser/internal/service/parser"

	"log/slog"
)

type Config interface {
	InitDefault() error
	GetConfigPath() string
	UpdateConfig(rootPath, outputPath string) error
	Clean() error
}

type Parser interface {
	Parse(rootDir, outputPath string) error
}

type Service struct {
	Config
	Parser
}

func NewService(log *slog.Logger, excludePatterns []string) *Service {
	patterns := excludePatterns
	if len(patterns) == 0 { // Another say "default"
		patterns = []string{
			".git/**", "node_modules/**", "vendor/**",
			"export/**", "dist/**", "build/**",
			"*.log", "*.sum", "*.exe", "*.dll", "*.so",
		}
	}

	return &Service{
		Config: config.NewConfService(),
		Parser: parser.NewParserService(log, patterns),
	}
}
