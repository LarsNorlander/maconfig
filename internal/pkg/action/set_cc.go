package action

import (
	"github.com/LarsNorlander/maconfig/internal/pkg/preset"
)

func SetCC(parameters map[string]interface{}, preset *preset.Preset) error {
	fader, err := preset.GetFaderAt(parameters["positions"].(int))
	if err != nil {
		return err
	}
	return fader.SetCC(byte(parameters["values"].(int)))
}
