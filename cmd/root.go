package cmd

import (
	"TypeCatParser/internal/service"
	"TypeCatParser/pkg/config"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

type cmdContext struct {
	srvc   *service.Service
	config *config.Config
	log    *slog.Logger
}

var ctx = &cmdContext{}

var rootCmd = &cobra.Command{
	Use:   "typecatparser",
	Short: "Parser project app",
}

func Execute(logg *slog.Logger, cfg *config.Config) {
	ctx.log = logg
	ctx.config = cfg
	ctx.srvc = service.NewService(logg)

	if err := rootCmd.Execute(); err != nil {
		handleErr(logg, err)
	}
}

func handleErr(log *slog.Logger, err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	log.Error("failed to execute command", slog.Any("error", err.Error()))
	os.Exit(1)
}
