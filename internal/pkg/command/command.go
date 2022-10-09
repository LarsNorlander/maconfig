package command

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type Command struct {
	Name       string            `yaml:"name"`
	Parameters map[string]string `yaml:"parameters"`
}

type Function func(parameters map[string]string, data []byte) (out []byte, err error)

func Apply(commands []Command, data []byte) (out []byte, err error) {
	// Copy the data
	result := data
	for i := 0; i < len(commands); i++ {
		// Lookup the function
		function, ok := Mapping[commands[i].Name]
		if !ok {
			return nil, errors.New("command not found")
		}
		// Apply the results
		result, err = function(commands[i].Parameters, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func ParseCommandsFile(path string) (commands []Command, err error) {
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
