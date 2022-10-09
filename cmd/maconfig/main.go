package main

import (
	"fmt"
	"github.com/LarsNorlander/maudio-oxypro49-preset-editor/internal/pkg/command"
	"github.com/spf13/cobra"
	"os"
)

const (
	CommandsFlagName = "commands"
	InputFlagName    = "input"
	OutputFlagName   = "output"
)

var (
	CommandsPath     string
	InputPresetPath  string
	OutputPresetPath string
)

var cmd = &cobra.Command{
	Use:  "maconfig",
	RunE: run,
}

func init() {
	cmd.Flags().StringVarP(&CommandsPath, CommandsFlagName, "c", "", "source file for commands")
	_ = cmd.MarkFlagRequired(CommandsFlagName)

	cmd.Flags().StringVarP(&InputPresetPath, InputFlagName, "i", "", "source preferences file")
	_ = cmd.MarkFlagRequired(InputFlagName)

	cmd.Flags().StringVarP(&OutputPresetPath, OutputFlagName, "o", "out.OxygenPro49UserPreset", "source preferences file")
}

func main() {
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, _ []string) error {
	// Read preference file
	preferenceFilePath, err := cmd.Flags().GetString(InputFlagName)
	if err != nil {
		return err
	}
	basePreferences, err := os.ReadFile(preferenceFilePath)
	if err != nil {
		return err
	}

	// Read commands file
	commandsFilePath, err := cmd.Flags().GetString(CommandsFlagName)
	if err != nil {
		return err
	}
	commands, err := command.ParseCommandsFile(commandsFilePath)
	if err != nil {
		return err
	}

	// Apply commands
	out, err := command.Apply(commands, basePreferences)
	if err != nil {
		return err
	}

	// Write the file out
	outFilePath, err := cmd.Flags().GetString(OutputFlagName)
	if err != nil {
		return err
	}
	err = os.WriteFile(outFilePath, out, 0644)
	if err != nil {
		return err
	}

	return nil
}
