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

func NewService(log *slog.Logger) *Service {
	return &Service{
		Config: config.NewConfService(),
		Parser: parser.NewParserService(log),
	}
}
