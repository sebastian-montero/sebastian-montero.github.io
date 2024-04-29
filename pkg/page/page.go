package page

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

func buildPageStruct(title string, content string) PageType {
	var page PageType
	page.Title = title
	page.Content = template.HTML(content)
	return page
}