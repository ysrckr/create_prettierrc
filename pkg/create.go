package pkg

import (
	"fmt"
	"os"
)

func CreateFile() {

	fbuf, err := os.ReadFile("/Users/yasar/github/projects/create_prettierrc/templates/template.yaml")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	fileData := string(fbuf)

	path := ".prettierrc.yaml"

	// open recently created file
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	// write to file
	_, err = file.WriteString(fileData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	// close file
	err = file.Close()
	if err != nil {
		fmt.Println("Error closing file:", err)
	}
}
