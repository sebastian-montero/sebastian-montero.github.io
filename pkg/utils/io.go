package utils

import (
	"fmt"
	"os"
)

func LoadMarkdownFromFile(filePath string) []byte {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error loading markdown file", err)
		os.Exit(1)
	}
	return content
}

func SaveHTML(content string, filename string) error {
	data := []byte(content)

	err := os.WriteFile(filename, data, 0644) // 0644 provides read/write permissions for owner and read for others
	if err != nil {
		fmt.Println("Error saving HTML file", err)
		os.Exit(1)
	}
	return nil
}