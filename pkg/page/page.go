package page

import (
	"html/template"
)

type PageType struct {
	Title   string
	Content template.HTML
	Sidebar template.HTML
}

func BuildPageStruct(title string, content string, sidebar string) PageType {
	var page PageType
	page.Title = title
	page.Content = template.HTML(content)
	page.Sidebar = template.HTML(sidebar)
	return page
}
