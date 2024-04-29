package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io/ioutil"
	"os"
)

type PageType struct {
	Title   string
	Content template.HTML
}

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

func applyTemplate(page PageType, tmplPath string) string {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		fmt.Println("Error parsing template", err)
		os.Exit(1)
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, page); err != nil {
		fmt.Println("Error executing template", err)
		os.Exit(1)
	}

	return out.String()
}

func buildPageStruct(title string, content string) PageType {
	var page PageType
	page.Title = title
	page.Content = template.HTML(content)
	return page
}

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

	var markdownBytes []byte = loadMarkdownFromFile(*markdownPath)
	var htmlStr string = markdownToHTML(markdownBytes)
	var page PageType = buildPageStruct(*title, htmlStr)
	var outStr string = applyTemplate(page, *tmplPath)
	saveHTML(outStr, *outPath)
}
