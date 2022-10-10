package action

import (
	"fmt"
	"github.com/LarsNorlander/maconfig/internal/pkg/preset"
)

func HelloWorld(parameters map[string]interface{}, _ *preset.Preset) error {
	name, exists := parameters["name"]
	if !exists {
		fmt.Println("Hello, world")
	} else {
		fmt.Printf("Hello, %s\n", name)
	}
	return nil
}
