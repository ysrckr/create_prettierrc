package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskForPath() []string {
	fmt.Println("Enter path to create file in: (.)")
	reader := bufio.NewReader(os.Stdin)
	path, _ := reader.ReadString('\n')

	if path == "" || path == "." {
		var err error
		path, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting current working directory:", err)
		}
	}

	fmt.Println("Creating file in:", path)

	pathSlice := strings.Split(path, "/")

	return pathSlice
}
