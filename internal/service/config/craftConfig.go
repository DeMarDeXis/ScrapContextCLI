package config

import (
	"TypeCatParser/pkg/config"
	"fmt"
	"os"
	"path/filepath"
)

type ConfigService struct{}

func NewConfService() *ConfigService {
	return &ConfigService{}
}

// InitDefault — creates default config
func (cs *ConfigService) InitDefault() error {
	const op = "config.service.InitDefault"

	if config.IsConfigExist() {
		return fmt.Errorf("%s: config already exists", op)
	}

	if err := os.MkdirAll(config.ConfigDir, 0755); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	path := filepath.Join(config.ConfigDir, config.ConfigFullFileName)
	if err := os.WriteFile(path, []byte(config.DefaultConfigTemplateString()), 0644); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// GetConfigPath — getter human-readable path to config
func (cs *ConfigService) GetConfigPath() string {
	return config.GetDefaultConfigPath()
}

// UpdateConfig - updates next values in config
func (cs *ConfigService) UpdateConfig(rootPath, outputPath string) error {
	const op = "config.service.UpdateConfig"

	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if cfg == nil {
		cfg = &config.Config{
			RootPath:   config.DefaultRootPath,
			OutputPath: config.DefaultOutputPath,
		}
	}

	if rootPath != "" {
		cfg.RootPath = rootPath
	}
	if outputPath != "" {
		cfg.OutputPath = outputPath
	}

	if err := config.SaveConfig(cfg); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Clean - removes config and export directory
func (cs *ConfigService) Clean() error {
	const op = "config.service.Clean"

	cfgPath := cs.GetConfigPath()
	if err := os.Remove(cfgPath); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("%s: remove config failed: %w", op, err)
		}
	}

	exportDir := "./export"
	if err := os.RemoveAll(exportDir); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("%s: remove export failed: %w", op, err)
		}
	}

	//filePath := "./export/context.md"
	//if err := os.Remove(filePath); err != nil {
	//	if !os.IsNotExist(err) {
	//		return fmt.Errorf("%s: remove export failed: %w", op, err)
	//	}
	//}

	return nil
}
