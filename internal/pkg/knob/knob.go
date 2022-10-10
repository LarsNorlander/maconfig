package knob

import "errors"

const (
	Size = 100
	// Property offset values
	ccOffset = 0x19
)

var (
	ErrCorruptKnob = errors.New("corrupt fader")
)

type Knob struct {
	data []byte
}

func New(data []byte) (*Knob, error) {
	if len(data) != 100 {
		return nil, ErrCorruptKnob
	}
	cc := data[ccOffset]
	if cc > 127 {
		return nil, ErrCorruptKnob
	}

	return &Knob{data: data}, nil
}

func (f *Knob) SetCC(val byte) error {
	if val > 127 {
		return errors.New("illegal CC value")
	}
	f.data[ccOffset] = val
	return nil
}
