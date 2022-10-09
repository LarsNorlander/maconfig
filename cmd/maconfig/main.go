package main

import (
	"fmt"
	"github.com/LarsNorlander/maudio-oxypro49-preset-editor/internal/pkg/command"
	"github.com/LarsNorlander/maudio-oxypro49-preset-editor/internal/pkg/preset"
	"github.com/spf13/cobra"
	"os"
)

const (
	CommandsFlName = "commands"
	InputFlName    = "input"
	OutputFlName   = "output"
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
	cmd.Flags().StringVarP(&CommandsPath, CommandsFlName, "c", "", "source file for commands")
	_ = cmd.MarkFlagRequired(CommandsFlName)

	cmd.Flags().StringVarP(&InputPresetPath, InputFlName, "i", "", "source preferences file")
	_ = cmd.MarkFlagRequired(InputFlName)

	cmd.Flags().StringVarP(&OutputPresetPath, OutputFlName, "o", "out.OxygenPro49UserPreset", "source preferences file")
}

func main() {
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, _ []string) error {
	// Read preference file
	presetPath, err := cmd.Flags().GetString(InputFlName)
	if err != nil {
		return err
	}
	prst, err := preset.ReadFile(presetPath)
	if err != nil {
		return err
	}

	// Read commands file
	commandsPath, err := cmd.Flags().GetString(CommandsFlName)
	if err != nil {
		return err
	}
	commands, err := command.ReadFile(commandsPath)
	if err != nil {
		return err
	}

	// Apply commands
	err = command.Apply(commands, prst)
	if err != nil {
		return err
	}

	// Write the file out
	outPath, err := cmd.Flags().GetString(OutputFlName)
	if err != nil {
		return err
	}
	err = preset.WriteFile(outPath, prst)
	if err != nil {
		return err
	}

	return nil
}
