package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sebastian-montero/ssg/pkg/models"
	"github.com/sebastian-montero/ssg/pkg/utils"
)

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

	var markdownBytes []byte = utils.LoadMarkdownFromFile(*markdownPath)
	var htmlStr string = utils.MarkdownToHTML(markdownBytes)
	var sidebarStr string = utils.ParseSidebar(htmlStr)

	page := models.PageType{
		Title:   *title,
		Content: htmlStr,
		Sidebar: sidebarStr,
	}

	var outStr string = utils.ApplyTemplate(page.ToHTML(), *tmplPath)
	utils.SaveHTML(outStr, *outPath)
}
