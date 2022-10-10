package preset

import (
	"errors"
	"github.com/LarsNorlander/maconfig/internal/pkg/fader"
	"github.com/LarsNorlander/maconfig/internal/pkg/knob"
	"os"
)

const (
	// Faders properties
	fadersStart = 0x016f
	fadersCount = 36
	// Knobs properties
	knobsStart = 0x731b
	knobsCount = 32
)

type Preset struct {
	data []byte
}

func (p *Preset) GetFaderAt(index int) (*fader.Fader, error) {
	if index >= fadersCount {
		return nil, errors.New("index out of bounds")
	}
	offset := fadersStart + (fader.Size * index)
	return fader.New(p.data[offset : offset+fader.Size])
}

func (p *Preset) GetKnobAt(index int) (*knob.Knob, error) {
	if index >= knobsCount {
		return nil, errors.New("index out of bounds")
	}
	offset := knobsStart + (knob.Size * index)
	return knob.New(p.data[offset : offset+knob.Size])
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
