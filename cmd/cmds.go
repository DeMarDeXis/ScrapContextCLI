package cmd

func init() {
	rootCmd.AddCommand(parseCmd)
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configHelpCmd)
	configCmd.AddCommand(configCleanCmd)

	configSetCmd.Flags().String("root_path", "", "Directory to parse (relative path)")
	configSetCmd.Flags().String("output_path", "", "Output file path for context dump")
	configCleanCmd.Flags().BoolP("force", "f", false, "Force clean config")
}
