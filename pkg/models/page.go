package models

import (
	"html/template"
)

type PageType struct {
	Title   string
	Content string
	Sidebar string
}

type PageHTMLType struct {
	Title   string
	Content template.HTML
	Sidebar template.HTML
}

func (p PageType) ToHTML() PageHTMLType {
	return PageHTMLType{
		Title:   p.Title,
		Content: template.HTML(p.Content),
		Sidebar: template.HTML(p.Sidebar),
	}
}

type GalleryType struct {
	Images string
}

type GalleryHTMLType struct {
	Images template.HTML
}

func (g GalleryType) ToHTML() GalleryHTMLType {
	return GalleryHTMLType{
		Images: template.HTML(g.Images),
	}
}
