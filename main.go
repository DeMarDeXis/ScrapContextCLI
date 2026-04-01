package main

import (
	"TypeCatParser/cmd"
	"TypeCatParser/pkg/config"
	"TypeCatParser/pkg/logger/handler/slogpretty"
	"log/slog"
	"os"
)

func main() {
	logg := setupPrettySlogLocal()
	logg.Info("Starting...")

	//if !config.IsConfigExist() {
	//	fmt.Fprintf(os.Stderr, "⚠️  Config not found: %s\n", config.GetDefaultConfigPath())
	//	fmt.Fprintf(os.Stderr, "💡 Run: ./parseDirContext.exe config init\n")
	//	os.Exit(0)
	//}

	cfg, err := config.LoadConfig()
	if err != nil {
		logg.Error("failed to load config", slog.Any("error", err.Error()))
		os.Exit(1)
	}

	// Если конфиг не загружен — просто логируем, но не падаем
	// (команда parse сама проверит и выдаст подсказку)
	if cfg == nil {
		logg.Debug("config not found — some commands may require it")
	} else {
		logg.Debug("config loaded", slog.Any("config", cfg))
	}

	logg.Debug("config loaded", slog.Any("config", cfg))

	logg.Info("App execute")
	cmd.Execute(logg, cfg)
}

func setupPrettySlogLocal() *slog.Logger {
	opts := slogpretty.PrettyHandlersOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handlerLog := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handlerLog)
}

// TODO:
// 		service.parser.filter.go should take the values from config.yaml [main]
//		delete legacy code service.config.craftConfig.go
//		maybe consts should be used in func [writeFileContent]
// 		legacy code in pkg/config/config.go
