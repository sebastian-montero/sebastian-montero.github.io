package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sebastian-montero/ssg/pkg/models"
	"github.com/sebastian-montero/ssg/pkg/utils"
)

func main() {
	contentPath := flag.String("img", "", "Path to content file")
	tmplPath := flag.String("tmpl", "", "Path to template file")
	outPath := flag.String("out", "", "Path to output HTML file")
	flag.Parse()

	if *contentPath == "" || *tmplPath == "" || *outPath == "" {
		fmt.Println("Please provide all required flags")
		os.Exit(1)
	}

	var contentBytes []byte = utils.LoadMarkdownFromFile(*contentPath)
	var contentStr string = string(contentBytes)
	var images string = utils.ParseImages(contentStr)

	page := models.GalleryType{
		Images: images,
	}

	var outStr string = utils.ApplyTemplate(page.ToHTML(), *tmplPath)
	utils.SaveHTML(outStr, *outPath)
}
