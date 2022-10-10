package fader

import "errors"

const (
	Size = 100
	// Property offset values
	ccOffset = 0x19
)

var (
	ErrCorruptFader = errors.New("corrupt fader")
)

type Fader struct {
	data []byte
}

func New(data []byte) (*Fader, error) {
	if len(data) != 100 {
		return nil, ErrCorruptFader
	}
	cc := data[ccOffset]
	if cc > 127 {
		return nil, ErrCorruptFader
	}

	return &Fader{data: data}, nil
}

func (f *Fader) SetCC(val byte) error {
	if val > 127 {
		return errors.New("illegal CC value")
	}
	f.data[ccOffset] = val
	return nil
}
