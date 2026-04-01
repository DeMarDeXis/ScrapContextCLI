package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const (
	configName         = "parserDump"
	configExt          = "yaml"
	ConfigDir          = "config"
	ConfigFullFileName = configName + "." + configExt
	DefaultRootPath    = "./"
	DefaultOutputPath  = "./export/context.md"
)

type Config struct {
	RootPath   string   `mapstructure:"root_path"`
	OutputPath string   `mapstructure:"output_path"`
	Exclude    []string `mapstructure:"exclude_patterns,omitempty"`
}

func LoadConfig() (*Config, error) {
	const op = "pkg.config.LoadConfig"
	var cfg Config
	viper.SetConfigName(configName)
	viper.SetConfigType(configExt)
	viper.AddConfigPath(ConfigDir)

	if err := viper.ReadInConfig(); err != nil {
		// If this is "file not found" error — it's not a FATAL error
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, nil // 👈 мягкий возврат
		}
		// Other error it's real error
		return nil, fmt.Errorf("%s.ReadInConfig: %w", op, err)
	}
	// ⚠️ Achtung: viper.ConfigFileNotFoundError — it's special error type, you need to import or check with errors.As. You can check with os.IsNotExist
	// Alt checker (universal)
	//if err := viper.ReadInConfig(); err != nil {
	//	if os.IsNotExist(err) || strings.Contains(err.Error(), "not found") {
	//		return nil, nil
	//	}
	//	return nil, fmt.Errorf("%s.ReadInConfig: %w", op, err)
	//}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("%s.Unmarshal: %w", op, err)
	}

	if cfg.RootPath == "" {
		cfg.RootPath = DefaultRootPath
	}
	if cfg.OutputPath == "" {
		cfg.OutputPath = DefaultOutputPath
	}

	return &cfg, nil
}

func SaveConfig(cfg *Config) error {
	const op = "pkg.config.SaveConfig"

	if err := os.MkdirAll(ConfigDir, 0755); err != nil {
		return fmt.Errorf("%s.MkdirAll: %w", op, err)
	}

	viper.SetConfigName(configName)
	viper.SetConfigType(configExt)
	viper.AddConfigPath(ConfigDir)

	viper.Set("root_path", cfg.RootPath)
	viper.Set("output_path", cfg.OutputPath)
	if len(cfg.Exclude) > 0 {
		viper.Set("exclude_patterns", cfg.Exclude)
	}

	//if err := viper.SafeWriteConfig(); err != nil {
	//	if os.IsExist(err) {
	//		return viper.WriteConfig()
	//	}
	//	return fmt.Errorf("%s.SafeWriteConfig: %w", op, err)
	//}

	if err := viper.SafeWriteConfig(); err != nil {
		// Если файл уже есть — используем WriteConfig (перезапись)
		// Проверяем по тексту ошибки, т.к. os.IsExist не срабатывает с viper
		if strings.Contains(err.Error(), "Already Exists") {
			return viper.WriteConfig()
		}
		return fmt.Errorf("%s.SafeWriteConfig: %w", op, err)
	}

	return nil
}

func IsConfigExist() bool {
	path := GetDefaultConfigPath()
	_, err := os.Stat(path)
	return err == nil
}

// GetDefaultConfigPath — returns full path to default config
// Always relative to current directory, when .exe is run
func GetDefaultConfigPath() string {
	return filepath.Join(ConfigDir, ConfigFullFileName)
}

// DefaultConfigTemplateString — returns default config template as string
func DefaultConfigTemplateString() string {
	return `# TypeCatParser — config-yaml
# File: config/parserDump.yaml

# Root directory to parse (relative by .exe)
root_path: "./"

# It's file for output result
output_path: "./export/context.md"

# Patter exception (type .gitignore)
# - папки: "vendor/**", "node_modules/**"
# - файлы: "*.log", "*.exe"
# - "**" означает рекурсивно
exclude_patterns:
  - ".git/**"
  - "vendor/**"
  - "node_modules/**"
  - "export/**"
  - "*.log"
  - "*.sum"
`
}
