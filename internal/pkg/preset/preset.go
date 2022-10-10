package preset

import (
	"errors"
	"github.com/LarsNorlander/maconfig/internal/pkg/fader"
	"os"
)

const (
	FadersStart = 0x016f
	FadersCount = 36
)

type Preset struct {
	data []byte
}

func (p *Preset) GetFaderAt(index int) (*fader.Fader, error) {
	if index >= FadersCount {
		return nil, errors.New("index out of bounds")
	}
	offset := FadersStart + (fader.Size * index)
	return fader.New(p.data[offset : offset+fader.Size])
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
