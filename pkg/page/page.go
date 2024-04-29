package page

import (
	"html/template"
)

type PageType struct {
	Title   string
	Content template.HTML
}

func BuildPageStruct(title string, content string) PageType {
	var page PageType
	page.Title = title
	page.Content = template.HTML(content)
	return page
}
