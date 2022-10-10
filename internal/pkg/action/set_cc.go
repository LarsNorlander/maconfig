package action

import (
	"fmt"
	"github.com/LarsNorlander/maconfig/internal/pkg/parser"
	"github.com/LarsNorlander/maconfig/internal/pkg/preset"
)

type notEnoughCCError struct {
	expected, provided int
}

func (e *notEnoughCCError) Error() string {
	return fmt.Sprintf("not enough CC values provided; expected %d, provided %d", e.expected, e.provided)
}

func SetCC(parameters map[string]interface{}, preset *preset.Preset) error {
	posStr := parameters["positions"].(string)
	positions, err := parser.ParseNumberList(posStr)
	if err != nil {
		return err
	}
	valStr := parameters["values"].(string)
	values, err := parser.ParseNumberList(valStr)

	if len(positions) > len(values) {
		return &notEnoughCCError{
			expected: len(positions),
			provided: len(values),
		}
	}

	for i, position := range positions {
		fader, err := preset.GetFaderAt(position)
		if err != nil {
			return err
		}
		err = fader.SetCC(byte(values[i]))
		if err != nil {
			return err
		}
	}

	return nil
}
