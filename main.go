package main

import (
	"fmt"
	"github.com/LarsNorlander/maconfig/internal/pkg/action"
	"github.com/LarsNorlander/maconfig/internal/pkg/preset"
	"github.com/spf13/cobra"
	"os"
)

const (
	ActionsFlName = "actions"
	InputFlName   = "input"
	OutputFlName  = "output"
)

var (
	ActionsPath      string
	InputPresetPath  string
	OutputPresetPath string
)

var cmd = &cobra.Command{
	Use:  "maconfig",
	RunE: run,
}

func init() {
	cmd.Flags().StringVarP(&ActionsPath, ActionsFlName, "a", "", "source file for actions")
	_ = cmd.MarkFlagRequired(ActionsFlName)

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

	// Read actions file
	actionsPath, err := cmd.Flags().GetString(ActionsFlName)
	if err != nil {
		return err
	}
	actions, err := action.ReadFile(actionsPath)
	if err != nil {
		return err
	}

	// Apply actions
	err = action.Apply(actions, prst)
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
