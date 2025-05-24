package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/russross/blackfriday/v2"
	"gopkg.in/yaml.v2"
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
		if strings.HasPrefix(key, "fa ") {
			htmlBuilder.WriteString(`<a href="` + iconMap[key] + `"><i class="` + key + `"></i></a>`)
		} else {
			htmlBuilder.WriteString(`<a href="` + iconMap[key] + `">` + key + `</a>`)
		}
	}

	return htmlBuilder.String()
}

func ApplyTemplate(page interface{}, tmplPath string) string {
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

type Image struct {
	Path        string  `yaml:"path"`
	Description string  `yaml:"description"`
	Title       string  `yaml:"title"`
	Latitude    float64 `yaml:"lat"`
	Longitude   float64 `yaml:"long"`
}

type Images struct {
	Images []Image `yaml:"images"`
}

func ParseImages(content string) string {
	var imagesData Images
	err := yaml.Unmarshal([]byte(content), &imagesData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var htmlTags strings.Builder
	for _, image := range imagesData.Images {
		html := fmt.Sprintf(`
	<img src="%s"
        data-title="%s" 
        data-description="%s"
        data-latitude="%f"
        data-longitude="%f"
        data-fullsrc="%s">`,
			image.Path, image.Title, image.Description, image.Latitude, image.Longitude, image.Path)
		htmlTags.WriteString(html)
		htmlTags.WriteString("\n")
	}
	return htmlTags.String()
}
