package fileutils

import "os"

func ReadTextFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		// Couldn't read file
		return "", err
	} else {
		return string(content), nil
	}
}
