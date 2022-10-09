package command

import (
	"github.com/LarsNorlander/maudio-oxypro49-preset-editor/internal/pkg/preset"
)

func SetCC(parameters map[string]interface{}, preset *preset.Preset) error {
	err := preset.SetFaderCC(parameters["positions"].(int), byte(parameters["values"].(int)))
	if err != nil {
		return err
	}
	return nil
}