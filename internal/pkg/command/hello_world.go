package command

import "fmt"

func HelloWorld(parameters map[string]string, data []byte) ([]byte, error) {
	name, exists := parameters["name"]
	if !exists {
		fmt.Println("Hello, world")
	} else {
		fmt.Printf("Hello, %s\n", name)
	}
	return data, nil
}
