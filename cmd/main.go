package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func saveHTML(content string, filename string) error {
	data := []byte(content)

	err := ioutil.WriteFile(filename, data, 0644) // 0644 provides read/write permissions for owner and read for others
	if err != nil {
		fmt.Println("Error saving HTML file", err)
		os.Exit(1)
	}
	return nil
}

func main() {
	markdownPath := flag.String("md", "", "Path to markdown file")
	tmplPath := flag.String("tmpl", "", "Path to template file")
	title := flag.String("title", "", "Title of the page")
	outPath := flag.String("out", "", "Path to output HTML file")
	flag.Parse()

	if *markdownPath == "" || *tmplPath == "" || *title == "" || *outPath == "" {
		fmt.Println("Please provide all required flags")
		os.Exit(1)
	}

	var markdownBytes []byte = markdown.loadMarkdownFromFile(*markdownPath)
	var htmlStr string = markdown.markdownToHTML(markdownBytes)
	var page PageType = page.buildPageStruct(*title, htmlStr)
	var outStr string = template.applyTemplate(page, *tmplPath)
	saveHTML(outStr, *outPath)
}
