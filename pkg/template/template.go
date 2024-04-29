package template

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/sebastian-montero/ssg/pkg/page"
)

func ApplyTemplate(page page.PageType, tmplPath string) string {
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
