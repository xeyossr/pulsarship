package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	flagvars "github.com/xeyossr/pulsarship/internal/cli/flag_variables"
	"github.com/xeyossr/pulsarship/internal/config"
	initShell "github.com/xeyossr/pulsarship/internal/init"
)

var RootCmd = &cobra.Command{
	Use:   "pulsarship",
	Short: "🚀 The minimal, fast, and customizable shell prompt ☄🌌️",
	Long:  "🚀🌌️ The minimal, fast, and customizable shell prompt written in Go 🐹",
}
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Prints the shell function used to execute pulsarship",
}
var GenConfig = &cobra.Command{
	Use:   "gen-config",
	Short: "Generates a default configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		path := config.GetConfigPath(flagvars.ConfigFlag)
		if err := config.WriteDefaultConfig(path); err != nil {
			fmt.Fprintln(os.Stderr, "Error generating config:", err)
			os.Exit(1)
		}
		fmt.Println("Configuration file generated at:", path)
	},
}

var PromptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Prints the full pulsarship prompt",
	Run: func(cmd *cobra.Command, args []string) {
		path := config.GetConfigPath(flagvars.ConfigFlag)
		var err error

		if flagvars.ShowRight {
			err = RunRightPrompt(path, os.Stdout)
		} else {
			err = RunPrompt(path, os.Stdout)
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	},
}

var InitBashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Prints Bash init script",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(initShell.BashInit())
	},
}

var InitZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Prints Zsh init script",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(initShell.ZshInit())
	},
}

var InitFishCmd = &cobra.Command{
	Use:   "fish",
	Short: "Prints Fish init script",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(initShell.FishInit())
	},
}

func init() {
	RootCmd.CompletionOptions.HiddenDefaultCmd = true

	RootCmd.PersistentFlags().BoolVar(&flagvars.DebugMode, "debug", false, "enable debug mode")
	_ = RootCmd.PersistentFlags().MarkHidden("debug")
	RootCmd.PersistentFlags().StringVarP(&flagvars.ConfigFlag, "config", "c", "", "Path to the config file")
	PromptCmd.Flags().BoolVarP(&flagvars.ShowRight, "right", "r", false, "Print the right prompt instead of left prompt")

	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(GenConfig)
	RootCmd.AddCommand(PromptCmd)

	InitCmd.AddCommand(InitBashCmd)
	InitCmd.AddCommand(InitZshCmd)
	InitCmd.AddCommand(InitFishCmd)
}
