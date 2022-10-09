package command

import (
	"errors"
	"github.com/LarsNorlander/maudio-oxypro49-preset-editor/internal/pkg/preset"
	"gopkg.in/yaml.v3"
	"os"
)

type Command struct {
	ID         string                 `yaml:"id"`
	Parameters map[string]interface{} `yaml:"parameters"`
}

type Function func(parameters map[string]interface{}, preset *preset.Preset) error

func Apply(commands []Command, preset *preset.Preset) error {
	// Copy the data
	for i := 0; i < len(commands); i++ {
		// Lookup the function
		command, ok := Mapping[commands[i].ID]
		if !ok {
			return errors.New("command not found")
		}
		// Apply the results
		err := command(commands[i].Parameters, preset)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadFile(path string) (commands []Command, err error) {
	// Read the file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Unmarshall it
	err = yaml.Unmarshal(data, &commands)
	if err != nil {
		return nil, err
	}
	return
}
