package markdown

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/russross/blackfriday/v2"
)

func LoadMarkdownFromFile(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error loading markdown file", err)
		os.Exit(1)
	}
	return content
}

func MarkdownToHTML(input []byte) string {
	var htmlBytes []byte = blackfriday.Run(input)
	var htmlStr string = string(htmlBytes)
	return htmlStr
}
