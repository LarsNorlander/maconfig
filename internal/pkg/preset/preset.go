package preset

import (
	"errors"
	"os"
)

const (
	FadersStart = 0x016f
	FadersLen   = 36

	FaderSize     = 100
	FaderCCOffset = 0x19
)

type Preset struct {
	data []byte
}

func (p *Preset) SetFaderCC(index int, val byte) error {
	if index >= FadersLen {
		return errors.New("index out of bounds")
	}
	if val > 127 {
		return errors.New("illegal CC value")
	}
	offset := FadersStart + (FaderSize * index) + FaderCCOffset
	p.data[offset] = val
	return nil
}

func ReadFile(path string) (*Preset, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Preset{data: data}, nil
}

func WriteFile(path string, preset *Preset) error {
	err := os.WriteFile(path, preset.data, 0644)
	if err != nil {
		return err
	}
	return nil
}
