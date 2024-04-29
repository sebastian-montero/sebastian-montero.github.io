package markdown

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io/ioutil"
	"os"
)

func loadMarkdownFromFile(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error loading markdown file", err)
		os.Exit(1)
	}
	return content
}

func markdownToHTML(input []byte) string {
	var htmlBytes []byte = blackfriday.Run(input)
	var htmlStr string = string(htmlBytes)
	return htmlStr
}