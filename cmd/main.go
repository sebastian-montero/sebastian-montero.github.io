package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sebastian-montero/ssg/pkg/page"
	"github.com/sebastian-montero/ssg/pkg/parse"
	"github.com/sebastian-montero/ssg/pkg/template"
)

func saveHTML(content string, filename string) error {
	data := []byte(content)

	err := os.WriteFile(filename, data, 0644) // 0644 provides read/write permissions for owner and read for others
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

	var markdownBytes []byte = parse.LoadMarkdownFromFile(*markdownPath)
	var htmlStr string = parse.MarkdownToHTML(markdownBytes)
	var sidebarStr string = parse.ParseSidebar(htmlStr)

	var page page.PageType = page.BuildPageStruct(*title, htmlStr, sidebarStr)
	var outStr string = template.ApplyTemplate(page, *tmplPath)
	saveHTML(outStr, *outPath)
}
