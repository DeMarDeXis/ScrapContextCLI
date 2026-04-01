package cmd

import "github.com/spf13/cobra"

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse the dir",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ctx.srvc.Parse(ctx.config.RootPath, ctx.config.OutputPath); err != nil {
			handleErr(ctx.log, err)
		}
	},
}
