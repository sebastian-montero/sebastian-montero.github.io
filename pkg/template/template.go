package template

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io/ioutil"
	"os"
)

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

