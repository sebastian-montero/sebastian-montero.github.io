package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"regexp"
	"strings"

	"github.com/russross/blackfriday/v2"
	"github.com/sebastian-montero/ssg/pkg/models"
)


func MarkdownToHTML(input []byte) string {
	var htmlBytes []byte = blackfriday.Run(input)
	var htmlStr string = string(htmlBytes)
	return htmlStr
}

func ParseSidebar(content string) string {
	re := regexp.MustCompile(`(?s)<!--(.*?)-->`)
	matches := re.FindStringSubmatch(content)
	if len(matches) == 0 {
		fmt.Println("Nothing to include in the sidebar.")
		return ""
	}
	commentContent := matches[1]

	iconMap := make(map[string]string)
	var orderedKeys []string

	lines := strings.Split(commentContent, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ": ")
		if len(parts) < 2 {
			continue
		}
		iconClass := strings.TrimSpace(parts[0])
		link := strings.TrimSpace(parts[1])
		if _, exists := iconMap[iconClass]; !exists {
			orderedKeys = append(orderedKeys, iconClass) // Append the key only if it's not already in the map
		}
		iconMap[iconClass] = link
	}
	var htmlBuilder strings.Builder
	for _, key := range orderedKeys {
		htmlBuilder.WriteString(`<a href="` + iconMap[key] + `"><i class="` + key + `"></i></a>`)
	}

	return htmlBuilder.String()
}

func ApplyTemplate(page models.PageHTMLType, tmplPath string) string {
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
