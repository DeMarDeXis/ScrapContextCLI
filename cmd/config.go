package cmd

import (
	"TypeCatParser/pkg/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage parserDump.yaml configuration",
	Long: `Configuration manager for TypeCatParser.

		Use subcommands to initialize, view, or update your parser settings.
		All paths are relative to the directory where you run the .exe file.`,
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate default parserDump.yaml",
	Long:  `Create a new config file with sensible defaults and helpful comments.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ctx.srvc.InitDefault(); err != nil {
			handleErr(ctx.log, err)
			return
		}

		fmt.Printf("✅ Config created: %s\n", ctx.srvc.GetConfigPath())
		fmt.Println("💡 Next: ./parseDirContext.exe config show")
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Update configuration values",
	Long: `Update one or more configuration values.

Examples:
  # Change root directory to parse
  ./parseDirContext.exe config set --root_path=./src

  # Change output file location
  ./parseDirContext.exe config set --output_path=./dump/context.md

  # Change both at once
  ./parseDirContext.exe config set --root_path=./pkg --output_path=./out/dump.md`,
	Run: func(cmd *cobra.Command, args []string) {
		rootPath, _ := cmd.Flags().GetString("root_path")
		outputPath, _ := cmd.Flags().GetString("output_path")

		if rootPath == "" && outputPath == "" {
			fmt.Fprintln(os.Stderr, "⚠️  No values to update. Use --root_path or --output_path")
			fmt.Fprintln(os.Stderr, "💡 Example: ./parseDirContext.exe config set --root_path=./src")
			return
		}

		if err := ctx.srvc.UpdateConfig(rootPath, outputPath); err != nil {
			handleErr(ctx.log, err)
			return
		}

		fmt.Println("✅ Config updated")
		if rootPath != "" {
			fmt.Printf("  Root path: %s\n", rootPath)
		}
		if outputPath != "" {
			fmt.Printf("  Output path: %s\n", outputPath)
		}
	},
}

// config show
var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Display current configuration",
	Long:  `Print current values from parserDump.yaml in a human-readable format.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			handleErr(ctx.log, err)
			return
		}
		if cfg == nil {
			fmt.Println("⚠️  Config not found. Please run: ./<<>>.exe config init")
			return
		}

		fmt.Println("📋 Current configuration:")
		fmt.Printf("   root_path:        %s\n", cfg.RootPath)
		fmt.Printf("   output_path:      %s\n", cfg.OutputPath)
		if len(cfg.Exclude) > 0 {
			fmt.Printf("   exclude_patterns: %v\n", cfg.Exclude)
		}
		fmt.Printf("\n📄 Config file: %s\n", config.GetDefaultConfigPath())
	},
}

var configCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove config file and export directory",
	Long: `Remove configuration and generated files.

This will delete:
  • config/parserDump.yaml
  • export/ directory (if exists)

⚠️  This action cannot be undone!`,
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")

		if !force {
			fmt.Println("⚠️  This action cannot be undone!")
			fmt.Println("   Are you sure you want to proceed? (y/n)")
			var confirmation string
			fmt.Scanln(&confirmation)

			if confirmation != "y" && confirmation != "Y" {
				fmt.Println("✅ Cancelled.")
				return
			}
		}

		if err := ctx.srvc.Clean(); err != nil {
			handleErr(ctx.log, err)
			return
		}

		fmt.Println("✅ Config cleaned")
	},
}

// its cmd out cmds list
var configHelpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show detailed help for config commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TypeCatParser Configuration Guide")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()
		fmt.Println("🔹 INIT — Create default config")
		fmt.Println("   ./parseDirContext.exe config init")
		fmt.Println("  ")
		fmt.Println("   Creates config/parserDump.yaml with:")
		fmt.Println("   • root_path: \"./\"           — parse from current dir")
		fmt.Println("   • output_path: \"./export/context.md\"")
		fmt.Println("   • exclude_patterns: common ignores")
		fmt.Println()
		fmt.Println("🔹 SET — Update values")
		fmt.Println("   ./parseDirContext.exe config set --root_path=./src")
		fmt.Println("   ./parseDirContext.exe config set --output_path=./dump/result.md")
		fmt.Println()
		fmt.Println("   Flags:")
		fmt.Println("   • --root_path=PATH    — directory to start parsing")
		fmt.Println("   • --output_path=FILE  — where to write the context dump")
		fmt.Println()
		fmt.Println("   Notes:")
		fmt.Println("   • Paths are relative to where you run .exe")
		fmt.Println("   • Both flags are optional — update only what you need")
		fmt.Println("   • File is saved immediately")
		fmt.Println()
		fmt.Println("🔹 SHOW — View current settings")
		fmt.Println("   ./parseDirContext.exe config show")
		fmt.Println()
		fmt.Println("   Prints current values in readable format.")
		fmt.Println()
		fmt.Println("🔹 CLEAN — Remove config and export directory")
		fmt.Println("   ./parseDirContext.exe config clean")
		fmt.Println()
		fmt.Println("Flags:")
		fmt.Println("   • --force   — skip confirmation prompt")
		fmt.Println()
		fmt.Println("🔹 TIPS")
		fmt.Println("   • Run 'config init' once, then use 'set' for tweaks")
		fmt.Println("   • Use 'parse' command after configuring")
		fmt.Println("   • Edit config/parserDump.yaml manually for advanced settings")
	},
}
