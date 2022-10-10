package action

import (
	"errors"
	"github.com/LarsNorlander/maconfig/internal/pkg/preset"
	"gopkg.in/yaml.v3"
	"os"
)

type Action struct {
	ID         string                 `yaml:"id"`
	Parameters map[string]interface{} `yaml:"parameters"`
}

type Function func(parameters map[string]interface{}, preset *preset.Preset) error

func Apply(actions []Action, preset *preset.Preset) error {
	// Copy the data
	for i := 0; i < len(actions); i++ {
		// Lookup the function
		action, ok := Mapping[actions[i].ID]
		if !ok {
			return errors.New("action not found")
		}
		// Apply the results
		err := action(actions[i].Parameters, preset)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadFile(path string) (actions []Action, err error) {
	// Read the file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Unmarshall it
	err = yaml.Unmarshal(data, &actions)
	if err != nil {
		return nil, err
	}
	return
}
