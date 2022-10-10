package action

import (
	"github.com/LarsNorlander/maconfig/internal/pkg/parser"
	"github.com/LarsNorlander/maconfig/internal/pkg/preset"
)

type HasCC interface {
	SetCC(val byte) error
}

func SetCC(parameters map[string]interface{}, preset *preset.Preset) error {
	// Parse all the parameters
	control, ok := parameters["control"].(string)
	if !ok {
		return ErrUnexpectedParameterType
	}
	posStr, ok := parameters["positions"].(string)
	if !ok {
		return ErrUnexpectedParameterType
	}
	positions, err := parser.ParseNumberList(posStr)
	if err != nil {
		return err
	}
	valStr, ok := parameters["values"].(string)
	if !ok {
		return ErrUnexpectedParameterType
	}
	values, err := parser.ParseNumberList(valStr)

	if len(positions) > len(values) {
		return &notEnoughValuesError{
			expected: len(positions),
			provided: len(values),
		}
	}

	var controls []HasCC
	if control == "fader" {
		for _, position := range positions {
			fader, err := preset.GetFaderAt(position)
			if err != nil {
				return err
			}
			controls = append(controls, fader)
		}
	}
	if control == "knob" {
		for _, position := range positions {
			fader, err := preset.GetKnobAt(position)
			if err != nil {
				return err
			}
			controls = append(controls, fader)
		}
	}

	for i := 0; i < len(controls); i++ {
		err := controls[i].SetCC(byte(values[i]))
		if err != nil {
			return err
		}
	}

	return nil
}
